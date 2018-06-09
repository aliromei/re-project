package handlers

import (
  "github.com/kataras/iris"
  "github.com/aliromei/re-project/model"
  "fmt"
  "gopkg.in/go-playground/validator.v9"
  "github.com/aliromei/re-project/helpers"
)

type (
  update struct {
    Name     string `json:"name" validate:"required"`
    Email    string `json:"email" validate:"required,email"`
    Password string `json:"password" validate:"required"`
  }

  reserve struct {
    BusId string  `json:"busId" validate:"required"`
    Seat  int     `json:"seat" validate:"required"`
  }
)

func Profile(ctx iris.Context) {
  user, err := model.ShowUser()
  if err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
    return
  }
  ctx.JSON(iris.Map{"code":iris.StatusOK, "data":user})
}

func UpdateProfile(ctx iris.Context) {
  var data update
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
      Name: data.Name,
      Email: data.Email,
      PlainPassword: data.Password,
    }
    if err := user.Update(); err != nil {
      ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
      return
    }
    ctx.JSON(iris.Map{"code":iris.StatusOK, "data":user})
  }
}

func Reserve(ctx iris.Context) {
  var data reserve
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
    if err := model.Reserve(data.BusId, data.Seat); err != nil {
      ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
      return
    }
    ctx.JSON(iris.Map{"code":iris.StatusOK})
  }
}