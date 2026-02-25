package tasks

import (
	"github.com/gin-gonic/gin"
	"github.com/go-minstack/auth"
)

func RegisterRoutes(r *gin.Engine, c *TaskController, jwt *auth.JwtService) {
	g := r.Group("/api/tasks", auth.Authenticate(jwt))
	g.GET("", c.list)
	g.POST("", c.create)
	g.GET("/:id", c.get)
	g.PATCH("/:id", c.update)
	g.DELETE("/:id", c.delete)
}
