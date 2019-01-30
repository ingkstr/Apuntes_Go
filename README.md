CURSO DE GO
===========

### Índice ###

 * [Instalación en Ubuntu](#instalación-en-ubuntu)
 * [Creación de entorno de trabajo](#creación-de-entorno-de-trabajo)
 * [Hola mundo](#hola-mundo)
 * [Recibir valores de teclado](#Recibir-valores-de-teclado)
 * [Variables y constantes](#variables-y-constantes)
 * [Funciones](#funciones)
 * [Datos numéricos](#Datos-numéricos)
 * [Cadenas de texto](#Cadenas-de-texto)
 * [Array y Slice](#Array-y-Slice)
 * [Condicionales](#Condicionales)
 * [For break](#For-break)

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

> touch 01_hello.go

> vi 02_hello.go

```
package main

import "fmt"

func main() {
  fmt.Print("Hello Go!!!")
}
```

Se pueden ejecutar de dos maneras.

 - `Interpretado` - Lo ejecuta en el momento

> go run 01_hello.go

 - `Compilado` - Genera un archivo sin extensión con el mismo nombre

> go build 01_hello.go

> ./01_hello

## Recibir valores de teclado ##

Se usa la palabra reservada "var", usando la sintaxis siguiente

```
var VARIABLE *TIPO_VARIABLE*
```

Se usa `scanf` para permitir entrada en teclado

> touch 02_entrada_datos.go

> 02_entrada_datos.go

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

> vi 03_variables-constantes.go

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

## Funciones ##

Se usa la palabra reservada `func`, seguido del nomnbre de la función y finaliza con el tipo de variable.

```
func getName() string {
  var name string
  name = "Sin nombre"
  fmt.Print ("Ingresa tu nombre:")
  fmt.Scanf("%s",&name)
  return name
}
```

Esta se puede mandar llamar de la siguiente manera

```
name := getName()
```

Se pueden hacer una función que devuelva más de un valor

```
func getVariables() (int, int, int){
  return 1,2,3
}
```

Se mandaría llamar de la siguiente manera

```
a, b, c := getVariables()
```

Para tener funciones con argumentos, se hace algo como esto

```
func sum(a int, b int) int {
  return a + b
}
```

Y se invoca de la siguiente manera

```
number := sum(50, 50)
```
## Datos numéricos ##

 * `int` : 32 bits comprendiendo la mitad para positivos como negativos
 * `unint` : 32 bits partiendo de cero
 * `int32` : Enteros de 32 bits
 * `int64` : Enteros de 64 bits
 * `float32` : Valor flotante cerrado a 32 bits  
 * `float64` : Valor más preciso usando 64 bits

## Cadenas de texto ##

 Unicode es el tipo de codificación de texto default en Go. Por lo que puede usar textos con acentos, caracteres especiales y japones.

 ```
 func main() {
   fmt.Println("もしもし!")
 }
```

Obtenemos lo siguiente:

```
もしもし!
```

Las cadenas son conformadas por bytes uno a uno. Para llamar a cada uno, se toma en cuenta como un arreglo.

> fmt.Println("hola"[0]) // Devuelve "104" ya que maneja ascii

> fmt.Println(string("hola"[0])) // Devuelve "h"

Para obtener la longitud

> fmt.Println(len("hola")) // Devuelve "4"


## Array y Slice ##

Ambas son estructuras de datos con valores del mismo tipo.

 * `Array` : Siempre tiene un tamaño fijo. No se puede expandir.

Se declara de las siguientes dos maneras

```
var arreglo_texto [2] string
arreglo_ints := [3] int {3,6,9}
```

Para asignar valores, se hace lo siguiente

```
  arreglo_texto[0] = "texto1"
  arreglo_texto[1] = "texto2"
```

 * `Slice` : Su tamaño es dinámico

Se declara de la siguiente manera solamente

```
var slice_texto [] string
```

La forma de cargarle información es la siguiente

```
slice_texto = append(slice_texto, "esto","es","1","slice")
```

## Condicionales ##

La sintaxis usada es

```
if [condicion] {
  //instrucciones
} else {
  //instrucciones
}
```

También se vale anidad Condicionales

```
if [condicion] {
  //instrucciones
} else if [condicion2] {
  //instrucciones
}
```

Igual se puede asignar valor en el mismo if

```
if variable := 2 ; variable == 4 {
  fmt.Println("Se cumple")
}
```

## For break ##

En Go no existe los bucles while ni do-while. Solo existe el `for`

La sintaxis es la siguientes

```
for VARIABLE := VALOR; CONDICION_FIN ; INCREMENTO {
  //Inatrucciones
}
```

Ejemplo

```
for i := 0 ; i < 5 ; i ++ {
  fmt.Println("Valor es: ",i)
}
```

Es posible usar un for sin declaración inicial ni incremento.

Ejemplo

```
c := 10

for ;c > 5 ; {
  c -= 2
  fmt.Println("c vale ",c)
}
```

Se puede manejar un bucle infinito, siempre que haya una condición que la acabe con un ` break`

 ```
 s := 1000

 for {
   s -= 1
   if (s == 0){
     fmt.Println("BREAK!!!")
     break
   }

 } 
 ```
