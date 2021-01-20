package urls

import (
	"testing"
	"strings"

	"github.com/stretchr/testify/require"
)

func TestGetLinkPreviewData(t *testing.T) {

	statusTownhall := LinkPreviewData{
		Site:         "YouTube",
		Title:        "Status Town Hall #67 - 12 October 2020",
		ThumbnailURL: "https://i.ytimg.com/vi/mzOyYtfXkb0/hqdefault.jpg",
	}

	previewData, err := GetLinkPreviewData("https://www.youtube.com/watch?v=mzOyYtfXkb0")
	require.NoError(t, err)
	require.Equal(t, statusTownhall.Site, previewData.Site)
	require.Equal(t, statusTownhall.Title, previewData.Title)
	require.Equal(t, statusTownhall.ThumbnailURL, previewData.ThumbnailURL)

	previewData, err = GetLinkPreviewData("https://youtu.be/mzOyYtfXkb0")
	require.NoError(t, err)
	require.Equal(t, statusTownhall.Site, previewData.Site)
	require.Equal(t, statusTownhall.Title, previewData.Title)
	require.Equal(t, statusTownhall.ThumbnailURL, previewData.ThumbnailURL)

	_, err = GetLinkPreviewData("https://www.test.com/unknown")
	require.Error(t, err)

}

func TestGetGiphyPreviewData(t *testing.T) {
	validGiphyLink := "https://giphy.com/gifs/FullMag-robot-boston-dynamics-dance-lcG3qwtTKSNI2i5vst"
	previewData, err := GetGiphyPreviewData(validGiphyLink)
	bostonDynamicsEthGifData := LinkPreviewData{
		Site:         "GIPHY",
		Title:        "Boston Dynamics Yes GIF by FullMag - Find & Share on GIPHY",
		ThumbnailURL: "https://media1.giphy.com/media/lcG3qwtTKSNI2i5vst/giphy.gif",
	}
	require.NoError(t, err)
	require.Equal(t, bostonDynamicsEthGifData.Site, previewData.Site)
	require.Equal(t, bostonDynamicsEthGifData.Title, previewData.Title)

	// Giphy oembed returns links to different servers: https://media1.giphy.com, https://media2.giphy.com and so on
	// We don't care about the server as long as other parts are equal, so we split at "." and ignore the first item
	require.Equal(t, strings.Split(bostonDynamicsEthGifData.ThumbnailURL, ".")[1:], strings.Split(previewData.ThumbnailURL, ".")[1:])

	invalidGiphyLink := "https://giphy.com/gifs/this-gif-does-not-exist-44444"
	_, err = GetGiphyPreviewData(invalidGiphyLink)
	require.Error(t, err)
}

func TestGetTenorPreviewData(t *testing.T) {
	validTenorLink := "https://tenor.com/view/robot-dance-do-you-love-me-boston-boston-dynamics-dance-gif-19998728"
	previewData, err := GetTenorPreviewData(validTenorLink)

	gifData := LinkPreviewData{
		Site:         "Tenor",
		Title:        "Annihere",
		ThumbnailURL: "https://media.tenor.com/images/975f6b95d188c277ebba62d9b5511685/tenor.gif",
	}
	require.NoError(t, err)
	require.Equal(t, gifData.Site, previewData.Site)
	require.Equal(t, gifData.Title, previewData.Title)
	require.Equal(t, gifData.ThumbnailURL, previewData.ThumbnailURL)

	invalidTenorLink := "https://giphy.com/gifs/this-gif-does-not-exist-44444"
	_, err = GetTenorPreviewData(invalidTenorLink)
	require.Error(t, err)
}
