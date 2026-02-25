package auth_domain

import "github.com/go-minstack/core"

func Register(app *core.App) {
	app.Provide(NewAuthService)
}

func RegisterService(app *core.App) {
	app.Provide(NewAuthController)
	app.Invoke(RegisterRoutes)
}
