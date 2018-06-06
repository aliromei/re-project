package handlers

import "github.com/kataras/iris"

func Profile(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"profile"})
}

func UpdateProfile(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"update profile"})
}