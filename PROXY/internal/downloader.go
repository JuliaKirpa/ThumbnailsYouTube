package internal

import (
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"net/http"
)

type Server struct {
	proto.UnimplementedThumbnailsServer
}

func (s *Server) Download(ctx context.Context, in *wrapperspb.StringValue) (*proto.Image, error) {

	response, err := http.Get(in.GetValue())
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		return nil, errors.New("Received not OK response code")
	}
	return &proto.Image{
		Status: "OK",
		Id:     01,
	}, nil

}
