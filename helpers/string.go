package helpers

import (
  "unicode"
)

func LowerFirst(str string) string {
  for i, v := range str {
    return string(unicode.ToLower(v)) + str[i+1:]
  }
  return ""
}