package main

import (
  "fmt"
  "io/ioutil"
  "strings"
  "../../../htmllinkparser"
)

func main() {
  bypeFileContent, err := ioutil.ReadFile("ex2.html")
  if err != nil {
    fmt.Print(err)
  }

  fileContent := string(bypeFileContent)

  fileReader := strings.NewReader(fileContent)
  links, err := htmllinkparser.Parse(fileReader)
  if err != nil {
    panic(err)
  }

  fmt.Printf("file content:\n%s\n", fileContent)
  fmt.Printf("links = %+v\n", links)
}
