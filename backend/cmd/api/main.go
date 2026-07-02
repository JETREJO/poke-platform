package main

import (
	"context"
	"log"
	"net/http"

	"pokemon-platform/backend/internal/database"
	"pokemon-platform/backend/internal/pokemon"
)

func main() {

	conn, err := database.Connect()

	if err != nil {
		log.Fatal(err)
	}

	// Cierra la conexión a la vase de datos cuando termine la función
	defer conn.Close(context.Background())

	log.Println("[MAIN] Database ready")

	// Creamos una instancia de nuestro "repository"
	repository := pokemon.NewRepository(conn)

	// Creamos una instancia para el service
	service := pokemon.NewService(repository)

	// Creamos una instancia de nuestro Handler
	// Ahora el Handler ya no se conecta directo al Repository, sino al Service
	handler := pokemon.NewHandler(service)

	// Con esto le indicamos al servidor que:
	// - "Cuando llegue una petición GET a 'pokemon/', ejecuta la función handler.GetAll()"
	http.HandleFunc("GET /pokemon", handler.GetAll)
	// - "Cuando llegue una petición POST a 'pokemon/', ejecuta la función handler.Create()"
	http.HandleFunc("POST /pokemon", handler.Create)

	log.Println("[MAIN] Server running on :8082")

	// // Ejecutamos el método GetAll() de nuestro repository
	// pokemons, err := repository.GetAll()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// // log.Printf("Found %d pokemon(s): ", len(pokemons))

	// // 1. Equivalente a escribir: "for (const pokemon of pokemons)".
	// // 2. "range" recorre todos los elementos de un Slice.
	// // 3. "range" devuelve dos valores: (index, value), por lo que el
	// //    guion bajo en esta sintaxis representaría al índice, pero como
	// //    para esto no nos interesa, usamos el guin bajo para indicar
	// //    que no lo vamos a usar.
	// for _, pokemon := range pokemons {
	// 	// - "%d" sirve para imprimir ENTEROS.
	// 	// - "%s" sirve para imprimir STRINGS (cadenas).
	// 	// - "%t" sirve para imprimir BOOLEANS.
	// 	log.Printf(
	// 		"ID: %d | Name: %s | Shiny: %t",
	// 		pokemon.ID,
	// 		pokemon.Name,
	// 		pokemon.Shiny,
	// 	)
	// }

	// 1. La función "ListenAndService" hace dos cosas en esta implementación:
	// - Abre el puerto 8080
	// - Se queda esperando peticiones
	//
	// 2. La función "ListenAndService" espera dos parámetros:
	// - El puerto a abrir
	// - El 'router' a utilizar
	//   Dejamos este segundo en NIL porque no queremos usa ninguno
	//   específico de momento. Después lo vamos a cambiar.
	err = http.ListenAndServe(":8082", nil)
	if err != nil {
		log.Fatal(err)
	}
}
