//https://www.youtube.com/watch?v=G58gN0lIbyI
package main

import (
	"fmt"
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {

	port := ":8000" //Ponerlo en un archivo de variables o algo así
	http.HandleFunc("/", Home)
	http.HandleFunc("/crear", Crear)
	http.HandleFunc("/actualizar", Actualizar)
	http.HandleFunc("/leer", Leer)
	http.HandleFunc("/borrar", Borrar)
	fmt.Println("Run server in the port", port)

	http.ListenAndServe(port, nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "home", nil)
}

func Crear(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "crear", nil)
}

func Actualizar(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "actualizar", nil)
}

func Leer(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "leer", nil)
}

func Borrar(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "borrar", nil)
}
