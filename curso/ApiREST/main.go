//https://www.youtube.com/watch?v=pQAV8A9KLwk
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `jason:Name`
	Content string `json:Content`
}

type allTask []task

var tasks = allTask{
	{
		ID:      1,
		Name:    "Tarea 1",
		Content: "Contenido tarea 1",
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)

	log.Fatal(http.ListenAndServe(":3000", router))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bienvenido")
}
