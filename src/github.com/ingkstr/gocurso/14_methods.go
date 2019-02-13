package main

import (
  "fmt"
)

type Course struct {
  Name string
  Slug string
  Skills []string
}

func (c Course) Subscribe (name string){
  fmt.Printf("El alumno %s se ha inscrito al curso %s. \n",name,c.Name)
}


func main(){
  course1 := Course {Name: "Go", Slug: "go", Skills: []string {"programar","abstraccion", "lógica"}}
  course1.Subscribe("Jose")
  course1.Subscribe("Emanuel")
}
