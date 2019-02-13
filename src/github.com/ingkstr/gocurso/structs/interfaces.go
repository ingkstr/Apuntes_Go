package structs

//Platzi interface de Course y Carrer
type Platzi interface {
  Subscribe(name string)
}

//InterfaceTest Función de pruebe
func InterfaceTest(){
  course := Course {Name: "Go", Slug: "go", Skills: []string {"programar","abstraccion", "lógica"}}
  courseint := new(InternalCourse)
  courseint.Name = "GoAdv"
  courseint.Slug = "goadv"
  courseint.Skills = []string {"programar","abstraccion", "lógica"}
  callSubscribe(course)
  callSubscribe(courseint)
}

//callSubscribe ejecuta una impresion de un usuario
func callSubscribe(p Platzi){
  p.Subscribe("Jose")
}
