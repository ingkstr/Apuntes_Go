package main

import (
  "fmt"
  "./maps"
)

func main(){
  fmt.Println("El valor del mapa completo es",maps.GetMap())
  fmt.Println("El valor del mapa en la llave1 es",maps.GetMap()["llave1"])
  fmt.Println("El valor del mapa en la llave2 es",maps.GetMap()["llave2"])
  fmt.Println("El valor del mapa 2 completo es",maps.GetMap2())
}
