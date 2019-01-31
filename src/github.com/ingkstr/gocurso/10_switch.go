package main

import (
  "fmt"
  )


func main() {
    var numero int
    fmt.Println("Dame un valor del 1 al 10")
    fmt.Scanf("%d",&numero)
    revision(numero)
}

func revision(numero int) {
  switch numero {
    case 1:
      fmt.Println("LE DISTE AL NUMERO UNO")
    default:
      fmt.Println("ES ENTRE DOS Y DIEZ")
  }

  switch {
  case numero % 2 == 0:
      fmt.Println("El número es par")
    default:
      fmt.Println("El número es non")
  }

}
