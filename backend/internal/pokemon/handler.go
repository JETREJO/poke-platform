package pokemon

// Importamos nuevas librerías
import (
	"encoding/json"
	"net/http"
)

// Este es el Struct de nuestro handler
type Handler struct {
	repository *Repository
}

// Este es el constructor.
func NewHandler(repository *Repository) *Handler {
	return &Handler{
		repository: repository,
	}
}

// 1. Esta función pertenece al Struct "Handler"
//
// 2. "http.ResponseWritter" representa la respuesta. Todo lo que
// se escriba en el "w.Write(...)" se enviará al navegador.
//
// 3. "*http.Request" representa la petición. Aquí se pueden leer
// cosas como params, headers, body, cookies, etc.
func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {

	pokemons, err := h.repository.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Header para indicar que vamos a devolver un JSON
	w.Header().Set("Content-Type", "application/json")

	// El enconder transoforma todo para devolverlo listo en la respuesta.
	// Esto nos evita:
	// - Recorrer el Slice
	// - Construir el JSON
	// - Convertir strings
	json.NewEncoder(w).Encode(pokemons)
}
