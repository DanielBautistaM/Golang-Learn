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

func main() {
	mux := mux.NewRouter()
	mux.HandleFunc("/", rutas.Home)
	mux.HandleFunc("/home", rutas.Nosotros)
	mux.HandleFunc("/parametros", rutas.Parametros)
	mux.HandleFunc("/query", rutas.Query)

	errorVariables := godotenv.Load()
	if errorVariables != nil {
		panic(errorVariables)
	}
	server := &http.Server{
		Addr:         "localhost:" + os.Getenv("PORT"),
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
