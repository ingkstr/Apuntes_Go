package main

import "fmt"

func main() {
  fmt.Println("******FOR CONVENCIONAL*******")
  for i := 0 ; i < 5 ; i ++ {
    fmt.Println("Valor es: ",i)
  }

  fmt.Println("******FOR CON SOLO CONDICIONAL*******")
  c := 10

  for ;c > 5 ; {
    c -= 2
    fmt.Println("c vale ",c)
  }

  fmt.Println("******FOR NO DEFINIDO*******")

  s := 1000

  for {
    s -= 1
    if (s == 0){
      fmt.Println("BREAK!!!")
      break
    }

  }

}
