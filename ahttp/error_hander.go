package ahttp

import (
	"context"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

// CustomErrorHandler 自定义错误处理器
func CustomErrorHandler(ctx context.Context, mux *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")

	// 始终返回 200
	w.WriteHeader(http.StatusOK)
	runtime.DefaultHTTPErrorHandler(ctx, mux, marshaler, w, r, err)
}
