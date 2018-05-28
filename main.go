package main

import (
	"os"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/aliromei/re-project/seed"
	"github.com/aliromei/re-project/connection"
	"github.com/aliromei/re-project/handlers"
	"github.com/aliromei/re-project/middlewares"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[len(os.Args) - 1] == "--seed" {
			seed.Run()
		}
	}

	app := iris.New()

	customLogger := logger.New(logger.Config{
		Status: true,
		IP: true,
		Method: true,
		Path: true,
		MessageContextKeys: []string{"logger_message"},
		MessageHeaderKeys: []string{"User-Agent"},
	})

	app.Use(customLogger)
	app.Use(middlewares.JsonOnly)

	connection.Dial()
	defer connection.Disconnect()

	app.Post("/register", handlers.Register)
	app.Post("/login", handlers.Login)
	app.Post("/logout", handlers.Logout)

	app.Post("/config", handlers.Config)

	authorized := app.Party("/", middlewares.Authorization)

	app.Get("/provinces", handlers.Provinces)

	user := authorized.Party("profile")
	user.Get("/", handlers.Profile)
	user.Post("/update", handlers.UpdateProfile)

	users := authorized.Party("/users", middlewares.AdminOnly)
	users.Get("/", handlers.UsersList)
	users.Get("/{id:string}", handlers.ShowUser)
	users.Post("/create", handlers.CreateUser)

	buses := authorized.Party("/buses")
	buses.Get("/", handlers.BusesList)
	buses.Get("/{id:string}", handlers.ShowBus)
	busesAdmin := buses.Party("/", middlewares.AdminOnly)
	busesAdmin.Post("/create", handlers.CreateBus)
	busesAdmin.Post("/{id:string}/update", handlers.UpdateBus)

	app.Run(iris.Addr("localhost:3333"))
}
