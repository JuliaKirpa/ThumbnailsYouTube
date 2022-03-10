package ThumbnailsYouTube_

import (
	"ThumbnailsYouTube_/pkg"
	"ThumbnailsYouTube_/pkg/proto"
)

type Server struct {
	Database *pkg.Database
	proto.UnimplementedThumbnailsServer
}

func New(db *pkg.Database) *Server {
	return &Server{Database: db}
}
