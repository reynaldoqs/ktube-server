package server

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	videos "github.com/reynaldoqs/ka/internal/core/service"
	"github.com/reynaldoqs/ka/internal/infrastructure/controller"
	videorepository "github.com/reynaldoqs/ka/internal/infrastructure/repositories"

)

const developerKey = "AIzaSyCuHwjTi26PHYP_v2qrBh9jRdgAeVzmXeg"

func RegisterRouter(port string) {
	chiDispatcher := chi.NewRouter()
	chiDispatcher.Use(middleware.RequestID)
	chiDispatcher.Use(middleware.RealIP)
	chiDispatcher.Use(middleware.Logger)
	chiDispatcher.Use(middleware.Recoverer)

	chiDispatcher.Use(cors.Handler(cors.Options{
		// AllowedOrigins: []string{"https://foo.com"}, // Use this to allow specific origin hosts
		AllowedOrigins:   []string{"*"},
		// AllowOriginFunc:  func(r *http.Request, origin string) bool { return true },
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	  }))

	youtubeRepo := videorepository.NewYoutubeService(developerKey)
	videosService := videos.NewService(youtubeRepo)

	videosController := controller.NewVideosController(videosService)

	chiDispatcher.Get("/videos/search", videosController.GetVideosByQuery)

	fmt.Printf("Chi HTTP server running on port %v\n", port)
	http.ListenAndServe(port, chiDispatcher)

}
