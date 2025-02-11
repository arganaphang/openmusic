package route

import (
	"net/http"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/entity"
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
	Services *service.Services
}

func NewSongRoute(engine *gin.Engine, services *service.Services) SongRoute {
	route := &songRoute{
		Services: services,
	}

	group := engine.Group("/songs")
	group.GET("", route.GetAll)
	group.GET(":id", route.GetByID)
	group.POST("", route.Create)
	group.PUT(":id", route.Update)
	group.DELETE(":id", route.Delete)

	return route
}

func (r songRoute) GetAll(c *gin.Context) {
	songs, err := r.Services.SongService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SongGetAllResponse{Success: true, Message: "get songs", Data: songs})
}

func (r songRoute) GetByID(c *gin.Context) {
	id := c.Param("id")
	song, err := r.Services.SongService.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SongGetByIDResponse{Success: true, Message: "get song", Data: song})
}

func (r songRoute) Create(c *gin.Context) {
	var body dto.SongCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	song, err := r.Services.SongService.Create(c, entity.Song{
		Title:     body.Title,
		Year:      body.Year,
		Genre:     body.Genre,
		Performer: body.Performer,
		Duration:  body.Duration,
		AlbumID:   body.AlbumID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SongCreateResponse{Success: true, Message: "song created", Data: song})
}

func (r songRoute) Update(c *gin.Context) {
	id := c.Param("id")
	var body dto.SongUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	song, err := r.Services.SongService.Update(c, id, entity.Song{
		Title:     body.Title,
		Year:      body.Year,
		Genre:     body.Genre,
		Performer: body.Performer,
		Duration:  body.Duration,
		AlbumID:   body.AlbumID,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SongUpdateResponse{Success: true, Message: "song updated", Data: song})
}

func (r songRoute) Delete(c *gin.Context) {
	id := c.Param("id")
	err := r.Services.SongService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SongDeleteResponse{Success: true, Message: "song deleted"})
}
