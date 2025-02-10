package route

import (
	"github.com/arganaphang/openmusic/internal/service"
	"github.com/gin-gonic/gin"
)

type PlaylistRoute interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type playlistRoute struct {
	Services service.Services
}

func NewPlaylistRoute(engine *gin.Engine, services service.Services) PlaylistRoute {
	route := &playlistRoute{
		Services: services,
	}

	group := engine.Group("/playlists")
	group.GET("/", route.GetAll)
	group.GET("/:id", route.GetByID)
	group.POST("/", route.Create)
	group.PUT("/:id", route.Update)
	group.DELETE("/:id", route.Delete)

	return route
}

func (r playlistRoute) GetAll(c *gin.Context) {}

func (r playlistRoute) GetByID(c *gin.Context) {}

func (r playlistRoute) Create(c *gin.Context) {}

func (r playlistRoute) Update(c *gin.Context) {}

func (r playlistRoute) Delete(c *gin.Context) {}
