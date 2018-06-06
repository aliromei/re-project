package handlers

import "github.com/kataras/iris"

func BusesList(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"buses list"})
}
func ShowBus(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"show bus"})
}
func CreateBus(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"create bus"})
}
func UpdateBus(ctx iris.Context) {
  ctx.JSON(iris.Map{"message":"update bus"})
}
