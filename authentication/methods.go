package authentication

import (
  "github.com/dgrijalva/jwt-go"
  "fmt"
  "golang.org/x/crypto/bcrypt"
)
func GenerateJWT(id string, isAdmin bool) (string, error) {
  claims["id"], claims["isAdmin"] = id, isAdmin
  Id, IsAdmin = id, isAdmin

  return token.SignedString(secret)
}

func DecodeJWT(JWT string) bool {
  claims := jwt.MapClaims{}
  _, err := jwt.ParseWithClaims(JWT, claims, func(token *jwt.Token) (interface{}, error) {
    return secret, nil
  })

  if err != nil {
    return false
  }

  fmt.Println(claims)

  for key, value := range claims {
    switch string(key) {
    case "id":
      Id = value.(string)
    case "isAdmin":
      IsAdmin = value.(bool)
    }
  }
  return true
}

func GeneratePassword(userPassword string) ([]byte, error) {
  return bcrypt.GenerateFromPassword([]byte(userPassword), bcrypt.DefaultCost)
}

func ValidatePassword(userPassword string, hashedPassword []byte) (bool, error) {
  if err := bcrypt.CompareHashAndPassword(hashedPassword, []byte(userPassword)); err != nil {
    return false, err
  }
  return true, nil
}