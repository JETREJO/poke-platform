package pokemon

// Importamos nuevas librerías
import (
	"encoding/json"
	"net/http"
)

// Este es el Struct de nuestro handler
type Handler struct {
	service *Service
}

// Este es el constructor.
func NewHandler(service *Service) *Handler {
	return &Handler{
		service: service,
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

	pokemons, err := h.service.GetAll()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Header para indicar que vamos a devolver un JSON
	w.Header().Set("Content-Type", "application/json")

	// El enconder transforma todo para devolverlo listo en la respuesta.
	// Esto nos evita:
	// - Recorrer el Slice
	// - Construir el JSON
	// - Convertir strings
	err = json.NewEncoder(w).Encode(pokemons)

	// Agregamos esta validación ya que 'Encode' también puede fallar, es raro
	// pero no es imposible. Por precaución y buena práctica, hay que validar
	// errores en todos los posibles fallos.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {

	var request CreatePokemonRequest // El tipo viene del DTO

	// "r.body" es el body que recibimos de la request, en este caso es un JSON.
	// Se le asigna a "err" para validar que el Decode no tenga errores.
	// - El Decode puede funcionar correctamente para decodeificar el body gracias
	//   a las líneas que agregamos en Struct del DTO (de CreatePokemonRequest).
	// - Usamos "&request" ya que del Decoder necesita ESCRIBIR dentro del
	//   Struct "request". (Si le pasáramos solo 'request' entonces estaría
	//   trabajando sobre una copia y no podría rellenar los campos originales).
	err := json.NewDecoder(r.Body).Decode(&request)

	// Si el body (el JSON) que recibimos viene mal formateado, entonces
	// regresamos un status de error. (En este caso: 404 bad request).
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	// Esto es lo que regreso de la función.
	// Solo regreso el mismo JSON codificado nuevamente (de momento).
	err = json.NewEncoder(w).Encode(request)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
