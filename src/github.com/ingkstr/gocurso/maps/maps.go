package maps

//GetMap crea y devuelve un mapa de ejemplo
func GetMap() map[string]int {
  mapTest := make(map[string]int)
  mapTest["llave1"] = 2
  mapTest["llave2"] = 100
  return mapTest
}

func GetMap2() map[string]int {
  mapTest2  := map[string]int{
    "Juan" : 19 ,
    "Mario" : 20 ,
    "Indeseado" : 99 , 
  }
  delete(mapTest2,"Indeseado")
  return mapTest2
}
