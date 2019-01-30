package main

import "fmt"

func main() {
  var edad int
  fmt.Print ("Ingresa tu edad:")
  fmt.Scanf("%d",&edad)
  revision(edad)

  if variable := 2 ; variable == 3 {
    fmt.Println("Se cumple, lo cual no lo creo")
  }
}

func revision(edad int) {
    if edad >= 18 {
      fmt.Println("Eres mayor de edad");
    } else {
      fmt.Println("Eres menor de edad");
    }
}
