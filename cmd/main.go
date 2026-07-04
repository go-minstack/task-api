package main

import (
	"github.com/go-minstack/go-minstack/auth"
	"github.com/go-minstack/go-minstack/core"
	mgin "github.com/go-minstack/go-minstack/gin"
	"github.com/go-minstack/go-minstack/sqlite"
	"gorm.io/gorm"
	"task-api/internal/authn"
	"task-api/internal/tasks"
	task_entities "task-api/internal/tasks/entities"
	"task-api/internal/users"
	user_entities "task-api/internal/users/entities"
)

func migrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&user_entities.User{},
		&task_entities.Task{},
	)
}

func main() {
	app := core.New(mgin.Module(), sqlite.Module(), auth.Module())

	users.Register(app)
	authn.Register(app)
	tasks.Register(app)

	app.Invoke(migrate)
	app.Run()
}
