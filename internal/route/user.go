package route

import (
	"github.com/arganaphang/openmusic/internal/service"
	"github.com/gin-gonic/gin"
)

type UserRoute interface {
	GetAll(c *gin.Context)
	GetByID(c *gin.Context)
	Create(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	Register(c *gin.Context)
	Login(c *gin.Context)
}

type userRoute struct {
	Services service.Services
}

func NewUserRoute(engine *gin.Engine, services service.Services) UserRoute {
	route := &userRoute{
		Services: services,
	}

	group := engine.Group("/users")
	group.GET("/", route.GetAll)
	group.GET("/:id", route.GetByID)
	group.POST("/", route.Create)
	group.PUT("/:id", route.Update)
	group.DELETE("/:id", route.Delete)

	return route
}

func (r userRoute) GetAll(c *gin.Context) {}

func (r userRoute) GetByID(c *gin.Context) {}

func (r userRoute) Create(c *gin.Context) {}

func (r userRoute) Update(c *gin.Context) {}

func (r userRoute) Delete(c *gin.Context) {}

func (r userRoute) Register(c *gin.Context) {}

func (r userRoute) Login(c *gin.Context) {}
