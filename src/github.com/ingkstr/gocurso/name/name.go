package name

import "fmt"

//Getname obtiene y retorna el nombre de una persona
func GetName() string {
    var name string
    fmt.Print("Ingresa tu nombre: ")
    fmt.Scanf("%s", &name)
    return name
}
