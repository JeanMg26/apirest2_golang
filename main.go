package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Empleado struct {
	Id     int    `json:id,omitempty`
	Nombre string `json:nombre,omitempty`
	Correo string `json:correo,omitempty`
}

var empleados []Empleado

func main() {

	empleados = []Empleado{
		{
			Id:     1,
			Nombre: "Giancarlo",
			Correo: "jeanmg25@gmail.com",
		},
		{
			Id:     2,
			Nombre: "Juancarlo",
			Correo: "juanca@gmail.com",
		},
	}

	// Rutas mux
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/empleados", GetEmpleados).Methods("GET")
	router.HandleFunc("/empleados/{id}", GetEmpleado).Methods("GET")
	router.HandleFunc("/empleados", StoreEmpleado).Methods("POST")
	router.HandleFunc("/empleados/{id}", UpdateEmpleado).Methods("PUT")
	router.HandleFunc("/empleados/{id}", DeleteEmpleado).Methods("DELETE")
	// Servidor}
	fmt.Println("Iniciando servidor...")
	http.ListenAndServe(":8000", router)
}

func GetEmpleados(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empleados)

}

func GetEmpleado(w http.ResponseWriter, r *http.Request) {
	empleado := mux.Vars(r)
	idEmpleado, error := strconv.Atoi(empleado["id"])
	if error != nil {
		fmt.Fprintf(w, "Ingrese un id correcto")
	}

	for _, item := range empleados {
		if item.Id == idEmpleado {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Empleado{})
}

func StoreEmpleado(w http.ResponseWriter, r *http.Request) {
	var empleado Empleado
	_ = json.NewDecoder(r.Body).Decode(&empleado)
	empleado.Id = len(empleados) + 1
	empleados = append(empleados, empleado)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(empleado)
}

func DeleteEmpleado(w http.ResponseWriter, r *http.Request) {
	empleado := mux.Vars(r)
	idEmpleado, error := strconv.Atoi(empleado["id"])
	if error != nil {
		fmt.Fprintf(w, "ingresar un id valido")
	}
	for index, item := range empleados {
		if item.Id == idEmpleado {
			empleados = append(empleados[:index], empleados[index+1:]...)
			fmt.Fprintf(w, "El elemento con id %v fue eliminado", idEmpleado)
		}
	}
}

func UpdateEmpleado(w http.ResponseWriter, r *http.Request) {
	empleado := mux.Vars(r)
	var updateEmpleado Empleado
	idEmpleado, error := strconv.Atoi(empleado["id"])
	if error != nil {
		fmt.Fprintf(w, "ingresar un id valido")
	}

	_ = json.NewDecoder(r.Body).Decode(&updateEmpleado)

	for index, item := range empleados {
		if item.Id == idEmpleado {
			// Eliminamos el empleado
			empleados = append(empleados[:index], empleados[index+1:]...)
			// Agregamos el empleado a editar

			updateEmpleado.Id = idEmpleado
			empleados = append(empleados, updateEmpleado)
			// w.Header().Set("Content-Type", "application/json")
			// json.NewEncoder(w).Encode(updateEmpleado)
			fmt.Fprintf(w, "El empleado con el id %v fue actualiazado", idEmpleado)
		}
	}
}
