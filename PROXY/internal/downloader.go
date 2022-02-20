package internal

import (
	"ThumbnailsYouTube_/PROXY/pkg"
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io"
	"net/http"
	"os"
)

type Server struct {
	proto.UnimplementedThumbnailsServer
}

func (s *Server) Download(ctx context.Context, in *wrapperspb.StringValue) (*proto.Image, error) {

	filename, url, err := pkg.ParceURL(in.GetValue())
	if err != nil {
		return nil, err
	}

	output, err := os.Create("PROXY/internal/assets/" + filename + ".jpg")
	if err != nil {
		return nil, err
	}

	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	io.Copy(output, response.Body)

	return &proto.Image{
		Status: "OK",
		Id:     01,
	}, nil
}
