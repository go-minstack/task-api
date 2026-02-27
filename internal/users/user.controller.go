package users

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-minstack/auth"
	"github.com/go-minstack/web"
	"task-api/internal/users/dto"
)

type UserController struct {
	service *UserService
}

func NewUserController(service *UserService) *UserController {
	return &UserController{service: service}
}

func (c *UserController) register(ctx *gin.Context) {
	var input user_dto.RegisterDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewErrorDto(err))
		return
	}
	user, err := c.service.Register(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewErrorDto(err))
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) me(ctx *gin.Context) {
	claims, _ := auth.ClaimsFromContext(ctx)
	id, _ := strconv.ParseUint(claims.Subject, 10, 64)
	user, err := c.service.Me(uint(id))
	if err != nil {
		ctx.JSON(http.StatusNotFound, web.NewErrorDto(err))
		return
	}
	ctx.JSON(http.StatusOK, user)
}
