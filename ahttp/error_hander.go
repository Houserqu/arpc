package ahttp

import (
	"context"
	"net/http"
	"strconv"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// CustomErrorHandler 自定义错误处理器
func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	// 默认返回 200
	httpStatusCode := 200

	md, ok := runtime.ServerMetadataFromContext(ctx)
	if ok {
		// 读取 gRPC Header 和 Trailer 中的 grpc-gateway-http-status
		if values, exist := md.HeaderMD["grpc-gateway-http-status"]; exist {
			httpStatusCode, _ = strconv.Atoi(values[0])
		} else if values, exist := md.TrailerMD["grpc-gateway-http-status"]; exist {
			httpStatusCode, _ = strconv.Atoi(values[0])
		}
	}

	w.WriteHeader(httpStatusCode)

	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}
