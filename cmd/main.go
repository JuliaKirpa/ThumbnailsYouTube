package main

import (
	"ThumbnailsYouTube_/internal"
	"ThumbnailsYouTube_/pkg"
	"flag"
	"os"
)

func main() {
	linc := os.Args
	async := flag.Bool("async", false, "asynchronous downloading")

	url, err := pkg.ParceURL(linc[1:])
	if err != nil {
		panic(err)
	}

	flag.Parse()
	if *async {
		for key, value := range url {
			go internal.Download(key, value)
		}
		return
	}
	for key, value := range url {
		internal.Download(key, value)
	}
}
