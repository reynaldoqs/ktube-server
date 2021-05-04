package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	service "github.com/reynaldoqs/ka/internal/core/ports/services"
)

// VideosController is instance of contrller for videos service
type VideosController struct {
	vserv service.VideoService
}

// NewVideosController creates a instance of video controller
func NewVideosController(vserv service.VideoService) *VideosController {
	return &VideosController{
		vserv: vserv,
	}
}

// GetVideosByQuery find a video
func (vc *VideosController) GetVideosByQuery(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	videos, err := vc.vserv.SearchKaraoke(req.URL.Query().Get("query"))
	if err != nil {
		fmt.Println(err)
		res.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(res).Encode(err.Error())
		return
	}

	res.WriteHeader(http.StatusOK)
	json.NewEncoder(res).Encode(videos)
}
