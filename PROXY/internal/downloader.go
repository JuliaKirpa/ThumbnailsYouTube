package internal

import (
	"ThumbnailsYouTube_/PROXY/pkg"
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"context"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io/ioutil"
	"net/http"
)

type Server struct {
	proto.UnimplementedThumbnailsServer
	pkg.DB
}

func (s *Server) Download(ctx context.Context, in *wrapperspb.StringValue) (*proto.Image, error) {

	filename, url, err := pkg.ParceURL(in.GetValue())
	if err != nil {
		return nil, err
	}

	image, err := s.CheckBase(filename)
	if err != nil {
		response, err := http.Get(url)
		if err != nil {
			return nil, err
		}

		defer response.Body.Close()

		img, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		return s.SaveToBase(filename, img)
	}

	return image, nil
}
