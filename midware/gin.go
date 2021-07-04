package midware

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/xuanxuan000/zzutil/log"
)

const convertfail = " convert fail"

func GetSpan(c *gin.Context) (span opentracing.Span) {
	temp, exist := c.Get("Span")
	if !exist {
		log.Error("Span", convertfail)
		return
	}
	span = temp.(opentracing.Span)
	return
}

func GetParentSpanContext(c *gin.Context) (ctx context.Context) {
	temp, exist := c.Get("ParentSpanContext")
	if !exist {
		log.Error("ParentSpanContext", convertfail)
		ctx = context.Background()
		return
	}
	ctx = temp.(context.Context)
	return
}
