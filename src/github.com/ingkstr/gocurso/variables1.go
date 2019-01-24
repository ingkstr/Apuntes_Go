package main 

import "fmt"

func main() { 
  var name string 
  fmt.Print ("Ingresa tu nombre:") 
  fmt.Scanf("%s",&name)
  fmt.Println("OK!") 
  fmt.Printf("Welcome %s to Go!!!\n",name) 
}
