package tracing

import (
	"fmt"
	"io"
	"log"

	opentracing "github.com/opentracing/opentracing-go"
	config "github.com/uber/jaeger-client-go/config"
)

// Init returns an instance of Jaeger Tracer that samples 100% of traces and logs all spans to stdout.
// TODO: https://github.com/jaegertracing/jaeger-client-go/blob/master/config/example_test.go
func Init() (opentracing.Tracer, io.Closer) {

	cfg, err := config.FromEnv()
	if err != nil {
		// parsing errors might happen here, such as when we get a string where we expect a number
		panic(fmt.Sprintf("Could not parse Jaeger env vars: %s", err.Error()))
	}

	tracer, closer, err := cfg.NewTracer()
	if err != nil {
		log.Printf("Could not initialize jaeger tracer: %s", err.Error())
		panic(fmt.Sprintf("Could not initialize jaeger tracer: %s", err.Error()))
	}

	return tracer, closer
}
