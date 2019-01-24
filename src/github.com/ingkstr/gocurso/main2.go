package main

import "fmt"

func main() {
  var name string
  fmt.Println ("Ingresa tu nombre:")
  fmt.Scanf("%s",&name)
  fmt.Printf("Welcome to Go %s!!!",name)
}
