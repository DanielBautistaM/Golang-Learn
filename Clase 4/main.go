package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"main.go/rutas"
)

// main es el punto de entrada de nuestra aplicación web
func main() {
	// Creamos un nuevo router usando Gorilla Mux
	mux := mux.NewRouter()

	// Definimos las rutas de nuestra aplicación
	// Cada ruta se asocia con su respectiva función manejadora
	mux.HandleFunc("/home", rutas.Home)                           // Ruta básica
	mux.HandleFunc("/notosotros", rutas.Nosotros)                // Otra ruta básica
	mux.HandleFunc("/parametros/{id}/{slug}", rutas.Parametros)  // Ruta con parámetros en URL
	mux.HandleFunc("/query/{id}/{slug}/{texto}", rutas.Query)    // Ruta con parámetros y query strings
	mux.HandleFunc("/estructura", rutas.Estructuras)             // Ruta para demostrar structs en templates

	// Configuración para servir archivos estáticos (CSS, JS, imágenes)
	// StripPrefix quita "/public" de la URL antes de buscar el archivo
	s := http.StripPrefix("/public", http.FileServer(http.Dir("./public/")))
	mux.PathPrefix("/public").Handler(s)

	// Cargamos variables de entorno desde el archivo .env
	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}

	// Obtenemos el puerto desde las variables de entorno
	puerto := os.Getenv("PORT")
	if puerto == "" {
		puerto = "8080" // Puerto por defecto si no está definido
	}

	// Iniciamos el servidor
	server := &http.Server{
		Addr:         "localhost:" + puerto,
		Handler:      mux,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	fmt.Println("Corriendo")
	log.Fatal(server.ListenAndServe())
}

/*
func main() {
	http.HandleFunc("/", func(response http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(response, "hola dasdasdsdsdwe")
	})
	fmt.Println("Corriendo")
	log.Fatal(http.ListenAndServe("Localhost:8081", nil))
}
*/
