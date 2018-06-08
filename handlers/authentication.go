package handlers

import (
  "github.com/kataras/iris"
  "gopkg.in/go-playground/validator.v9"
  "github.com/aliromei/re-project/helpers"
  "github.com/aliromei/re-project/model"
  "fmt"
)

type (
  register struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
  }
  login struct {
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
  }
)

func Register(ctx iris.Context) {
  var data register
  if err := ctx.ReadJSON(&data); err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":err})
    return
  } else if err := validate.Struct(data); err != nil {
    errs := make(map[string]string)
    for _, err := range err.(validator.ValidationErrors) {
      errs[helpers.LowerFirst(err.Field())] = err.Tag()
    }
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":errs})
    return
  } else {
    var user = model.User{
      Name:          data.Name,
      Email:         data.Email,
      PlainPassword: data.Password,
    }
    if err := user.Create(false); err != nil {
      ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
      return
    }
    ctx.JSON(iris.Map{"code":iris.StatusOK, "data":user})
  }
}

func Login(ctx iris.Context) {
  var data login
  if err := ctx.ReadJSON(&data); err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":err})
    return
  } else if err := validate.Struct(data); err != nil {
    errs := make(map[string]string)
    for _, err := range err.(validator.ValidationErrors) {
      errs[helpers.LowerFirst(err.Field())] = err.Tag()
    }
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":errs})
    return
  } else {
    var user = model.User{
      Email:         data.Email,
      PlainPassword: data.Password,
    }
    if err := user.Login(); err != nil {
      ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
      return
    }
    ctx.JSON(iris.Map{"code":iris.StatusOK, "data":user})
  }
}

func Logout(ctx iris.Context) {
  if err := model.Logout(); err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
    return
  }
  ctx.JSON(iris.Map{"code":iris.StatusOK})
}
