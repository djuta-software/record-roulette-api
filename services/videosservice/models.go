package videosservice

// Video Represents a video
type Video struct {
	VideoID string `json:"videoId"`
	Title   string `json:"title"`
	Image   string `json:"image"`
}

// Channel Represents a channel
type Channel struct {
	ChannelID string `json:"channelId"`
	Title     string `json:"title"`
	Image     string `json:"image"`
}
