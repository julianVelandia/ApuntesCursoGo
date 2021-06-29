package main

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//w.Header().Add("nombre", "valor")
		//fmt.Println("El método es: " + r.Method)

		//fmt.Fprintf(w, "Hola mundo")
		http.Redirect(w, r, "/redireccion", http.StatusMovedPermanently)

	})

	http.HandleFunc("/redireccion", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Redireccionamiento\n \n")

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

		//Para trabajar con la url se usa r

		fmt.Println(r.URL.Query())
		parametro := r.URL.Query().Get("parametro")
		if len(parametro) != 0 {
			fmt.Println(parametro)
		}

		//Generar query
		fmt.Println(r.URL)
		values := r.URL.Query()
		values.Add("nombre", "Julián")
		values.Add("edad", "21")
		values.Add("profesión", "ingeniero")

		r.URL.RawQuery = values.Encode()
		fmt.Println(r.URL)
		//http.NotFound(w, r)
		//http.Error(w, "primer error", http.StatusConflict)

		//Headers
		fmt.Println(r.Header)

		//crear urls

		url := createURL()
		fmt.Println("la url final es:", url)

		//Generar peticiones
		request, err := http.NewRequest("GET", url, nil)

		if err != nil {
			panic(err)
		}
		cliente := &http.Client{}
		response, err := cliente.Do(request)

		if err != nil {
			panic(err)
		}
		fmt.Println("header: ", response.Header)
		fmt.Println("body", response.Body)
		fmt.Println("status", response.Status)
	})

	log.Fatal(http.ListenAndServe("localhost:3000", nil))

}

func createURL() string {
	u, err := url.Parse("/params")
	if err != nil {
		panic(err)
	}

	u.Host = "localhost:300"
	u.Scheme = "http"

	query := u.Query() //map
	query.Add("nombre", "valor")
	u.RawQuery = query.Encode()
	return u.String()
}
