package authn

import "github.com/go-minstack/core"

func Register(app *core.App) {
	app.Provide(NewAuthService)
	app.Provide(NewAuthController)
	app.Invoke(RegisterRoutes)
}
