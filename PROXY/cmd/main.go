package main

import (
	"ThumbnailsYouTube_/PROXY/internal"
	"ThumbnailsYouTube_/PROXY/pkg"
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"google.golang.org/grpc"
	"net"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go ServerStart(wg)

	db, err := pkg.ConnectToBase()
	if err != nil {
		panic(err)
	}
	defer db.Close()

	wg.Wait()
}

func ServerStart(group sync.WaitGroup) {
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
	group.Done()
}
