//Programa para sumar y restar dos valores

package main

import "fmt"

func main() {
  valor1 := getValor("1")
  valor2 := getValor("2") 
  fmt.Printf("La suma es %d\n", suma(valor1, valor2))
  fmt.Printf("La resta es %d\n", resta(valor1, valor2))

}

func getValor(label string) int {
  var valor int
  fmt.Printf("Ingresa valor %s:\n",label)
  fmt.Scanf("%d",&valor)
  return valor
}

func getVariables() (int, int, int){
  return 1,2,3
}

func suma(a int, b int) int {
  return a + b
}

func resta(a int, b int) int {
  return a - b
}
