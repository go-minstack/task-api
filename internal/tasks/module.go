package tasks

import (
	"github.com/go-minstack/core"
	task_repos "task-api/internal/tasks/repositories"
)

func Register(app *core.App) {
	app.Provide(task_repos.NewTaskRepository)
	app.Provide(NewTaskService)
}

func RegisterService(app *core.App) {
	app.Provide(NewTaskController)
	app.Invoke(RegisterRoutes)
}
