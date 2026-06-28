package main

import (
	"context"
	"log"

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

	// Ejecutamos el método GetAll() de nuestro repository
	pokemons, err := repository.GetAll()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Found %d pokemon(s): ", len(pokemons))

	// 1. Equivalente a escribir: "for (const pokemon of pokemons)".
	// 2. "range" recorre todos los elementos de un Slice.
	// 3. "range" devuelve dos valores: (index, value), por lo que el
	//    guion bajo en esta sintaxis representaría al índice, pero como
	//    para esto no nos interesa, usamos el guin bajo para indicar
	//    que no lo vamos a usar.
	for _, pokemon := range pokemons {
		// - "%d" sirve para imprimir ENTEROS.
		// - "%s" sirve para imprimir STRINGS (cadenas).
		// - "%t" sirve para imprimir BOOLEANS.
		log.Printf(
			"ID: %d | Name: %s | Shiny: %t",
			pokemon.ID,
			pokemon.Name,
			pokemon.Shiny,
		)
	}
}
