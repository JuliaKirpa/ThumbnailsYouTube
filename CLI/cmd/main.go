package main

import (
	"ThumbnailsYouTube_/PROXY/pkg"
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/wrapperspb"
	"os"
	"time"
)

func main() {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(cwt, "localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	client := proto.NewThumbnailsClient(conn)

	linc := os.Args
	async := flag.Bool("async", false, "asynchronous downloading")

	url, err := pkg.ParceURL(linc[1:])
	if err != nil {
		panic(err)
	}

	flag.Parse()
	if *async {
		go client.DownloadAsync(context.Background())
	}
	for _, value := range url {
		fmt.Println(client.Download(context.Background(), &wrapperspb.StringValue{Value: value}))
	}
}
