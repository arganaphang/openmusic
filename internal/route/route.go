package route

import (
	"github.com/arganaphang/openmusic/internal/service"
	"github.com/gin-gonic/gin"
)

func New(engine *gin.Engine, services *service.Services) {
	// Album Route
	NewAlbumRoute(engine, services)
	// Song Route
	NewSongRoute(engine, services)
	// Playlist Route
	NewPlaylistRoute(engine, services)
	// User Route
	NewUserRoute(engine, services)
}
