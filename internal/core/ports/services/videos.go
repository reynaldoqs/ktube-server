package service

import "github.com/reynaldoqs/ka/internal/core/domain"

// VideoService represents a service
type VideoService interface {
	SearchKaraoke(query string) ([]*domain.Video, error)
}
