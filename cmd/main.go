package main

import (
	"ThumbnailsYouTube_/internal"
	"ThumbnailsYouTube_/pkg"
	"os"
)

func main() {
	lite := os.Args

	url, err := pkg.ParceURL(lite)
	if err != nil {
		panic(err)
	}

	internal.Download("pic.jpg", url)
}
