package authentication

import (
  "github.com/dgrijalva/jwt-go"
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

func GenerateHashedPassword(password string) ([]byte, error) {
  return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func CompareHashedAndPassword(hashedPassword []byte, password []byte) error {
  return bcrypt.CompareHashAndPassword(hashedPassword, password)
}