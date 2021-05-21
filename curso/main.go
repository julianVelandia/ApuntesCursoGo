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

	fmt.Println("Tu nombre es:", nombre, "Tu edad es:", edad)

}