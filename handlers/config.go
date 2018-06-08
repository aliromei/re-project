package handlers

import (
  "github.com/kataras/iris"
  "github.com/aliromei/re-project/seed"
)

func Config(ctx iris.Context) {
  config, err := seed.Config()
  if err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":err})
    return
  }
  ctx.JSON(iris.Map{"code":iris.StatusOK, "data":config})
}
