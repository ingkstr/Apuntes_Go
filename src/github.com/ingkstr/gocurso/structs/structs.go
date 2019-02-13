package structs

import "fmt"

//Course estructura de curso
type Course struct {
  Name string
  Slug string
  Skills []string
}

//Subscribe Función de curso para mostrar quien se ha suscrito
func (c Course) Subscribe (name string){
  fmt.Printf("El alumno %s se ha inscrito al curso %s. \n",name,c.Name)
}

//Carrer estructura de carrera
type InternalCourse struct {
  Course
}

//Subscribe Función de carrera para mostrar quien se ha suscrito
func (c InternalCourse) Subscribe (name string){
  fmt.Printf("El alumno %s se ha inscrito al curso interno %s. \n",name,c.Name)
}
