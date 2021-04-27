package tracing

import (
	"fmt"
	"io"

	opentracing "github.com/opentracing/opentracing-go"
	config "github.com/uber/jaeger-client-go/config"
)

// Initialize the Tracer using environment variables
func Init() (opentracing.Tracer, io.Closer) {
	cfg, err := config.FromEnv()
	if err != nil {
		panic(fmt.Sprintf("Could not parse Jaeger env vars: %s", err.Error()))
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		panic(fmt.Sprintf("Could not initialize jaeger tracer: %s", err.Error()))
	}

	return tracer, closer
}
