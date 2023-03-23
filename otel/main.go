package main

import (
	"context"
	"fmt"
	"net/http"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
)

const serviceName = "AdderSvc"

func main() {
	ctx := context.Background()
	{
		tp, err := setUpTracing(ctx, serviceName)
		if err != nil {
			panic(err)
		}
		defer tp.Shutdown(ctx)

		mp, err := setupMetrics(ctx, serviceName)
		if err != nil {
			panic(err)
		}
		defer mp.Shutdown(ctx)
	}
	go serviceA(ctx, 7071)
	serviceB(ctx, 7072)
}

func serviceB(ctx context.Context, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/serviceB", serviceBHandler)

	// Wrap passed handler in span with given name
	handler := otelhttp.NewHandler(mux, "serverB.http")
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}

	fmt.Println("ServerB listening on", port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func serviceBHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("myTracer").Start(r.Context(), "serviceB_HttpHandler")
	defer span.End()

	// Context is needed, to attach log with this span
	sbLogger := NewLoggerWithCtx(zap.NewExample(), ctx)
	sbLogger.Info(ctx, "Service B Handler")
	answer := add(ctx, 42, 8, sbLogger)

	w.Header().Add("SVC-RESPONSE", fmt.Sprint(answer))
	fmt.Fprintf(w, "hello from serviceB: Answer is: %d", answer)
}

func add(ctx context.Context, x, y int64, logger *LoggerWithCtx) int64 {
	
	ctx, span := otel.Tracer("myTracer").Start(
		ctx,
		"add",
		// add labels/tags/resources(if any) that are specific to this scope.
		trace.WithAttributes(attribute.String("component", "addition")),
		trace.WithAttributes(attribute.String("someKey", "someValue")),
		trace.WithAttributes(attribute.Int("age", 89)),
	)
	defer span.End()
	{
		// Need to understand
		counter, _ := global.MeterProvider().
			Meter("instrumentation/package/name", metric.WithInstrumentationVersion("0.0.1")).
			Int64Counter("add_counter",
				instrument.WithDescription("how many times add function has been called."))

		counter.Add(
			ctx,
			1,
			// labels/tags
			attribute.String("component", "addition"),
			attribute.Int("age", 89),
		)
	}
	// Context passed as argument is different from context passed in logger
	// logger has context of Handler while context passed as argument is of add function
	logger.Info(ctx, fmt.Sprintf("Add function completed successfully. Arguments %v, %v", x, y))
	return x + y
}

func serviceA(ctx context.Context, port int) {
	mux := http.NewServeMux()
	mux.HandleFunc("/serviceA", serviceAHandler)
	// Wrap passed handler in span with given name
	handler := otelhttp.NewHandler(mux, "serverA.http")

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: handler,
	}
	fmt.Println("ServerA listening on", port)
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}

func serviceAHandler(w http.ResponseWriter, r *http.Request) {
	ctx, span := otel.Tracer("myTracer").Start(r.Context(), "serviceAHandler")
	defer span.End()

	cli := &http.Client{
		// NewTransport wraps the provided http.RoundTripper with one that starts a span
		// and injects the span context into the outbound request headers.
		Transport: otelhttp.NewTransport(http.DefaultTransport),
	}
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://127.0.0.1:7072/serviceB", nil)
	if err != nil {
		panic(err)
	}
	resp, err := cli.Do(req)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != http.StatusOK {
		panic(fmt.Errorf("no OK response"))
	}
	w.Header().Add("SVC-RESPONSE", resp.Header.Get("SVC-RESPONSE"))
}
