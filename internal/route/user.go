package route

import (
	"net/http"
	"strconv"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/entity"
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
	Services *service.Services
}

func NewUserRoute(engine *gin.Engine, services *service.Services) UserRoute {
	route := &userRoute{
		Services: services,
	}
	{
		engine.POST("/register", route.Register)
		engine.POST("/login", route.Login)
	}
	group := engine.Group("/users")
	{
		group.GET("", route.GetAll)
		group.GET(":id", route.GetByID)
		group.POST("", route.Create)
		group.PUT(":id", route.Update)
		group.DELETE(":id", route.Delete)
	}

	return route
}

func (r userRoute) GetAll(c *gin.Context) {
	users, err := r.Services.UserService.GetAll(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserGetAllResponse{Success: false, Message: "get users", Data: users})
}

func (r userRoute) GetByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	user, err := r.Services.UserService.GetByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserGetByIDResponse{Success: false, Message: "get user", Data: user})
}

func (r userRoute) Create(c *gin.Context) {
	var body dto.UserCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	user, err := r.Services.UserService.Create(c, entity.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.UserCreateResponse{Success: true, Message: "user created", Data: user})
}

func (r userRoute) Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	var body dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	user, err := r.Services.UserService.Update(c, id, entity.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserUpdateResponse{Success: true, Message: "user updated", Data: user})
}

func (r userRoute) Delete(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	if err := r.Services.UserService.Delete(c, id); err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserDeleteResponse{Success: false, Message: "delete user"})
}

func (r userRoute) Register(c *gin.Context) {
	var body dto.RegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	token, err := r.Services.UserService.Register(c, entity.User{
		Name:     body.Name,
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.RegisterResponse{Success: true, Message: "register success", Data: token})
}

func (r userRoute) Login(c *gin.Context) {
	var body dto.RegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	token, err := r.Services.UserService.Login(c, entity.User{
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Success: false, Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.RegisterResponse{Success: true, Message: "login success", Data: token})
}
