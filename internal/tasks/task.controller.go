package tasks

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-minstack/auth"
	"github.com/go-minstack/web"
	"task-api/internal/tasks/dto"
)

type TaskController struct {
	service *TaskService
}

func NewTaskController(service *TaskService) *TaskController {
	return &TaskController{service: service}
}

func (c *TaskController) list(ctx *gin.Context) {
	claims, _ := auth.ClaimsFromContext(ctx)
	tasks, err := c.service.List(claims)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewErrorDto(err))
		return
	}
	ctx.JSON(http.StatusOK, tasks)
}

func (c *TaskController) create(ctx *gin.Context) {
	var input dto.CreateTaskDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewErrorDto(err))
		return
	}
	claims, _ := auth.ClaimsFromContext(ctx)
	task, err := c.service.Create(claims, input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, web.NewErrorDto(err))
		return
	}
	ctx.JSON(http.StatusCreated, task)
}

func (c *TaskController) get(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewErrorDto(err))
		return
	}
	claims, _ := auth.ClaimsFromContext(ctx)
	task, err := c.service.Get(claims, id)
	if err != nil {
		status := http.StatusNotFound
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}
		ctx.JSON(status, web.NewErrorDto(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) update(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewErrorDto(err))
		return
	}
	var input dto.UpdateTaskDto
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewErrorDto(err))
		return
	}
	claims, _ := auth.ClaimsFromContext(ctx)
	task, err := c.service.Update(claims, id, input)
	if err != nil {
		status := http.StatusInternalServerError
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}
		ctx.JSON(status, web.NewErrorDto(err))
		return
	}
	ctx.JSON(http.StatusOK, task)
}

func (c *TaskController) delete(ctx *gin.Context) {
	id, err := parseID(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, web.NewErrorDto(err))
		return
	}
	claims, _ := auth.ClaimsFromContext(ctx)
	if err := c.service.Delete(claims, id); err != nil {
		status := http.StatusNotFound
		if err.Error() == "forbidden" {
			status = http.StatusForbidden
		}
		ctx.JSON(status, web.NewErrorDto(err))
		return
	}
	ctx.Status(http.StatusNoContent)
}

func parseID(ctx *gin.Context) (uint, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 10, 64)
	if err != nil {
		return 0, errors.New("invalid id")
	}
	return uint(id), nil
}
