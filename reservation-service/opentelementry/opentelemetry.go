package opentelementry

import (
	"go.opentelemetry.io/otel/codes"
	otelTrace "go.opentelemetry.io/otel/trace"
	"log"
	"os"

	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

var Tp *trace.TracerProvider

const ServiceName = "reservation-service"

func InitTracer() (*trace.TracerProvider, error) {
	url := os.Getenv("JAEGER_ENDPOINT")
	if len(url) > 0 {
		return initJaegerTracer(url)
	} else {
		return initFileTracer()
	}
}

func initFileTracer() (*trace.TracerProvider, error) {
	log.Println("Initializing tracing to traces.json")
	f, err := os.Create("traces.json")
	if err != nil {
		return nil, err
	}
	exporter, err := stdouttrace.New(
		stdouttrace.WithWriter(f),
		stdouttrace.WithPrettyPrint(),
	)
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithSampler(trace.AlwaysSample()),
	), nil
}

func initJaegerTracer(url string) (*trace.TracerProvider, error) {
	log.Printf("Initializing tracing to jaeger at %s\n", url)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(jaeger.WithEndpoint(url)))
	if err != nil {
		return nil, err
	}
	return trace.NewTracerProvider(
		trace.WithBatcher(exporter),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(ServiceName),
		)),
	), nil
}

func HttpError(err error, span otelTrace.Span, status int) {
	log.Println(err.Error())
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
}
