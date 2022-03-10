package internal

import (
	"ThumbnailsYouTube_/pkg"
	"ThumbnailsYouTube_/pkg/proto"
)

type Server struct {
	DB *pkg.DB
	proto.UnimplementedThumbnailsServer
}

func New(db *pkg.DB) *Server {
	return &Server{
		DB:                            db,
		UnimplementedThumbnailsServer: proto.UnimplementedThumbnailsServer{},
	}
}
