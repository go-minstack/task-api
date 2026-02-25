package main

import (
	"github.com/go-minstack/auth"
	"github.com/go-minstack/core"
	mgin "github.com/go-minstack/gin"
	"github.com/go-minstack/sqlite"
	auth_domain "task-api/internal/auth"
	"task-api/internal/tasks"
	task_entities "task-api/internal/tasks/entities"
	"task-api/internal/users"
	user_entities "task-api/internal/users/entities"
	"gorm.io/gorm"
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
	auth_domain.Register(app)
	tasks.Register(app)

	users.RegisterService(app)
	auth_domain.RegisterService(app)
	tasks.RegisterService(app)

	app.Invoke(migrate)
	app.Run()
}
