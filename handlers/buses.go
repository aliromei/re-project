package handlers

import (
  "github.com/kataras/iris"
  "gopkg.in/go-playground/validator.v9"
  "github.com/aliromei/re-project/helpers"
  "github.com/aliromei/re-project/model"
  "fmt"
)

type (
  Origin struct {
    ProvinceId int `json:"provinceId,number" validate:"required"`
    CityId     int `json:"cityId,number" validate:"required"`
  }

  Destination struct {
    ProvinceId int `json:"provinceId,number" validate:"required"`
    CityId     int `json:"cityId,number" validate:"required"`
  }

  Bus struct {
    Model       string      `json:"model,number" validate:"required"`
    Seats       int         `json:"seats,number" validate:"required"`
    Origin      Origin      `json:"origin,number" validate:"required"`
    Destination Destination `json:"destination,number" validate:"required"`
  }

  UpdateBusStatus struct {
    Status int `json:"status,number" validate:"required"`
  }
)

func CreateBus(ctx iris.Context) {
  var data Bus
  if err := ctx.ReadJSON(&data); err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":err})
    return
  } else if err := validate.Struct(data); err != nil {
    errs := make(map[string]string)
    for _, err := range err.(validator.ValidationErrors) {
      fmt.Println(err)
      errs[helpers.LowerFirst(err.Field())] = err.Tag()
    }
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":errs})
    return
  } else {
    var Origin = model.Address{
      ProvinceId: data.Origin.ProvinceId,
      CityId:     data.Origin.CityId,
    }
    var Destination = model.Address{
      ProvinceId: data.Destination.ProvinceId,
      CityId:     data.Destination.CityId,
    }
    var bus = model.Bus{
      Model: data.Model,
      Seats: data.Seats,
      Origin: Origin,
      Destination: Destination,
    }
    if err := bus.Create(); err != nil {
      ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
      return
    }
    ctx.JSON(iris.Map{"code":iris.StatusOK, "data":bus})
  }
}

func UpdateBus(ctx iris.Context) {
  var data UpdateBusStatus
  if err := ctx.ReadJSON(&data); err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":err})
    return
  } else if err := validate.Struct(data); err != nil {
    errs := make(map[string]string)
    for _, err := range err.(validator.ValidationErrors) {
      fmt.Println(err)
      errs[helpers.LowerFirst(err.Field())] = err.Tag()
    }
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":errs})
    return
  } else {
    bus, err := model.UpdateBus(ctx.Params().Get("id"), data.Status)
    if err != nil {
      ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
      return
    }
    ctx.JSON(iris.Map{"code":iris.StatusOK, "data":bus})
  }
}

func ShowBus(ctx iris.Context) {
  var bus model.Bus
  if err := bus.ShowBus(ctx.Params().Get("id")); err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
  }
  ctx.JSON(iris.Map{"code":iris.StatusOK, "data":bus})
}

func BusesList(ctx iris.Context) {
  buses, err := model.BusesList()
  if err != nil {
    ctx.JSON(iris.Map{"code":iris.StatusBadRequest, "errors":fmt.Sprintf("%v", err)})
  }
  ctx.JSON(iris.Map{"code":iris.StatusOK, "data":buses})
}