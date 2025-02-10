package route

import (
	"github.com/arganaphang/openmusic/internal/service"
	"github.com/gin-gonic/gin"
)

type SongRoute interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type songRoute struct {
	Services service.Services
}

func NewSongRoute(engine *gin.Engine, services service.Services) SongRoute {
	route := &songRoute{
		Services: services,
	}

	group := engine.Group("/songs")
	group.GET("/", route.GetAll)
	group.GET("/:id", route.GetByID)
	group.POST("/", route.Create)
	group.PUT("/:id", route.Update)
	group.DELETE("/:id", route.Delete)

	return route
}

func (r songRoute) GetAll(c *gin.Context) {}

func (r songRoute) GetByID(c *gin.Context) {}

func (r songRoute) Create(c *gin.Context) {}

func (r songRoute) Update(c *gin.Context) {}

func (r songRoute) Delete(c *gin.Context) {}
