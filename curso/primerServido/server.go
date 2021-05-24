//https://www.youtube.com/watch?v=juVYPx0UG80
package main

import (
	"net/http" // librería para levantar el servidor
)

func main(){

	//Rutas
	http.HandleFunc("/", home)
	http.HandleFunc("/ruta1", ruta1)


	//Levanta un servidor en el puerto 3000
	http.ListenAndServe(":3000",nil)
}

func home(w http.ResponseWriter, r *http.Request){//Request que recibe y datos de envio
	w.Write([]byte("<h1>Hola mundo</h1>"))
}

func ruta1(w http.ResponseWriter, r *http.Request){
	w.Write([]byte("<h1>Ruta 1</h1>"))

}