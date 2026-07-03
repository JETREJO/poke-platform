package http

import (
	"net/http"

	"pokemon-platform/backend/internal/pokemon"
)

func NewRouter(pokemonHandler *pokemon.Handler) *http.ServeMux {
	// ServeMux será nuestro enrutador explícito y aquí manualmente
	// vamos a registrar las rutas de nuestra API.

	router := http.NewServeMux()

	// Con esto le indicamos al servidor que:
	// - "Cuando llegue una petición GET a 'pokemon/', ejecuta la función handler.GetAll()"

	router.HandleFunc("GET /pokemon", pokemonHandler.GetAll)
	router.HandleFunc("GET /pokemon/{id}", pokemonHandler.GetByID)
	router.HandleFunc("POST /pokemon", pokemonHandler.Create)

	return router
}
