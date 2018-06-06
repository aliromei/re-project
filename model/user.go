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
  token, err := authentication.GenerateJWT(string(this.Id.Hex()), false)
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

func (this *User) Login() error {
  USER := connect("users")
  newUser := new(User)
  if err := USER.Find(bson.M{"email":this.Email}).One(&newUser); err != nil {
    return err
  } else {
    if _, err := authentication.ValidatePassword(string(this.Password), newUser.Password); err != nil {
      return err
    } else {
      if JWT, err := authentication.GenerateJWT(string(newUser.Id), newUser.IsAdmin); err != nil {
        return err
      } else {
        newUser.Token = JWT
        this = newUser
        return USER.Update(bson.M{"_id":this.Id}, bson.M{"$set":bson.M{"token":this.Token,"updatedAt":time.Now()}})
      }
    }
  }
}

func Logout() error {
  USER := connect("users")
  fmt.Println(authentication.Id)
  return USER.Update(bson.M{"_id":bson.ObjectIdHex(authentication.Id)}, bson.M{"$set":bson.M{"token":"","updatedAt":time.Now()}})
}