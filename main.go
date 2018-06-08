package main

import (
  "github.com/kataras/iris"
  "github.com/kataras/iris/middleware/logger"
  "github.com/aliromei/re-project/handlers"
  "github.com/aliromei/re-project/middlewares"
  "os"
  "github.com/aliromei/re-project/seed"
  "github.com/aliromei/re-project/connection"
)

func main() {
  connection.Dial()
  defer connection.Disconnect()

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

  if len(os.Args) > 1 {
    if os.Args[len(os.Args) - 1] == "--seed" {
      seed.Run()
    }
  }

  app.Post("/register", handlers.Register)
  app.Post("/login", handlers.Login)

  authorized := app.Party("/", middlewares.Authorization)

  authorized.Get("/config", handlers.Config)
  authorized.Post("/logout", handlers.Logout)

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
