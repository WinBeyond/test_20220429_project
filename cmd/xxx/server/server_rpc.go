// Package server provides grpc server implement
package server

import (
	"context"
	"log"

	"xxx/internal/xxx"
	pb "github.com/skema-repo/WinBeyond/grpc-go/xxx/xxx"
)

// Xxx
type rpcXxxServer struct{
    pb.UnimplementedXXAABBServer
}

// NewServer: Create new grpc server instance
func NewServer() pb.XXAABBServer {
	svr := &rpcXxxServer {
		// init custom fileds
	}
	return svr
}
// Heathcheck
func (s *rpcXxxServer) Heathcheck(ctx context.Context, req *pb.HealthcheckRequest) (rsp *pb.HealthcheckResponse,err error) {
	// implement business logic here ...
	// ...

	log.Printf("Received from Heathcheck request: %v", req)
	rsp = &pb.HealthcheckResponse{
		// Msg: "Hello " + req.GetMsg(),
	}

	service := xxx.NewService()
	service.Helloworld()
	return rsp,err
}
// Helloworld
func (s *rpcXxxServer) Helloworld(ctx context.Context, req *pb.HelloRequest) (rsp *pb.HelloReply,err error) {
	// implement business logic here ...
	// ...

	log.Printf("Received from Helloworld request: %v", req)
	rsp = &pb.HelloReply{
		// Msg: "Hello " + req.GetMsg(),
	}

	service := xxx.NewService()
	service.Helloworld()
	return rsp,err
}
