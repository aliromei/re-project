package handlers

import "github.com/kataras/iris"

func Config(ctx iris.Context) {
	ctx.JSON(iris.Map{"message":"config"})
}
