package middlewares

import (
  "github.com/kataras/iris"
)

func JsonOnly(ctx iris.Context) {
  if ctx.Request().Header.Get("Content-Type") == "application/json" {
    ctx.Next()
  } else {
    ctx.StatusCode(400)
    ctx.JSON(iris.Map{"code":400,"error":"Content-Type must be JSON"})
  }
}
