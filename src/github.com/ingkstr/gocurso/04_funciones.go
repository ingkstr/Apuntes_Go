package main

import "fmt"

const HELLO_USER string = "Hola %s, Bienvenido al fascinante mundo de Go\n"

func main() {
  name := getName()
  number := sum(50, 50)
  a, b, c := getVariables()
  fmt.Printf(HELLO_USER, name)
  fmt.Println(number,a,b,c)
}

func getName() string {
  var name string
  name = "Sin nombre"
  fmt.Print ("Ingresa tu nombre:")
  fmt.Scanf("%s",&name)
  return name
}

func getVariables() (int, int, int){
  return 1,2,3
}

func sum(a int, b int) int {
  return a + b
}
