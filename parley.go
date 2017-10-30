package main

import (
  "io/ioutil"
  "fmt"
  "encoding/base64"
  "os"
)

func main() {
  if len(os.Args) < 2 {
    fmt.Println("No filepath specified.")
    os.Exit(1)
  }
  path := os.Args[1]
  if path == "" {
    fmt.Println("No filepath specified.")
    os.Exit(1)
  }
  img, err := ioutil.ReadFile(os.Args[1])
  if err == nil {
    encodedData := base64.StdEncoding.EncodeToString(img)
    fmt.Println(fmt.Sprintf(`data:image/png;base64,%s`, encodedData))
  } else {
    fmt.Println(err)
    os.Exit(1)
  }
}
