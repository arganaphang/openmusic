package route

import (
	"github.com/arganaphang/openmusic/internal/service"
	"github.com/gin-gonic/gin"
)

type AlbumRoute interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type albumRoute struct {
	Services service.Services
}

func NewAlbumRoute(engine *gin.Engine, services service.Services) AlbumRoute {
	route := &albumRoute{
		Services: services,
	}

	group := engine.Group("/albums")
	group.GET("/", route.GetAll)
	group.GET("/:id", route.GetByID)
	group.POST("/", route.Create)
	group.PUT("/:id", route.Update)
	group.DELETE("/:id", route.Delete)

	return route
}

func (r albumRoute) GetAll(c *gin.Context) {}

func (r albumRoute) GetByID(c *gin.Context) {}

func (r albumRoute) Create(c *gin.Context) {}

func (r albumRoute) Update(c *gin.Context) {}

func (r albumRoute) Delete(c *gin.Context) {}
