package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/diazjf/meow-micro/chat"
	"google.golang.org/grpc"
)

const (
	GRPCAddress = "localhost:5001"
	RESTPort    = ":5002"
	defaultName = "Cat"
)

type Cat struct {
	name string `json:"name"`
}

func sendMeow(w http.ResponseWriter, r *http.Request) {

	// Set up a connection to the GRPC server
	conn, err := grpc.Dial(GRPCAddress, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// create a from the protot
	c := chat.NewChatServiceClient(conn)

	// Grab what was sent in the request
	reqBody, _ := ioutil.ReadAll(r.Body)

	var cat Cat
	err = json.Unmarshal(reqBody, &cat)
	if err != nil {
		cat = Cat{
			name: defaultName,
		}
	}
	log.Printf("Request: %+v", cat)
	log.Printf("Cat Name: %s", cat.name)

	resp, err := c.SayHello(context.Background(), &chat.Message{Body: cat.name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Meowwwwwwwww: %+v", resp)
}

func main() {
	http.HandleFunc("/", sendMeow)
	log.Fatal(http.ListenAndServe(RESTPort, nil))
}
