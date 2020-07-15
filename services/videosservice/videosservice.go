package videosservice

import (
	"math/rand"
	"strings"

	"djuta.software/record-roulette-api/repositories/videosrepository"
	"djuta.software/record-roulette-api/utils/slice"
)

// GetRandomVideos returns an slice of Videos
func GetRandomVideos(channelIDs []string) ([]Video, error) {
	videos := []Video{}
	for _, channelID := range channelIDs {
		videosOut, errOut := getVideosForChannel(channelID)
		select {
		case err := <-errOut:
			return nil, err
		case channelVideos := <-videosOut:
			videos = append(videos, channelVideos...)
		}
		channelVideos := <-videosOut
		videos = append(videos, channelVideos...)
	}
	shuffleVideos(videos)
	return videos, nil
}

// GetChannel returns Channels
func GetChannel(channelURL string) (Channel, error) {
	channelURLFragments := strings.Split(channelURL, "/")
	channelID := channelURLFragments[len(channelURLFragments)-1]
	channelUsername := ""
	_, isChannelID := slice.Find(channelURLFragments, "channel")

	if !isChannelID {
		channelUsername = channelID
		channelID = ""
	}

	channelDTO, err := videosrepository.GetChannel(channelID, channelUsername)
	if err != nil {
		return Channel{}, err
	}
	channel := Channel{
		ChannelID: channelDTO.ChannelID,
		Title:     channelDTO.Title,
		Image:     channelDTO.Image,
	}
	return channel, nil
}

func getVideosForChannel(channelID string) (<-chan []Video, <-chan error) {
	videosChan := make(chan []Video)
	errChan := make(chan error)
	go func() {
		defer close(videosChan)
		defer close(errChan)
		videos := []Video{}
		videoDTOs, err := videosrepository.GetVideosForChannel(channelID)
		if err != nil {
			errChan <- err
			return
		}
		for _, videoDTO := range videoDTOs {
			video := Video{
				VideoID: videoDTO.VideoID,
				Title:   videoDTO.Title,
				Image:   videoDTO.Image,
			}
			videos = append(videos, video)
		}
		videosChan <- videos
	}()
	return videosChan, errChan
}

func shuffleVideos(slice []Video) {
	for i := 1; i < len(slice); i++ {
		r := rand.Intn(i + 1)
		if i != r {
			slice[r], slice[i] = slice[i], slice[r]
		}
	}
}
