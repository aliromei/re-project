package authentication

import (
  "github.com/dgrijalva/jwt-go"
)

var (
  Id         string
  IsAdmin    bool
  secret = []byte("re-project-server")
  token = jwt.New(jwt.SigningMethodHS256)
  claims = token.Claims.(jwt.MapClaims)
)