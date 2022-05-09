package main

import (
	"context"
	"log"
	"time"

	pb "github.com/skema-repo/WinBeyond/grpc-go/xxx/xxx"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	Address     = "localhost:50051"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewXXAABBClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

    // heathcheckReply
	heathcheckReply, err := c.Heathcheck (ctx, &pb.HealthcheckRequest {
		// Msg: "world",
	})
	HandleReply("heathcheckReply", heathcheckReply, err)
	// log.Printf("Greeting: %s", heathcheckReply.GetMsg())

    // helloworldReply
	helloworldReply, err := c.Helloworld (ctx, &pb.HelloRequest {
		// Msg: "world",
	})
	HandleReply("helloworldReply", helloworldReply, err)
	// log.Printf("Greeting: %s", helloworldReply.GetMsg())
}

// HandleReply handle the reply
func HandleReply(replyName string, reply interface{}, err error) {
	if err != nil {
		log.Fatalf("%s could not greet: %v", replyName, err)
	}
	log.Printf(replyName + " from server: %v\n", reply)
}