package main

import "fmt"

func main() {
  var emptyStr string
  if ( emptyStr == "") {
    fmt.Println("empty string is not nil type")
  }

  fmt.Println("Way to initialize two dimonension map")

  twoDimon := make(map[string]interface{}, 0)
  if (twoDimon["data"] == nil) {
      twoDimon["data"] = make(map[string]string, 0)
  }
  twoDimon["data"].(map[string]string)["1"] = "001"
  fmt.Println("Initialize two dimonension map end")

}
