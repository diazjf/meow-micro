package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/diazjf/meow-micro/chat"
	"google.golang.org/grpc"

	"github.com/diazjf/meow-micro/tracing"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

const (
	GRPCAddress = "localhost:5001"
	RESTPort    = ":5002"
	defaultName = "Cats of the World!"
)

type Cat struct {
	Name string `json:"name"`
}

func main() {
	// Add tracer for the HTTPHandle function
	os.Setenv("JAEGER_SERVICE_NAME", "meow-client")
	tracer, closer := tracing.Init()
	defer closer.Close()

	opentracing.SetGlobalTracer(tracer)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("send-meow-communication", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		// Get env variable for GRPC Server Address
		sAddr := os.Getenv("GRPC_SERVER_ADDRESS")
		if sAddr == "" {
			sAddr = GRPCAddress
		}
		log.Printf("GRPC Server set to %v", sAddr)
		log.Printf("GRPC Connection Started")
		// Set up a connection to the GRPC server
		conn, err := grpc.Dial(sAddr, grpc.WithInsecure(), grpc.WithBlock())
		if err != nil {
			log.Fatalf("did not connect: %v", err)
		}
		defer conn.Close()
		log.Printf("GRPC Connection Established")

		// create a from the proto
		c := chat.NewChatServiceClient(conn)

		// Grab what was sent in the request
		reqBody, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error: %v", err.Error())
		}

		// curl 127.0.0.1:5002 -X POST -d "{\"name\": \"Meower\"}"
		// lookup https://stackoverflow.com/questions/15672556/handling-json-post-request-in-go
		var cat Cat
		err = json.Unmarshal(reqBody, &cat)
		if err != nil {
			log.Printf("Error: %v", err.Error())
			log.Printf("Not talking to any cats.")
		} else {
			// perform GRPC call
			resp, err := c.SayHello(context.Background(), &chat.Message{Body: cat.Name})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Sending Message: %+v", resp)
		}

	})
	log.Printf("Client Started")
	log.Fatal(http.ListenAndServe(RESTPort, nil))
}
