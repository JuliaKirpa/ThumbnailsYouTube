package test

import (
	"ThumbnailsYouTube_/internal"
	"ThumbnailsYouTube_/pkg/proto"
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"log"
	"net"
	"testing"
	"time"
)

func initGRPCServerHTTP2() {
	lis, err := net.Listen("tcp", "localhost:50051")

	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	proto.RegisterThumbnailsServer(s, &internal.Server{})

	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
func TestDownload(t *testing.T) {
	initGRPCServerHTTP2()
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := proto.NewThumbnailsClient(conn)

	url := "https://www.youtube.com/watch?v=lCbAt7bp0H4"

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.Download(ctx, &wrapperspb.StringValue{Value: url})
	if err != nil {
		log.Fatalf("Could not add product: %v", err)
	}
	log.Printf("Res %s", r.String())
}
