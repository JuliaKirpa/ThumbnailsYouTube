package test

import (
	"ThumbnailsYouTube_/PROXY/pkg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestParceURL(t *testing.T) {
	name, url, err := pkg.ParceURL("https://www.youtube.com/watch?v=F5tSoaJ93ac&list=RDF5tSoaJ93ac&start_radio=1")

	require.NoError(t, err)
	require.NotEmpty(t, name, url)
	require.Equal(t, "F5tSoaJ93ac", name)
	require.Equal(t, "https://img.youtube.com/vi/F5tSoaJ93ac/0.jpg", url)
}
