package videos

import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/reynaldoqs/ka/internal/core/domain"
	repository "github.com/reynaldoqs/ka/internal/core/ports/repositories"

)

// Service is a instance for videos service
type Service struct {
	videosRepo repository.VideoRepository
}

// NewService creates a instance of a service
func NewService(vr repository.VideoRepository) *Service {
	return &Service{
		videosRepo: vr,
	}
}

// SearchKaraoke query a karaoke video
func (s *Service) SearchKaraoke(query string) ([]*domain.Video, error) {

	// add keywork karaoke
	keyword := ""
	query = fmt.Sprintf("%s %s", query, keyword)
	videos, err := s.videosRepo.SearchVideo(query)
	if err != nil {
		err = errors.Wrap(err, "videos.SearchKaraoke")
	}

	return videos, nil

}
