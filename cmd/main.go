package main

import (
	"ThumbnailsYouTube_/internal"
	"ThumbnailsYouTube_/pkg"
	"ThumbnailsYouTube_/pkg/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("error from lister gRPC server: %s", err)
	}

	db := pkg.New()
	defer db.Close()

	er := db.PrepareBase()
	if er != nil {
		log.Fatalf("err from preparing db %s", er)
	}

	server := internal.New(db)

	GRPCServ := grpc.NewServer()
	proto.RegisterThumbnailsServer(GRPCServ, server)

	if err := GRPCServ.Serve(lis); err != nil {
		log.Fatalf("error from serv gRPC server: %s", err)
	}

	defer GRPCServ.GracefulStop()
}
