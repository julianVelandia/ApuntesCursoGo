//https://www.youtube.com/watch?v=G58gN0lIbyI
package main

import (
	"net/http"
	"text/template"
)

var plantillas = template.Must(template.ParseGlob("plantillas/*"))

func main() {
	http.HandleFunc("/", Home)

	http.ListenAndServe(":8080", nil)
}

func Home(w http.ResponseWriter, r *http.Request) {
	plantillas.ExecuteTemplate(w, "home", nil)
}
