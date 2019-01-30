package main

import "fmt"

func main() {
  hola := "hola"
  fmt.Println("Impresión de japones: もしもし!")
  fmt.Println("Ejemplo de impresión de caracter en ascii: ",hola[0])
  fmt.Println("Ejemplo de impresión de caracter en texto: ",string(hola[0]))
  fmt.Println("Ejemplo de tamaño de texto: ",len(hola))
}
