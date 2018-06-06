package model

import (
  "github.com/aliromei/re-project/authentication"
  "gopkg.in/mgo.v2/bson"
  "errors"
  "time"
)

func (this *User) Create() error {
  if err := this.uniqueEmailCheck(); err != nil {
    return err
  }
  password, err := authentication.GenerateHashedPassword(this.PlainPassword)
  if err != nil {
    return err
  }
  this.Password = string(password)

  this.Id = bson.NewObjectId()
  token, err := authentication.GenerateJWT(string(this.Id.Hex()), false)
  if err != nil {
    return err
  }
  this.Token = token

  this.CreatedAt, this.UpdatedAt = time.Now(), time.Now()

  this.Insert()

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
    if err := authentication.CompareHashedAndPassword([]byte(newUser.Password), []byte(this.PlainPassword)); err != nil {
      return err
    } else {
      if JWT, err := authentication.GenerateJWT(string(newUser.Id.Hex()), newUser.IsAdmin); err != nil {
        return err
      } else {
        newUser.Token = JWT
        if err := USER.Update(bson.M{"_id":newUser.Id}, bson.M{"$set":bson.M{"token":newUser.Token,"updatedAt":time.Now()}}); err != nil {
          return err
        } else {
          if err := USER.Find(bson.M{"_id":newUser.Id}).One(&this); err != nil {
            return err
          } else {
            return nil
          }
        }
      }
    }
  }
}

func Logout() error {
  USER := connect("users")
  return USER.Update(bson.M{"_id":bson.ObjectIdHex(authentication.Id)}, bson.M{"$set":bson.M{"token":"","updatedAt":time.Now()}})
}