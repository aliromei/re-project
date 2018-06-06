package middlewares

import (
  "github.com/kataras/iris"
  "github.com/aliromei/re-project/authentication"
  "strings"
  "fmt"
)

func Authorization(ctx iris.Context) {
  if JWT := ctx.Request().Header.Get("Authorization"); JWT != "" {
    Jwt := strings.Split(JWT, " ")
    authentication.DecodeJWT(Jwt[1])
    fmt.Println(authentication.Id)
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
