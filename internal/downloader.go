package internal

import (
	"ThumbnailsYouTube_/pkg"
	"ThumbnailsYouTube_/pkg/proto"
	"context"
	"errors"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"io/ioutil"
	"log"
	"net/http"
)

type Server struct {
	proto.UnimplementedThumbnailsServer
}

func (s *Server) Download(ctx context.Context, in *wrapperspb.StringValue) (*proto.Image, error) {

	db, err := pkg.ConnectToBase()
	if err != nil {
		log.Fatal("error to connecting DB")
	}
	defer db.Close()

	filename, url, err := pkg.ParceURL(in.GetValue())
	if err != nil {
		return nil, errors.New("can't parce URL")
	}

	image, err := db.CheckBase(filename)
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

		return db.SaveToBase(filename, img)
	}

	return image, nil
}
