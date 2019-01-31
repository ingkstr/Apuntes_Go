package main

import (
  "fmt"
  "strings"
  )


func main() {
  var text = "Hola mundo. Hola amigos. Hola a todos sus habitantes!"
  fmt.Println("Texto original: ",text)  //Original
  fmt.Println("Texto en mayúsculas: ",strings.ToUpper(text))  //Convierte a mayusculas
  fmt.Println("Texto en minúsculas: ",strings.ToLower(text))  //Convierte a minúsculas
  fmt.Println("Reemplazo total: ",strings.Replace(text,"Hola","HELLO", -1))  // Remplaza todos los "Hola" por ""HELLO
  fmt.Println("Reemplazo de la primera coincidencia: ",strings.Replace(text,"Hola","HELLO", 1))  // Remplaza el primer "Hola" por ""HELLO
  fmt.Println("Reemplazo de las primera y segunda coincidencia: ",strings.Replace(text,"Hola","HELLO", 2))  // Remplaza el segundo "Hola" por ""HELLO
  separa_espacio := strings.Split(text," ") //Separación por espacios
  fmt.Println("**********División por espacios************")
  for i := range separa_espacio {
    fmt.Println(separa_espacio[i])
  }
  separa_punto := strings.Split(text,".") //Separación por puntos
  fmt.Println("**********División por puntos************")
  for i := range separa_punto {
    fmt.Println(separa_punto[i])
  }
}
