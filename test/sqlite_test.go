package test

import (
	mock_pkg "ThumbnailsYouTube_/test/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSaveToBase(t *testing.T) {
	req := require.New(t)

	mockController := gomock.NewController(t)
	defer mockController.Finish()

	mockExtractor := mock_pkg.NewMockDatabase(mockController)

	mockExtractor.SaveToBase()
}
