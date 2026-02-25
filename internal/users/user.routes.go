package users

import (
	"github.com/gin-gonic/gin"
	"github.com/go-minstack/auth"
)

func RegisterRoutes(r *gin.Engine, c *UserController, jwt *auth.JwtService) {
	g := r.Group("/api/users")
	g.POST("/register", c.register)
	g.GET("/me", auth.Authenticate(jwt), c.me)
}
