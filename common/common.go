package common

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"net"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/xuanxuan000/zzutil/midware"
)

func GetMD5(data []byte) string {
	md5Contain := md5.New()
	md5Contain.Write(data)
	return hex.EncodeToString(md5Contain.Sum(nil))
}

//GetOutBoundIP 获取本机IP
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:53")
	if err != nil {
		return
	}
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	ip = strings.Split(localAddr.String(), ":")[0]
	return
}

//支持	1.*gin.Context	2.context.Context
//GetSpanFrom 根据父对象生成子span
func NewSpan(parent interface{}, operation string) (span opentracing.Span, ctx context.Context, err error) {
	switch parent := parent.(type) {
	case *gin.Context:
		clue := midware.GetParentSpanContext(parent)
		span, ctx = opentracing.StartSpanFromContext(clue, operation)
	// case opentracing.SpanContext:
	// 	span = opentracing.StartSpan(
	// 		operation,
	// 		opentracing.ChildOf(parent),
	// 		ext.SpanKindRPCClient,
	// 	)
	case context.Context:
		span, ctx = opentracing.StartSpanFromContext(parent, operation)
	default:
		span = opentracing.StartSpan(operation)
	}
	if ctx == nil {
		ctx = opentracing.ContextWithSpan(context.Background(), span)
	}
	span.SetTag(string(ext.Component), operation)
	// defer span.Finish()
	return
}
