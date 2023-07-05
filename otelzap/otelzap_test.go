package otelzap

import (
	"context"
	"errors"
	"fmt"
	"testing"

	"go.uber.org/zap/zaptest/observer"

	"github.com/stretchr/testify/require"
	"go.opentelemetry.io/otel/attribute"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.opentelemetry.io/otel/semconv"
	"go.uber.org/zap"

	"github.com/rlakhtakia/opentelemetry-go-extra/otelutil"
)

type Test struct {
	log     func(ctx context.Context, log *Logger)
	require func(t *testing.T, event otelutil.Event)
}

func TestOtelZap(t *testing.T) {
	tests := []Test{
		{
			log: func(ctx context.Context, log *Logger) {
				log.Ctx(ctx).Info("hello")
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "INFO", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				requireCodeAttrs(t, m)
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log.InfoContext(ctx, "hello")
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "INFO", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				requireCodeAttrs(t, m)
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log.Ctx(ctx).Warn("hello", zap.String("foo", "bar"))
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "WARN", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				foo, ok := m["foo"]
				require.True(t, ok)
				require.Equal(t, "bar", foo.AsString())

				requireCodeAttrs(t, m)
			},
		}, {
			log: func(ctx context.Context, log *Logger) {
				err := errors.New("some error")
				log.Ctx(ctx).Error("hello", zap.Error(err))
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "ERROR", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				excTyp, ok := m[semconv.ExceptionTypeKey]
				require.True(t, ok)
				require.Equal(t, "*errors.errorString", excTyp.AsString())

				excMsg, ok := m[semconv.ExceptionMessageKey]
				require.True(t, ok)
				require.Equal(t, "some error", excMsg.AsString())

				requireCodeAttrs(t, m)
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log = log.Clone(WithStackTrace(true))
				log.Ctx(ctx).Info("hello")
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				stack, ok := m[semconv.ExceptionStacktraceKey]
				require.True(t, ok)
				require.NotZero(t, stack.AsString())

				requireCodeAttrs(t, m)
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log.Sugar().ErrorwContext(ctx, "hello", "foo", "bar")
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "ERROR", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				foo, ok := m["foo"]
				require.True(t, ok)
				require.NotZero(t, foo.AsString())

				requireCodeAttrs(t, m)
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log.Sugar().ErrorfContext(ctx, "hello %s", "world")
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "ERROR", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello world", msg.AsString())

				tpl, ok := m[logTemplateKey]
				require.True(t, ok)
				require.Equal(t, "hello %s", tpl.AsString())

				requireCodeAttrs(t, m)
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log.Sugar().Ctx(ctx).Errorw("hello", "foo", "bar")
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "ERROR", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				foo, ok := m["foo"]
				require.True(t, ok)
				require.NotZero(t, foo.AsString())

				requireCodeAttrs(t, m)
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log.Sugar().InfowContext(ctx, "hello", "foo", "bar")
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "INFO", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello", msg.AsString())

				foo, ok := m["foo"]
				require.True(t, ok)
				require.NotZero(t, foo.AsString())

				requireCodeAttrs(t, m)
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log.Sugar().InfowContext(ctx, "sugary logs require keyAndValues to come in pairs", "so this is invalid, but it shouldn't panic")
			},
			require: func(t *testing.T, event otelutil.Event) {
				// no panic? success!
			},
		},
		{
			log: func(ctx context.Context, log *Logger) {
				log.Sugar().Ctx(ctx).Errorf("hello %s", "world")
			},
			require: func(t *testing.T, event otelutil.Event) {
				m := attrMap(event.Attributes)

				sev, ok := m[logSeverityKey]
				require.True(t, ok)
				require.Equal(t, "ERROR", sev.AsString())

				msg, ok := m[logMessageKey]
				require.True(t, ok)
				require.Equal(t, "hello world", msg.AsString())

				tpl, ok := m[logTemplateKey]
				require.True(t, ok)
				require.Equal(t, "hello %s", tpl.AsString())

				requireCodeAttrs(t, m)
			},
		},
	}

	logger := New(zap.L(), WithMinLevel(zap.InfoLevel))

	for i, test := range tests {
		test := test
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			sr := otelutil.NewSpanRecorder()
			provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(sr))
			tracer := provider.Tracer("test")

			ctx := context.Background()
			ctx, span := tracer.Start(ctx, "main")

			test.log(ctx, logger)

			span.End()

			spans := sr.Ended()
			require.Equal(t, 1, len(spans))

			events := spans[0].Events()
			require.Equal(t, 1, len(events))

			event := otelutil.Event{
				Name:                  events[0].Name,
				Time:                  events[0].Time,
				Attributes:            events[0].Attributes,
				DroppedAttributeCount: events[0].DroppedAttributeCount,
			}
			require.Equal(t, "log", event.Name)
			test.require(t, event)
		})
	}

	t.Run("providing extra fields to be recorded on the span, and logged", func(t *testing.T) {
		sr := otelutil.NewSpanRecorder()
		provider := sdktrace.NewTracerProvider(sdktrace.WithSpanProcessor(sr))
		tracer := provider.Tracer("test")

		ctx := context.Background()
		ctx, span := tracer.Start(ctx, "main")

		core, observedLogs := observer.New(zap.InfoLevel)
		logger := New(zap.New(core), WithMinLevel(zap.InfoLevel))
		loggerWithCtx := logger.Ctx(ctx).Clone(WithExtraFields(
			zap.String("foo", "bar"),
			zap.String("MyTraceIDKey", span.SpanContext().TraceID().String()),
		))
		loggerWithCtx.Info("hello")

		span.End()

		spans := sr.Ended()
		require.Equal(t, 1, len(spans))

		events := spans[0].Events()
		require.Equal(t, 1, len(events))

		event := events[0]
		require.Equal(t, "log", event.Name)

		m := attrMap(event.Attributes)
		foo, ok := m["foo"]
		require.True(t, ok)
		require.Equal(t, "bar", foo.AsString())

		_, ok = m["MyTraceIDKey"]
		require.True(t, ok)
		requireCodeAttrs(t, m)

		require.Equal(t, 1, observedLogs.Len())
		require.Equal(t, "hello", observedLogs.All()[0].Message)
		require.Equal(t, zap.InfoLevel, observedLogs.All()[0].Level)

		contextMap := observedLogs.All()[0].ContextMap()
		require.Equal(t, "bar", contextMap["foo"])
		require.Equal(t, span.SpanContext().TraceID().String(), contextMap["MyTraceIDKey"])
	})
}

func requireCodeAttrs(t *testing.T, m map[attribute.Key]attribute.Value) {
	fn, ok := m[semconv.CodeFunctionKey]
	require.True(t, ok)
	require.Contains(t, fn.AsString(), "otelzap.TestOtelZap")

	file, ok := m[semconv.CodeFilepathKey]
	require.True(t, ok)
	require.Contains(t, file.AsString(), "otelzap/otelzap_test.go")

	_, ok = m[semconv.CodeLineNumberKey]
	require.True(t, ok)
}

func attrMap(attrs []attribute.KeyValue) map[attribute.Key]attribute.Value {
	m := make(map[attribute.Key]attribute.Value, len(attrs))
	for _, kv := range attrs {
		m[kv.Key] = kv.Value
	}
	return m
}
