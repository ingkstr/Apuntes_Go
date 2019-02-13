package main

import (
  "fmt"
  "./name"
  "./numbers"
  )


func main(){
  nombre := name.GetName()
  fmt.Println("El nombre es ", nombre)
  fmt.Println("La suma de 30 + 50 es", numbers.GetSuma(30,50))
  fmt.Println("La resta de 120 - 40 es", numbers.GetResta(120,40))
}
