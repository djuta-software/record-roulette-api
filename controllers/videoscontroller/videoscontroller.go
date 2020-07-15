package videoscontroller

import (
	"net/http"
	"strings"

	"djuta.software/record-roulette-api/services/videosservice"
	"djuta.software/record-roulette-api/utils/api"
)

// GetRandomVideos returns an array of Videos
func GetRandomVideos(writer http.ResponseWriter, request *http.Request) {
	maxChannels := 5
	channelIds := request.URL.Query().Get("channelIds")
	channels := strings.Split(channelIds, ",")
	if len(channels) == 1 && channels[0] == "" {
		api.SendBadRequestResponse(writer, "channelIds is required")
		return
	}
	if len(channels) > maxChannels {
		channels = channels[:5]
	}
	videos, err := videosservice.GetRandomVideos(channels)
	if err != nil {
		api.SendInternalServerError(writer, err)
		return
	}
	api.SendOkResponse(writer, videos)
}
