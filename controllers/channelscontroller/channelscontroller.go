package channelscontroller

import (
	"net/http"

	"djuta.software/record-roulette-api/services/videosservice"
	"djuta.software/record-roulette-api/utils/api"
)

// GetChannel returns an array of Channels
func GetChannel(writer http.ResponseWriter, request *http.Request) {
	channelURL := request.URL.Query().Get("channelUrl")
	if channelURL == "" {
		api.SendBadRequestResponse(writer, "channelUrl is required")
		return
	}
	channel, err := videosservice.GetChannel(channelURL)
	if err != nil {
		api.SendInternalServerError(writer, err)
		return
	}
	api.SendOkResponse(writer, channel)
}
