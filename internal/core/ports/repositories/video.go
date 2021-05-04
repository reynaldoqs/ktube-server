package repository

import "github.com/reynaldoqs/ka/internal/core/domain"

// VideoRepository define a repository
type VideoRepository interface {
	SearchVideo(query string) ([]*domain.Video, error)
}
