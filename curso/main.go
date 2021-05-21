package main //Define el paquete main

import "fmt" // Importa fmt

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

}

//Funciones

func funcion1(){
	fmt.Println("Función 1")
}