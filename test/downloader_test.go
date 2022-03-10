package test

import (
	"ThumbnailsYouTube_/internal"
	"ThumbnailsYouTube_/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
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

}
