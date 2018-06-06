package middlewares

import (
  "github.com/kataras/iris"
  "fmt"
)

func Authorization(ctx iris.Context) {
  if ctx.Request().Header.Get("Authorization") != "" {
    fmt.Println("have Authorization") // todo: check for user & get role of user
    role := "admin" // todo: save user.Role to role
    ctx.Values().Set("role", role)
  } else {
    ctx.JSON(iris.Map{"code":iris.StatusForbidden,"error":"Permission Denied"})
  }
}

func AdminOnly(ctx iris.Context) {

}
