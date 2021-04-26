package chat

import (
	"log"
	"os"

	"github.com/diazjf/meow-micro/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
	"golang.org/x/net/context"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, in *Message) (*Message, error) {
	// Add Tracer for SayHello
	os.Setenv("JAEGER_SERVICE_NAME", "say-hello")
	tracer, closer := tracing.Init()
	defer closer.Close()

	//Span until the end of this function
	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, context.Background())
	span := tracer.StartSpan("format", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	log.Printf("Received message: %s", in.Body)
	return &Message{Body: in.Body}, nil
}
