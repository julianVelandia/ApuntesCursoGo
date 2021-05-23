package Paquete1 //Esta línea hace que la carpeta se vuelva un paquete

import "fmt"

//FuncionPaquete1 Las funciones que inician con mayúscula son funciones públicas
func FuncionPaquete1(){ 
	fmt.Println("función paquete 1")
}


//funcionPrivada función privada porque inicia con minúscula
func funcionPrivada(){
	fmt.Println("función paquete 1 privada")
}