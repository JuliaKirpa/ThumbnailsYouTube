package main

import (
	"ThumbnailsYouTube_/PROXY/pkg"
	"ThumbnailsYouTube_/PROXY/pkg/proto"
	"context"
	"flag"
	"google.golang.org/grpc"
	"os"
	"time"
)

func main() {
	cwt, _ := context.WithTimeout(context.Background(), time.Second*5)
	conn, err := grpc.DialContext(cwt, "50051", grpc.WithInsecure(), grpc.WithBlock())
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
		for _, value := range url {
			go client.DownloadAsync(value)
		}
		return
	}
	for _, value := range url {
		client.Download(value)
	}
}
