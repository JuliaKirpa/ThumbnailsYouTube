package test

import (
	"ThumbnailsYouTube_/pkg"
	"github.com/stretchr/testify/require"
	"log"
	"testing"
)

func TestSaveToBase(t *testing.T) {

	storage := pkg.New()
	defer storage.Close()

	er := storage.PrepareBase()
	if er != nil {
		log.Fatalf("err from preparing db %s", er)
	}

	filename := "lCbAt7bp0H4"
	image := []byte{214, 46, 220, 83, 160, 73, 40, 39, 201, 155, 19, 202, 3, 11, 191, 178, 56,
		74, 90, 36, 248, 103, 18, 144, 170, 163, 145, 87, 54, 61, 34, 220, 222,
		207, 137, 149, 173, 14, 92, 120, 206, 222, 158, 28, 40, 24, 30, 16, 175,
		108, 128, 35, 230, 118, 40, 121, 113, 125, 216, 130, 11, 24, 90, 48, 194,
		240, 105, 44, 76, 34, 57, 249, 228, 125, 80, 38, 9, 136, 29, 117, 207, 139,
		168, 181, 85, 137, 126, 10, 126, 242, 120, 247, 121, 8, 100, 12, 201, 171,
		38, 226, 193, 180, 190, 117, 177, 87, 143, 242, 213, 11, 44, 180, 113, 93,
		106, 99, 179, 68, 175, 211, 164, 116, 64, 148, 226, 254, 172, 147}

	img, err := storage.SaveToBase(filename, image)

	require.NoError(t, err)
	require.NotZero(t, img.Id)
	require.Equal(t, "downloaded", img.Status)

	defer func() {
		err := storage.Clean(img.Id)
		if err != nil {
			log.Fatalf("err to del from db: %s", err)
		}
		storage.Close()
	}()
}

func TestCheckBase(t *testing.T) {
	storage := pkg.New()
	defer storage.Close()

	er := storage.PrepareBase()
	if er != nil {
		log.Fatalf("err from preparing db %s", er)
	}

	filenameExist := "lCbAt7bp0H4"

	img, err := storage.CheckBase(filenameExist)
	require.NoError(t, err)
	require.Equal(t, "already downloaded", img.Status)
	require.NotZero(t, img.Id)

	filenameNotExist := "lCbAt7bp0H6"

	resp, err := storage.CheckBase(filenameNotExist)
	require.Error(t, err)
	require.Empty(t, resp)

	defer func() {
		storage.Close()
	}()
}
