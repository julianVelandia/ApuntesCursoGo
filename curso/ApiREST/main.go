//https://www.youtube.com/watch?v=pQAV8A9KLwk
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type task struct {
	ID      int    `json:ID`
	Name    string `jason:Name`
	Content string `json:Content`
}

type allTask []task

var tasks = allTask{ //Tareas a interactuar con el Api
	{
		ID:      1,
		Name:    "Tarea 1",
		Content: "Contenido tarea 1",
	},
	{
		ID:      2,
		Name:    "Tarea 2",
		Content: "Contenido tarea 2",
	},
	{
		ID:      3,
		Name:    "Tarea  3",
		Content: "Contenido tarea 3",
	},
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", indexRoute)
	router.HandleFunc("/tareas", getTasks).Methods("GET")
	router.HandleFunc("/tareas", createTask).Methods("POST")
	router.HandleFunc("/tareas/{id}", getTask).Methods("GET")
	router.HandleFunc("/tareas/{id}", deleteTask).Methods("DELETE")
	router.HandleFunc("/tareas/{id}", updateTask).Methods("PUT")

	log.Fatal(http.ListenAndServe(":3000", router))
}

func indexRoute(w http.ResponseWriter, r *http.Request) {
	//Vista Home
	fmt.Fprintf(w, "Bienvenido")
}

func getTasks(w http.ResponseWriter, r *http.Request) {
	//Retorna todas las tareas

	w.Header().Set("Content-Type", "application/json") //Cabecera http
	json.NewEncoder(w).Encode(tasks)
}

func createTask(w http.ResponseWriter, r *http.Request) {
	//Crea una nueva tarea

	var newTask task
	reqBody, err := ioutil.ReadAll(r.Body) //Información que el cliente está enviando al servidor

	if err != nil {
		fmt.Fprintf(w, "Inserta una tarea")
	}

	json.Unmarshal(reqBody, &newTask) //Se asigna el dato a newTask

	newTask.ID = len(tasks) + 1 //Id autogenerado de acuerdo al número de tareas

	tasks = append(tasks, newTask)

	w.Header().Set("Content-Type", "application/json") //Cabecera http
	w.WriteHeader(http.StatusCreated)

	json.NewEncoder(w).Encode(newTask) //Responde al cliente con tarea añadida

}

func getTask(w http.ResponseWriter, r *http.Request) {
	//Retorna una tarea

	vars := mux.Vars(r) //Extrae las variables y las asigna a vars

	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID invalido")
		return
	}

	for _, task := range tasks {
		if task.ID == taskID {
			w.Header().Set("Content-Type", "application/json") //Cabecera http
			json.NewEncoder(w).Encode(task)
		}
	}
}

func deleteTask(w http.ResponseWriter, r *http.Request) {
	//Borra una tarea

	vars := mux.Vars(r) //Extrae las variables y las asigna a vars

	taskID, err := strconv.Atoi(vars["id"])

	if err != nil {
		fmt.Fprintf(w, "ID invalido")
		return
	}

	for i, task := range tasks {
		if task.ID == taskID {
			//Conserva todo lo que está antes y lo
			//concatena con lo que está después del indice
			tasks = append(tasks[:i], tasks[i+1:]...)

		}
	}
}

func updateTask(w http.ResponseWriter, r *http.Request) {
	//Actualiza una tarea

	vars := mux.Vars(r) //Extrae las variables y las asigna a vars

	taskID, err := strconv.Atoi(vars["id"])

	var updatedTask task

	if err != nil {
		fmt.Fprintf(w, "ID invalido")
		return
	}

	reqBody, err := ioutil.ReadAll((r.Body))

	if err != nil {
		fmt.Fprintf(w, "Datos invalidos")
		return
	}

	json.Unmarshal(reqBody, &updatedTask)

	for i, task := range tasks {
		if task.ID == taskID {
			//Elimina la tarea y la vuelve a escribir

			//Conserva todo lo que está antes y lo
			//concatena con lo que está después del indice
			tasks = append(tasks[:i], tasks[i+1:]...)

			updatedTask.ID = taskID
			tasks = append(tasks, updatedTask)
		}
	}
}
