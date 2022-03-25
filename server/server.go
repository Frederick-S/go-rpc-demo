package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"

	"example.com/playground/pb"
)

type SumService struct {
}

func (sumService *SumService) Sum(sumRequest *pb.SumRequest, sumReplay *pb.SumReply) error {
	sumReplay.Result = sumRequest.A + sumRequest.B

	return nil
}

func main() {
	sumService := &SumService{}
	rpc.Register(sumService)
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")

	if err != nil {
		log.Fatal("listen error:", err)
	}

	fmt.Println("Listening on port 1234")
	http.Serve(listener, nil)
}
