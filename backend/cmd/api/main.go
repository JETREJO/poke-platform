package main

import (
	"context"
	"log"
	"net/http"

	"pokemon-platform/backend/config"
	internalHttp "pokemon-platform/backend/internal/http"

	"pokemon-platform/backend/internal/database"
	"pokemon-platform/backend/internal/pokemon"
)

func main() {

	config.Load()

	conn, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// Cierra la conexión a la base de datos cuando termine la función
	defer conn.Close(context.Background())

	log.Println("[MAIN] Database ready")

	// Creamos una instancia de nuestro "repository"
	repository := pokemon.NewRepository(conn)

	// Creamos una instancia para el service
	service := pokemon.NewService(repository)

	// Creamos una instancia de nuestro Handler
	// Ahora el Handler ya no se conecta directo al Repository, sino al Service
	handler := pokemon.NewHandler(service)

	router := internalHttp.NewRouter(handler)

	log.Printf("[MAIN] Server running on :%s", config.App.ServerPort)

	// 1. La función "ListenAndService" hace dos cosas en esta implementación:
	// - Abre el puerto 8080
	// - Se queda esperando peticiones
	//
	// 2. La función "ListenAndService" espera dos parámetros:
	// - El puerto a abrir
	// - El 'router' a utilizar. En este caso  es el que creamos en la carpeta 'internal/http'
	err = http.ListenAndServe(":"+config.App.ServerPort, router)
	if err != nil {
		log.Fatal(err)
	}
}
