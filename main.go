package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	// Establecer los encabezados CORS
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")   // Permite solicitudes desde el puerto 3000
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE") // Métodos permitidos
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")           // Encabezados permitidos

	// Si es una solicitud OPTIONS, solo devolvemos una respuesta vacía
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "Backend is running!")
}

func main() {
	http.HandleFunc("/healthcheck", healthCheckHandler)
	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}
