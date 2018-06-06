package model

import (
  "github.com/aliromei/re-project/authentication"
  "gopkg.in/mgo.v2/bson"
  "fmt"
  "errors"
  "time"
)

func (this *User) Create() error {
  if err := this.uniqueEmailCheck(); err != nil {
    return err
  }
  password, err := authentication.GeneratePassword(this.PlainPassword)
  if err != nil {
    return err
  }
  this.Password = password

  this.Id = bson.NewObjectId()
  token, err := authentication.GenerateJWT(string(this.Id), false)
  if err != nil {
    return err
  }
  this.Token = token

  this.CreatedAt, this.UpdatedAt = time.Now(), time.Now()

  this.Insert()
  fmt.Println(this.Id)

  return nil
}

func (this *User) Insert() error {
  USER := connect("users")
  err := USER.Insert(this)
  if err != nil {
    return err
  }
  USER.Find(bson.M{"id":this.Id}).One(this)

  return nil
}

func (this *User) uniqueEmailCheck() error  {
  USER := connect("users")
  count, err := USER.Find(bson.M{"email":this.Email}).Count()
  if err != nil {
    return err
  } else if count > 0 {
    return errors.New("we have a user with the email you entered")
  } else {
    return nil
  }
}