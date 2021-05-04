package domain

import "time"

// Video respresents a video data
type Video struct {
	Kind    string      `json:"kind"`
	Etag    string      `json:"etag"`
	ID      IdData      `json:"id"`
	Snippet SnippetData `json:"snippet"`
}

type IdData struct {
	Kind       string `json:"kind"`
	VideoID    string `json:"videoId"`
	ChannelID  string `json:"channelId"`
	PlaylistID string `json:"playlistId"`
}

type SnippetData struct {
	PublishedAt  time.Time      `json:"publishedAt"`
	ChannelID    string         `json:"channelId"`
	Title        string         `json:"title"`
	Description  string         `json:"description"`
	Thumbnails   ThumbnailsData `json:"thumbnails"`
	ChannelTitle string         `json:"channelTitle"`
}

type ThumbnailsData struct {
	Key KeyData `json:"key"`
}

type KeyData struct {
	URL    string `json:"url"`
	Width  int64  `json:"width"`
	Height int64  `json:"height"`
	Extra  string `json:"extra"`
}
