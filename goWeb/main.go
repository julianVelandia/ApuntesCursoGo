package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Add("nombre", "valor")
		//fmt.Println("El método es: " + r.Method)

		fmt.Fprintf(w, "Hola mundo")
		//http.Redirect(w, r, "/adios", http.StatusMovedPermanently)

	})

	http.HandleFunc("/adios", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "El método es GET")
		case "POST":
			fmt.Fprintf(w, "El método es POST")
		case "PUT":
			fmt.Fprintf(w, "El método es PUT")
		case "DELETE":
			fmt.Fprintf(w, "El método es DELETE")
		}

		fmt.Fprintf(w, "\nHola mundo")

		//Para trabajar con la url se usa r

		fmt.Println(r.URL.Query())
		parametro := r.URL.Query().Get("parametro")
		if len(parametro) != 0 {
			fmt.Println(parametro)
		}

		//http.Redirect(w, r, "/adios", http.StatusMovedPermanently)

		//http.NotFound(w, r)
		//http.Error(w, "primer error", http.StatusConflict)
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}
