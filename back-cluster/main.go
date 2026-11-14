package main

import (
	"fmt"
	"net/http"
)

func main() {
	// 1. Registrar la ruta raíz "/" y su función manejadora
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Escribir la respuesta "Hola Mundo" en el cuerpo de la respuesta HTTP
		fmt.Fprintln(w, "Hola Mundo")
	})

	// 2. Iniciar el servidor en el puerto 8080
	// Se imprime un mensaje en consola para saber que está corriendo
	fmt.Println("Servidor escuchando en http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		// Si hay un error al iniciar el servidor (ej. puerto ocupado), se muestra
		fmt.Printf("Error al iniciar el servidor: %v\n", err)
	}
}
