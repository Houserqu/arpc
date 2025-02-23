package arpc

import (
	"context"
	"fmt"
	"log"
	"strings"

	uuid "github.com/satori/go.uuid"
	"google.golang.org/grpc/metadata"
)

// 定义日志级别
const (
	INFO  = "INFO"
	WARN  = "WARN"
	ERROR = "ERROR"
	DEBUG = "DEBUG"
)

// logWithFields 根据日志级别和字段输出日志
func LogWithFields(ctx context.Context, level string, format string, a ...any) {
	// 获取 traceID 和时间戳
	traceID := getRequestID(ctx)

	// 构建日志字段
	logEntry := []string{
		"level=" + level,
		"requestID=" + traceID,
		fmt.Sprintf(format, a...),
	}

	// 打印日志
	log.Println(strings.Join(logEntry, " "))
}

// getRequestID 从上下文中获取 requestID
func getRequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if ok {
		if requestIDs, exists := md["request-id"]; exists {
			return requestIDs[0]
		}
	}
	requestID := uuid.NewV4().String()
	ctx = metadata.NewOutgoingContext(ctx, metadata.Pairs("request-id", requestID))
	return requestID
}

// Info 输出信息级别日志
func LogInfo(ctx context.Context, format string, a ...any) {
	LogWithFields(ctx, INFO, format, a...)
}

// Warn 输出警告级别日志
func LogWarn(ctx context.Context, format string, a ...any) {
	LogWithFields(ctx, WARN, format, a...)
}

// Error 输出错误级别日志
func LogError(ctx context.Context, format string, a ...any) {
	LogWithFields(ctx, ERROR, format, a...)
}

// Debug 输出调试级别日志
func LogDebug(ctx context.Context, format string, a ...any) {
	LogWithFields(ctx, DEBUG, format, a...)
}
