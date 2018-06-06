package middlewares

import (
  "github.com/kataras/iris"
  "github.com/aliromei/re-project/authentication"
)

func Authorization(ctx iris.Context) {
  if JWT := ctx.Request().Header.Get("Authorization"); JWT != "" {
    authentication.DecodeJWT(JWT)
    ctx.Next()
  } else {
    ctx.JSON(iris.Map{"code":iris.StatusForbidden,"error":"Permission Denied"})
  }
}

func AdminOnly(ctx iris.Context) {
  if authentication.IsAdmin {
    ctx.Next()
  } else {
    ctx.JSON(iris.Map{"code":iris.StatusForbidden,"error":"Permission Denied"})
  }
}
