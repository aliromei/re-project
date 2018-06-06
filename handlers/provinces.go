package handlers

import "github.com/kataras/iris"

func Provinces(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"provinces"})
}
