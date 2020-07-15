package youtubeservice

import (
	"errors"
	"os"
	"strings"

	"djuta.software/record-roulette-api/utils/httpclient"
	"djuta.software/record-roulette-api/utils/url"
)

const baseURL = "https://www.googleapis.com/youtube/v3"
const maxResults = "50"

// GetVideosForChannel returns a slice of VideoSearchResults
func GetVideosForChannel(channelID string) ([]VideoResult, error) {
	requestURL := createVideoSearchURL(channelID)
	var searchResponse VideoResponse
	err := httpclient.GetJSON(requestURL, &searchResponse)
	if err != nil {
		return nil, err
	}
	return searchResponse.Items, nil
}

// GetChannel returns a ChannelSearchResult with a given id or username
func GetChannel(channelID string, channelUsername string) (ChannelResult, error) {
	requestURL := createChannelURL(channelID, channelUsername)
	var response ChannelResponse
	err := httpclient.GetJSON(requestURL, &response)
	if err != nil || len(response.Items) == 0 {
		err = errors.New("No items")
		return ChannelResult{}, err
	}
	return response.Items[0], nil
}

func createVideoSearchURL(channelID string) string {
	playlistID := strings.Replace(channelID, "UC", "UU", 1)
	return url.New(
		baseURL,
	).AddPath(
		"/playlistItems",
	).AddQueryParam(
		"maxResults", maxResults,
	).AddQueryParam(
		"key", os.Getenv("YOUTUBE_API_KEY"),
	).AddQueryParam(
		"part", "snippet",
	).AddQueryParam(
		"playlistId", playlistID,
	).GetURL()
}

func createChannelURL(channelID string, channelUsername string) string {
	u := url.New(
		baseURL,
	).AddPath(
		"/channels",
	).AddQueryParam(
		"key", os.Getenv("YOUTUBE_API_KEY"),
	).AddQueryParam(
		"part", "snippet",
	)
	if channelID != "" {
		u.AddQueryParam("id", channelID)
	}
	if channelUsername != "" {
		u.AddQueryParam("forUsername", channelUsername)
	}
	return u.GetURL()
}
