package main

import (
  "fmt"
)

type Course struct {
  Name string
  Slug string
  Skills []string
}

type Carrer struct {
  Name string
  Courses []Course
}

func main(){
  course1 := Course {Name: "Go", Slug: "go", Skills: []string {"programar","abstraccion", "lógica"}}
  fmt.Println(course1)
  course2 := new (Course)
  course2.Name = "Java"
  course2.Slug = "java"
  course2.Skills = []string {"c","programar","computacion"}
  fmt.Println(course2)

  courses := []Course {course1, *course2}
  carrer1 := Carrer {Name: "Programacion" ,Courses: courses}
  fmt.Println(carrer1)
}
