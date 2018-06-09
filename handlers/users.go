package handlers

import (
  "github.com/kataras/iris"
  "fmt"
  "github.com/aliromei/re-project/model"
  "gopkg.in/go-playground/validator.v9"
  "github.com/aliromei/re-project/helpers"
)

type (
  create struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
    IsAdmin  bool   `json:"isAdmin" validate:"required"`
  }
)

func UsersList(ctx iris.Context) {
  users, err := model.UsersList()
  if err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
  }
  ctx.JSON(iris.Map{"code":iris.StatusOK, "data":users})
}

func ShowUser(ctx iris.Context) {
  user, err := model.ShowUserA(ctx.Params().Get("id"))
  if err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
    return
  }
  ctx.JSON(iris.Map{"code":iris.StatusOK, "data":user})
}

func CreateUser(ctx iris.Context) {
  var data create
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
      IsAdmin:       data.IsAdmin,
    }
    if err := user.Create(true); err != nil {
      ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
      return
    }
    ctx.JSON(iris.Map{"code":iris.StatusOK, "data":user})
  }
}
