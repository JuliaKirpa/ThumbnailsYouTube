package main

import (
	"ThumbnailsYouTube_/internal"
	"ThumbnailsYouTube_/pkg/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	ServerStart()
}

func ServerStart() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		panic(err)
	}
	server := &internal.Server{}

	GRPCServ := grpc.NewServer()
	proto.RegisterThumbnailsServer(GRPCServ, server)

	if err := GRPCServ.Serve(lis); err != nil {
		panic(err)
	}
}
