package main

import "fmt"

func main() {

  var arreglo_texto [2] string
  arreglo_ints := [3] int {3,6,9}
  arreglo_texto[0] = "texto1"
  arreglo_texto[1] = "texto2"

  var slice_texto [] string
  slice_texto = append(slice_texto, "esto","es","1","slice")

  fmt.Println("Contenido de arreglo de texto: ",arreglo_texto)
  fmt.Println("Contenido de arreglo de enteros: ",arreglo_ints)
  fmt.Println("Contenido de slice de texto: ",slice_texto)

}
