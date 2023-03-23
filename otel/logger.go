package main

import (
	"context"

	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

type LoggerWithCtx struct {
	l       *zap.Logger
	// Parent context which will be same across a handler
	// Can't remove it because it attache
	context context.Context
}

func NewLoggerWithCtx(logger *zap.Logger, ctx context.Context) *LoggerWithCtx {
	return &LoggerWithCtx{
		l:       logger,
		context: ctx,
	}
}

func (lCtx *LoggerWithCtx) Info(currentCtx context.Context, msg string, fields ...zap.Field) {
	span := trace.SpanFromContext(currentCtx)
	if span.IsRecording() {
		context := span.SpanContext()
		spanField := zap.String("span_id", context.SpanID().String())
		traceField := zap.String("trace_id", context.TraceID().String())
		traceFlags := zap.Int("trace_flags", int(context.TraceFlags()))
		fields = append(fields, []zap.Field{spanField, traceField, traceFlags}...)

		attrs := make([]attribute.KeyValue, 0)
		logSeverityKey := attribute.Key("log.severity")
		logMessageKey := attribute.Key("log.message")
		attrs = append(attrs, logSeverityKey.String("Info"))
		attrs = append(attrs, logMessageKey.String(msg))
		span.AddEvent("log", trace.WithAttributes(attrs...))
	}
	lCtx.l.Info(msg, fields...)
}
