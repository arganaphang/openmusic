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
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "get songs", "data": songs})
}

func (r songRoute) GetByID(c *gin.Context) {
	id := c.Param("id")
	song, err := r.Services.SongService.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "get song", "data": song})
}

func (r songRoute) Create(c *gin.Context) {
	var body dto.SongCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"success": true, "message": "song created", "data": song})
}

func (r songRoute) Update(c *gin.Context) {
	id := c.Param("id")
	var body dto.SongUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
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
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "song updated", "data": song})
}

func (r songRoute) Delete(c *gin.Context) {
	id := c.Param("id")
	err := r.Services.SongService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "message": "song deleted", "data": nil})
}
