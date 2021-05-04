package videorepository

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/pkg/errors"
	"github.com/reynaldoqs/ka/internal/core/domain"
	"google.golang.org/api/googleapi/transport"
	"google.golang.org/api/youtube/v3"

)

// YoutubeService instance of youtube services
type YoutubeService struct {
	serv *youtube.Service
}

// NewYoutubeService creates a instnace of youtube services
func NewYoutubeService(devKey string) *YoutubeService {
	client := &http.Client{
		Transport: &transport.APIKey{Key: devKey},
	}
	service, err := youtube.New(client)
	if err != nil {
		err = errors.Wrap(err, "youtube.NewYoutubeService")
	}
	return &YoutubeService{
		serv: service,
	}
}

// SearchVideo returns a list of videos
func (s *YoutubeService) SearchVideo(query string) ([]*domain.Video, error) {
	//fquery := flag.String("query", query, "Search term")
	//maxResults := flag.Int64("max-results", 4, "Max Youtube results")
	//flag.Parse()
	call := s.serv.Search.List([]string{"id", "snippet"}).Q(query).MaxResults(5).VideoType("any").Type("video")
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error making search API call: %v", err)
	}
	var data []*domain.Video
	for i, sr := range response.Items {

		newIdVideo := domain.IdData{
			ChannelID:  sr.Id.ChannelId,
			Kind:       sr.Id.Kind,
			PlaylistID: sr.Id.PlaylistId,
			VideoID:    sr.Id.VideoId,
		}
		publishedAt, _ := time.Parse("January 2, 2006", sr.Snippet.PublishedAt)
		newThumbnailsKey := domain.KeyData{
			URL:    sr.Snippet.Thumbnails.Default.Url,
			Height: sr.Snippet.Thumbnails.Default.Height,
			Width:  sr.Snippet.Thumbnails.Default.Width,
			Extra:  sr.Snippet.Thumbnails.High.Url,
		}
		newThumbnails := domain.ThumbnailsData{
			Key: newThumbnailsKey,
		}
		newSnippetData := domain.SnippetData{
			PublishedAt:  publishedAt,
			ChannelID:    sr.Snippet.ChannelId,
			Title:        sr.Snippet.Title,
			ChannelTitle: sr.Snippet.ChannelTitle,
			Description:  sr.Snippet.Description,
			Thumbnails:   newThumbnails,
		}
		newVideo := domain.Video{Etag: sr.Etag, ID: newIdVideo, Snippet: newSnippetData, Kind: sr.Kind}
		fmt.Print(i, newVideo)
		data = append(data, &newVideo)
	}
	println(response.Items)
	return data, nil
}
