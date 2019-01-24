package main

import "fmt"

const HELLO_USER string = "Hola %s %s, Bienvenido al fascinante mundo de Go\n"

func main() {
  var name string
  lastname := "<Apellido>"
  var number = 100
  var (
  a = 1
  b = 2
  c = 3
  )
  fmt.Print ("Ingresa tu nombre:")
  fmt.Scanf("%s",&name)
  fmt.Print ("Ingresa tu apellido:")
  fmt.Scanf("%s",&lastname)
  fmt.Println("OK!")
  fmt.Printf(HELLO_USER, name, lastname)
  fmt.Println(number,a,b,c)
}
