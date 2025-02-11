package route

import (
	"net/http"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/entity"
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
	Services *service.Services
}

func NewAlbumRoute(engine *gin.Engine, services *service.Services) AlbumRoute {
	route := &albumRoute{
		Services: services,
	}

	group := engine.Group("/albums")
	{
		group.GET("", route.GetAll)
		group.GET(":id", route.GetByID)
		group.POST("", route.Create)
		group.PUT(":id", route.Update)
		group.DELETE(":id", route.Delete)
	}
	return route
}

func (r albumRoute) GetAll(c *gin.Context) {
	albums, err := r.Services.AlbumService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.AlbumGetAllResponse{Success: true, Message: "get albums", Data: albums})
}

func (r albumRoute) GetByID(c *gin.Context) {
	id := c.Param("id")
	album, err := r.Services.AlbumService.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.AlbumGetByIDResponse{Success: true, Message: "get album", Data: album})
}

func (r albumRoute) Create(c *gin.Context) {
	var data dto.AlbumCreateRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	album, err := r.Services.AlbumService.Create(c, entity.Album{
		Name: data.Name,
		Year: data.Year,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.AlbumCreateResponse{Success: true, Message: "create album", Data: album})
}

func (r albumRoute) Update(c *gin.Context) {
	id := c.Param("id")
	var data dto.AlbumUpdateRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	album, err := r.Services.AlbumService.Update(c, id, entity.Album{
		Name: data.Name,
		Year: data.Year,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.AlbumUpdateResponse{Success: true, Message: "update album", Data: album})
}

func (r albumRoute) Delete(c *gin.Context) {
	id := c.Param("id")
	err := r.Services.AlbumService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.AlbumDeleteResponse{Success: true, Message: "delete album"})
}
