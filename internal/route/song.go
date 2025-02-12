package route

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/arganaphang/openmusic/internal/dto"
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
	var params dto.SongGetAllRequest
	if err := c.BindQuery(&params); err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	songs, err := r.Services.SongService.GetAll(c, params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SongGetAllResponse{
		Status:  "success",
		Message: "get songs",
		Data:    dto.SongGetAllResponseData{Songs: songs},
	})
}

func (r songRoute) GetByID(c *gin.Context) {
	id := c.Param("id")
	song, err := r.Services.SongService.GetByID(c, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SongGetByIDResponse{
		Status:  "success",
		Message: "get song",
		Data:    dto.SongGetByIDResponseData{Song: song},
	})
}

func (r songRoute) Create(c *gin.Context) {
	var body dto.SongCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	song, err := r.Services.SongService.Create(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.SongCreateResponse{
		Status:  "success",
		Message: "song created",
		Data:    dto.SongCreateResponseData{Song: song},
	})
}

func (r songRoute) Update(c *gin.Context) {
	id := c.Param("id")
	var body dto.SongUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	_, err := r.Services.SongService.GetByID(c, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	song, err := r.Services.SongService.Update(c, id, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(
		http.StatusOK,
		dto.SongUpdateResponse{
			Status:  "success",
			Message: "song updated",
			Data:    dto.SongUpdateResponseData{Song: song},
		},
	)
}

func (r songRoute) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := r.Services.SongService.GetByID(c, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	err = r.Services.SongService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.SongDeleteResponse{Status: "success", Message: "song deleted"})
}
