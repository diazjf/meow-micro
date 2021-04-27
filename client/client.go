package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		span := tracer.StartSpan("send-grpc-request", ext.RPCServerOption(spanCtx))
		defer span.Finish()

		// sleep so we can show DT
		time.Sleep(1 * time.Second)

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

		sleep(r)

		// curl 127.0.0.1:5002 -X POST -d "{\"name\": \"Meower\"}"
		var cat Cat
		err = json.Unmarshal(reqBody, &cat)
		if err != nil {
			log.Printf("Error: %v", err.Error())
			log.Printf("Not talking to any cats.")
		} else {

			if cat.Name == "" {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("400 - Missing name field in body!"))
			} else {
				// perform GRPC call
				resp, err := c.SayHello(context.Background(), &chat.Message{Body: cat.Name})
				if err != nil {
					log.Fatalf("could not greet: %v", err)
				}
				log.Printf("Sending Message: %+v", resp)
				w.WriteHeader(http.StatusOK)
				w.Write([]byte("200 - Meow sent: " + resp.Body))
			}
		}
	})
	log.Printf("Client Started")
	log.Fatal(http.ListenAndServe(RESTPort, nil))
}

func sleep(r *http.Request) {
	os.Setenv("JAEGER_SERVICE_NAME", "meow-client")
	tracer, closer := tracing.Init()
	defer closer.Close()

	spanCtx, _ := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
	span := tracer.StartSpan("sleep", ext.RPCServerOption(spanCtx))
	defer span.Finish()

	fmt.Printf("Sleeping..........")
	time.Sleep(2 * time.Second)
}
