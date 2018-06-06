package handlers

import "github.com/kataras/iris"

func UsersList(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"users list"})
}

func ShowUser(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"show user"})
}

func CreateUser(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"create user"})
}
