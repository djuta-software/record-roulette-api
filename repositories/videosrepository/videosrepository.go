package videosrepository

import (
	"djuta.software/record-roulette-api/services/youtubeservice"
)

// GetVideosForChannel returns a slice of VideoDTOS
func GetVideosForChannel(channelID string) ([]VideoDTO, error) {
	videos := []VideoDTO{}
	results, err := youtubeservice.GetVideosForChannel(channelID)
	if err != nil {
		return nil, err
	}
	for _, result := range results {
		video := VideoDTO{
			VideoID: result.Snippet.ResourceID.VideoID,
			Title:   result.Snippet.Title,
			Image:   result.Snippet.Thumbnails.Default.URL,
		}
		videos = append(videos, video)
	}
	return videos, nil
}

// GetChannel returns a ChannelDTO
func GetChannel(channelID string, channelUsername string) (ChannelDTO, error) {
	result, err := youtubeservice.GetChannel(channelID, channelUsername)
	if err != nil {
		return ChannelDTO{}, err
	}
	channel := ChannelDTO{
		ChannelID: result.ID,
		Title:     result.Snippet.Title,
		Image:     result.Snippet.Thumbnails.Default.URL,
	}
	return channel, nil
}
