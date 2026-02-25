package authn

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-minstack/web"
	authn_dto "task-api/internal/authn/dto"
	"task-api/internal/users/dto"
)

type AuthController struct {
	service *AuthService
}

func NewAuthController(service *AuthService) *AuthController {
	return &AuthController{service: service}
}

func (c *AuthController) login(ctx *gin.Context) {
	var input dto.LoginDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewErrorDto(err))
		return
	}
	token, err := c.service.Login(input)
	if err != nil {
		ctx.JSON(http.StatusUnauthorized, web.NewErrorDto(err))
		return
	}
	ctx.JSON(http.StatusOK, authn_dto.TokenDto{Token: token})
}
