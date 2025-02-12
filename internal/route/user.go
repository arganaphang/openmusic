package route

import (
	"database/sql"
	"errors"
	"net/http"

	"github.com/arganaphang/openmusic/internal/dto"
	"github.com/arganaphang/openmusic/internal/service"
	"github.com/arganaphang/openmusic/pkg"
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
	RefreshToken(c *gin.Context)
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
		engine.PUT("/refresh-token", route.RefreshToken)
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
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserGetAllResponse{
		Status:  "success",
		Message: "get users",
		Data:    dto.UserGetAllResponseData{Users: users},
	})
}

func (r userRoute) GetByID(c *gin.Context) {
	id := c.Param("id")
	user, err := r.Services.UserService.GetByID(c, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserGetByIDResponse{
		Status:  "success",
		Message: "get user",
		Data:    dto.UserGetByIDResponseData{User: user},
	})
}

func (r userRoute) Create(c *gin.Context) {
	var body dto.UserCreateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	userExist, err := r.Services.UserService.GetByUsername(c, body.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	if userExist != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: "user already exist"})
		return
	}
	user, err := r.Services.UserService.Create(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.UserCreateResponse{
		Status:  "success",
		Message: "user created",
		Data:    dto.UserCreateResponseData{User: user},
	})
}

func (r userRoute) Update(c *gin.Context) {
	id := c.Param("id")
	var body dto.UserUpdateRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	user, err := r.Services.UserService.Update(c, id, body)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserUpdateResponse{
		Status:  "success",
		Message: "user updated",
		Data:    dto.UserUpdateResponseData{User: user},
	})
}

func (r userRoute) Delete(c *gin.Context) {
	id := c.Param("id")
	_, err := r.Services.UserService.GetByID(c, id)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusNotFound, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	err = r.Services.UserService.Delete(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.UserDeleteResponse{Status: "success", Message: "delete user"})
}

func (r userRoute) Register(c *gin.Context) {
	var body dto.RegisterRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	userExist, err := r.Services.UserService.GetByUsername(c, body.Username)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	if userExist != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: "user already exist"})
		return
	}
	token, err := r.Services.UserService.Register(c, body)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.RegisterResponse{
		Status:  "success",
		Message: "register success",
		Data:    dto.RegisterResponseData{Token: token},
	})
}

func (r userRoute) Login(c *gin.Context) {
	var body dto.LoginRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	_, err := r.Services.UserService.GetByUsername(c, body.Username)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	token, err := r.Services.UserService.Login(c, body)
	if err != nil && errors.Is(err, pkg.ErrUnauthorized) {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, dto.LoginResponse{
		Status:  "success",
		Message: "login success",
		Data:    dto.LoginResponseData{Token: token},
	})
}

func (r userRoute) RefreshToken(c *gin.Context) {
	var body dto.RefreshTokenRequest
	if err := c.ShouldBindJSON(&body); err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	token, err := r.Services.UserService.RefreshToken(c, body)
	if err != nil && errors.Is(err, pkg.ErrUnauthorized) {
		c.JSON(http.StatusUnauthorized, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.CommonResponse{Status: "fail", Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, dto.RefreshTokenResponse{
		Status:  "success",
		Message: "refresh token success",
		Data:    dto.RefreshTokenResponseData{Token: token},
	})
}
