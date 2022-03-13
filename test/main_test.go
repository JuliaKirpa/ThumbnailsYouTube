package test

import (
	"ThumbnailsYouTube_/internal"
	"ThumbnailsYouTube_/pkg"
	"ThumbnailsYouTube_/pkg/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"testing"
)

func TestMain(m *testing.M) {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	storage := pkg.New()
	defer storage.Close()

	er := storage.PrepareBase()
	if er != nil {
		log.Fatalf("err from preparing db %s", er)
	}

	proto.RegisterThumbnailsServer(s, &internal.Server{DB: storage})

	reflection.Register(s)
	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()
}
