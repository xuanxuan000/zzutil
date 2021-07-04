package midware

import (
	"context"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// opentracing.SetGlobalTracer()  //设置全局的tracer
// opentracing.StartSpan() //开始一个span
// opentracing.GlobalTracer().Inject() //把span信息注入到http请求里面
// opentracing.GlobalTracer().Extract()  //从http请求里面解析出上一个服务的span信息
// const defaultComponentName = "net/http"
const JaegerOpen = 1
const JaegerHostPort = "127.0.0.1:6831"

func OpenTracing(tracer opentracing.Tracer) gin.HandlerFunc {
	return func(c *gin.Context) {
		if JaegerOpen == 1 {
			var parentSpan opentracing.Span
			// tracer, closer := NewJaegerTracer()
			// defer closer.Close()
			//从http请求里面解析出上一个服务的span信息
			spCtx, err := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
			if err != nil {
				parentSpan = tracer.StartSpan(c.Request.URL.Path)
				defer parentSpan.Finish()
			} else {
				parentSpan = opentracing.StartSpan(
					c.Request.URL.Path,
					opentracing.ChildOf(spCtx),
					opentracing.Tag{Key: string(ext.Component), Value: "HTTP"},
					ext.SpanKindRPCServer,
				)
				defer parentSpan.Finish()
			}
			ctx := opentracing.ContextWithSpan(context.Background(), parentSpan)
			c.Set("Tracer", tracer)
			c.Set("Span", parentSpan)
			// c.Set("ParentSpanContext", parentSpan.Context())
			c.Set("ParentSpanContext", ctx)
		}
		c.Next()
	}
}
