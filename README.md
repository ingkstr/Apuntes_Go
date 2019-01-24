CURSO DE GO
===========

### Índice ###

 * [Instalación en Ubuntu](#instalación-en-ubuntu)
 * [Creación de entorno de trabajo](#creación-de-entorno-de-trabajo)
 * [Hola mundo](#hola-mundo)
 * [Recibir valores de teclado](#Recibir-valores-de-teclado)
 * [Variables y constantes](#variables-y-constantes)

## Instalación en Ubuntu ##

```
sudo add-apt-repository ppa:longsleep/golang-backports
sudo apt-get update
sudo apt-get install golang-go
```

## Creación de entorno de trabajo ##
Suponemos que tendremos nuestro entorno en /home/USUARIO/go.

Para más información, pueden ver en https://github.com/golang/go/wiki/SettingGOPATH

```
cd
mkdir go
export GOPATH=$HOME/go
source ~/.bash_profile
```

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

> touch hello.go

> vi hello.go

```
package main 

import "fmt" 

func main() { 
  fmt.Print("Hello Go!!!") 
}
```

Se pueden ejecutar de dos maneras.

 - `Interpretado` - Lo ejecuta en el momento

> go run main.go 

 - `Compilado` - Genera un archivo sin extensión con el mismo nombre

> go build main.go

> ./main

## Recibir valores de teclado ##

Se usa la palabra reservada "var", usando la sintaxis siguiente

```
var VARIABLE *TIPO_VARIABLE*
```

Se usa `scanf` para permitir entrada en teclado

> touch variables1.go ; vi variables1.go

```
package main 

import "fmt"

func main() { 
  var name string 
  fmt.Print ("Ingresa tu nombre:") 
  fmt.Scanf("%s",&name) 
  fmt.Println("OK!")
  fmt.Printf("Welcome %s to Go!!!\n",name) 
}
```

El comando print se diferencia en 3 modalidades

 - `fmt.Print` imprime un texto sin salto de línea
 - `fmt.Println` imprime un texto con salto de línea
 - `fmt.Print` imprime un texto sin salto de línea, pero permitiendo edición ingresando variables

## Variables y constantes ##

También se pueden crear variables de la siguiente manera, con la que intuye Go el tipo de variable

```
variable2 := "Esta variable es texto"
```

Go también infiere el tipo de dato usando `var` e igualando con una variable inicial.

```
var number = 100"
```

Para múltiples variables, se hace lo siguiente

```
var (
  a = 1
  b = 2
  c = 3
  )
```

Para declarar constantes se usa la palabra `const`. Igualmente, declaramos el tipo o dejamos que Go lo intuya.

La declaración de constantes se hace fuera de la función main.

```
const helloworld string = "Hola %s %s, bienvenido a Go!!!"
const testCont = "Text"

```

Siendo así, tendríamos un programa como este

```
package main

import "fmt"

const HELLO_USER string = "Hola %s %s, Bienvenido al fascinante mundo de Go\n"

func main() {
  var name string
  lastname := "<Apellido>"
  var number = 100
  var (
  a = 1
  b = 2
  c = 3
  )
  fmt.Print ("Ingresa tu nombre:")
  fmt.Scanf("%s",&name)
  fmt.Print ("Ingresa tu apellido:")
  fmt.Scanf("%s",&lastname)
  fmt.Println("OK!")
  fmt.Printf(HELLO_USER, name, lastname)
  fmt.Println(number,a,b,c)
}
```

Y su salida sería algo como esto:

```
Ingresa tu nombre:Emanuel
Ingresa tu apellido:Castelan
OK!
Hola Emanuel Castelan, Bienvenido al fascinante mundo de Go
100 1 2 3
```
