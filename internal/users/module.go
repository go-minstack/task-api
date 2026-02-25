package users

import (
	"github.com/go-minstack/core"
	user_repos "task-api/internal/users/repositories"
)

func Register(app *core.App) {
	app.Provide(user_repos.NewUserRepository)
	app.Provide(NewUserService)
	app.Provide(NewUserController)
	app.Invoke(RegisterRoutes)
}
