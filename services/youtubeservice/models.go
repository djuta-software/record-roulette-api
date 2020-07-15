package youtubeservice

// VideoResultID represents a search result id
type VideoResultID struct {
	VideoID string
}

// Image represents an image
type Image struct {
	URL string
}

// SnippetThumbnails represents thumbnails
type SnippetThumbnails struct {
	Default Image
}

// Snippet represents a snippet
type Snippet struct {
	ResourceID VideoResultID
	Title      string
	Thumbnails SnippetThumbnails
}

// VideoResult represents a search result
type VideoResult struct {
	Snippet Snippet
}

// VideoResponse represents a search response
type VideoResponse struct {
	Items []VideoResult
}

// ChannelResult represents a search result
type ChannelResult struct {
	ID      string
	Snippet Snippet
}

// ChannelResponse represents a search response
type ChannelResponse struct {
	Items []ChannelResult
}
