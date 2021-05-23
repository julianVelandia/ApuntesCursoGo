package main //Define el paquete main


// Importa librerías
import (
	"fmt"
	"strings"
	//"github.com/julianVelandia/ApuntesCursoGo/curso/Paquete1"
)


//Estructuras
//Análogo de clases en go
type Estructura struct{
	atributo1 string
	atributo2 [] int
} 

//Herencia
type Estructura2 struct{
	atributo3 string
	atributo4 [] int
	Estructura // Herencia
} 



//Métodos
//metodo1 
func (e Estructura) metodo1(parametro1 string){
	fmt.Println("Método 1: ", parametro1," Clase: ",e.atributo1)
}



func main()  {
	/*
	Comentarios largos
	*/
	fmt.Println("Hola") //Imprimir

	/*
	Tipos de datos
	-Enteros
	-Flotantes
	-Cadenas
	-Booleanos
	-Derivados
	*/

	//Definir variables
	var nombre string
	nombre = "Julián"
	fmt.Println(nombre)

	//Definir muchas variables
	var num1,num2,num3,num4 int = 1,2,3,4

	//Definir muchs variables de diferente tipo de dato
	var(
		num5 float64
		str1 string 
		bool1 bool
		str2 = "Cadena1"
	)
	num5=5
	str1 = "Cadena2"
	bool1 = true


	//Concatenar
	fmt.Println("Números:",num1,num2,num3,num4,num5,"Cadenas:",str1,str2,"Bool:",bool1)

	//Para no usar var (tipado dinámico)
	num6 := 10
	str3 := "String"

	fmt.Println(num6,str3)

	//Constantes
	const con1 = 10 //Se pueden defir fuera del main

	//Recibir variables por teclado
	var(
		nom string
		edad int
	)

	fmt.Println("Ingrese su nombre:")
	fmt.Scanln(&nom)

	fmt.Println("Ingrese su edad:")
	fmt.Scanln(&edad)

	fmt.Println("Tu nombre es:", nom, "Tu edad es:", edad)

	/*
	Operadores 
	+ - * /
	== != < > <= >=
	not: !
	And: &&
	Or: ||
	*/

	//Condicionales
	if true{
		fmt.Println("verdadero")
	}else{
		fmt.Println("No se ejecuta")
	}

	if num1 > 3{
		fmt.Println("Numero mayor que 3")
	} 


	//Incrementos
	//a +=1 Es igual a a++
	a := 0
	a++
	fmt.Println(a)

	//Bucles
	//No existe while ni do-while, solo for
	i := 0
	for i<5{
		fmt.Println(i)
		i++	
	}

	for c := 0; c<10;c++{
		
		if c == 3{
			fmt.Println("No se ejecuta el 3")
			continue
		}

		if c == 7{
			fmt.Println("Se corta el bucle en 7")
			break
		}
		fmt.Println(c)
	}


	//Estructuras de datos
	//Array de tamaño estático
	var arr1 [3] string
	arr1[0] = "Dato0"
	arr1[1] = "Dato1"
	arr1[2] = "Dato2"

	arr2 := [2] int {1,2}

	fmt.Println(arr1)
	fmt.Println(arr2)

	//Slicen
	var sli1 [] string

	sli1 = append(sli1,"dato")

	fmt.Println(sli1,"dato")

	//Funciones
	funcion1()
	nombreFunc2,edadFunc2 := funcion2("Julián")
	fmt.Println(nombreFunc2,edadFunc2)

	//Funciones con strings
	//documentación https://golang.org/pkg/strings/

	texto := "Hola mundo"
	fmt.Println(strings.ToUpper(texto))
	fmt.Println(strings.Count(texto,"a"))


	//Paquetes
	//Paquete1.FuncionPaquete1()

	//Maps
	//Análogos a los diccionarios en python

	mapa1 := map[string] int{
		"Llave1":4,
		"Llave2":3,
		"Llave3":2,
		"Llave4":1,
	}

	fmt.Print("Mapa 1: ",mapa1)
	fmt.Println()

	//Agregar
	mapa1["NuevaClave"] = 5
	fmt.Print("Mapa 1 con nuevo valor: ",mapa1)

	//Imprimir
	fmt.Print("Mapa 1 por Clave: ",mapa1["Llave2"])

	//Eliminar
	delete(mapa1, "Llave4")
	fmt.Print("Mapa 1 eliminando un valor: ",mapa1)
	

	//Declararlo vacio
	mapa2 := make(map[string] string)
	mapa2["uno"] = "UNO"

	fmt.Println(mapa2)

	
	//Métodp Range
	arr3 := []string{
		"uno","dos","tres",
	}

	for i, num := range arr3{
		fmt.Println(i," ==> ",num)
	}

	for k,v := range mapa1{
		fmt.Println("Key: ",k,"Value: ",v)
	}


	//Estructuras
	est1 := Estructura{
		atributo1 : "Atributo 1",
		atributo2 : []int {
			1,2,3,4,
		},
	}

	fmt.Println("Estructuras: ",est1)
	fmt.Println("Atributo 1: ",est1.atributo1)
	fmt.Println("Metodos")
	est1.metodo1("Parametro")

	//Herencia
	est2 := new(Estructura2) //Estructura vacía
	est2.atributo1 = "Atributo1"
	est2.atributo2 = []int {1,2,3,}
	est2.atributo3 = "Atributo3"
	est2.atributo4 = []int {4,5,6,}
	fmt.Println(est2)
}

//Funciones

func funcion1(){
	fmt.Println("Función 1")
}

func funcion2(varNombre string) (string,int){ //Indica lo que devuelve la función
	
	var varEdad int
	varEdad = 21

	return varNombre,varEdad
}

