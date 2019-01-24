CURSO DE GO
===========

## Instalación en Ubuntu ##

´´´
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go
´´´

## Creación de entorno de trabajo ##
Suponemos que tendremos nuestro entorno en /home/USUARIO/go.

Para más información, pueden ver en https://github.com/golang/go/wiki/SettingGOPATH

´´´
cd
mkdir go
export GOPATH=$HOME/go
source ~/.bash_profile
´´´

Hay que armar un esquema de carpetas similar a este.

go 

└── src 

    └── github.com 

        └── USUARIO

            └── PROYECTO

## Hola mundo ##

Debemos ir a la localidad de nuestro proyecto

> cd go/src/github.com/USUARIO/PROYECTO

Un "homa mundo sería de la siguiente manera

> touch main.go

> vi main.go

´´´
package main 

import "fmt" 

func main() { 
  fmt.Print("Hello Go!!!") 
}
´´´

Un "hola mundo" con entrada de teclado sería así.

> touch main2.go

> touch main2.go

´´´
package main 

import "fmt"

func main() { 
  var name string 
  fmt.Println ("Ingresa tu nombre:") 
  fmt.Scanf("%s",&name) 
  fmt.Printf("Hello bitch %s!!!",name) 
}
´´´

Se pueden ejecutar de dos maneras.

 - ´Interpretado´ - Lo ejecuta en el momento

> go run main.go 

 - ´Compilado´ - Genera un archivo sin extensión con el mismo nombre

> go build main.go

> ./main
