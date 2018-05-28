package handlers

import "github.com/kataras/iris"

func Register(ctx iris.Context) {
	ctx.JSON(iris.Map{"message":"signup"})
}

func Login(ctx iris.Context) {
	ctx.JSON(iris.Map{"message":"login"})
}

func Logout(ctx iris.Context) {
	ctx.JSON(iris.Map{"message":"logout"})
}
