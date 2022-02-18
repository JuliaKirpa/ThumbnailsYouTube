package main

import (
	"ThumbnailsYouTube_/PROXY/internal"
	"ThumbnailsYouTube_/PROXY/pkg"
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {
	go ServerStart()

	db, err := pkg.ConnectToBase()
	if err != nil {
		panic(err)
	}
	defer db.Close()
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
