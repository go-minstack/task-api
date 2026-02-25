package auth_domain

import "github.com/gin-gonic/gin"

func RegisterRoutes(r *gin.Engine, c *AuthController) {
	g := r.Group("/api/auth")
	g.POST("/login", c.login)
}
