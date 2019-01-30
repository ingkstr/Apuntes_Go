//Programa para sumar y restar dos valores

package main

import "fmt"

func main() {
  valor1 := getValor("1")
  valor2 := getValor("2")
  fmt.Printf("La suma es %f\n", suma(valor1, valor2))
  fmt.Printf("La resta es %f\n", resta(valor1, valor2))

}

func getValor(label string) float64 {
  var valor float64
  fmt.Printf("Ingresa valor %s:\n",label)
  fmt.Scanf("%f",&valor)
  return valor
}


func suma(a float64, b float64) float64 {
  return a + b
}

func resta(a float64, b float64) float64 {
  return a - b
}
