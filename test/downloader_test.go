package test

import (
	"ThumbnailsYouTube_/pkg/proto"
	"context"
	"fmt"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"testing"
	"time"
)

func TestDownload(t *testing.T) {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	require.NoError(t, err)
	defer conn.Close()
	c := proto.NewThumbnailsClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	urlNotExist := "https://www.youtube.com/watch?v=ceRTzKuCLnw"

	image, err := c.Download(ctx, &wrapperspb.StringValue{Value: urlNotExist})
	fmt.Println(image.String())
	require.NoError(t, err)
	require.Equal(t, "downloaded", image.GetStatus())
	require.NotZero(t, image.GetId())

}
