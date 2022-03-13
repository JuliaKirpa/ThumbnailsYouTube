package internal

import (
	"ThumbnailsYouTube_/pkg"
	"ThumbnailsYouTube_/pkg/proto"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io/ioutil"
	"net/http"
)

func (s *Server) Download(ctx context.Context, in *wrapperspb.StringValue) (*proto.Image, error) {

	filename, url, err := pkg.ParceURL(in.GetValue())
	if err != nil {
		return nil, errors.New("can't parce URL")
	}

	image, err := s.DB.CheckBase(filename)
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

		sdImage, err := s.DB.SaveToBase(filename, img)
		return &proto.Image{
			Status: sdImage.Status,
			Id:     sdImage.Id,
		}, nil
	}

	return &proto.Image{
		Status: image.Status,
		Id:     image.Id,
	}, nil
}

//func (s *Server) DownloadAsync(stream proto.Thumbnails_DownloadAsyncServer) error {
//	for {
//		url, err := stream.Recv()
//		if err == io.EOF {
//			return nil
//		}
//		if err != nil {
//			return err
//		}
//
//	}
//}
