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

// El HANDLER no es una función que devuelva datos, envía directo la respuesta  al cliente.
// La respuesta del Handler viene en "W" -> HttpResponseWriter
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

	// Invocamos al método Create() del Service
	pokemon, err := h.service.Create(request)

	// Retornamos un error del servidor en caso de que la query no se haya
	// ejecutado correctamente.
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Seteamos los headers a la respuesta que vamos a enviar.
	w.Header().Set("Content-Type", "application/json")
	// - Este seta el status de creado (201 Created).
	// - Si no se hace esto, Go responde autompaticamente con un '200 OK'
	// - Siempre escribir el 'WriteHeader' primero antes que el body, ya que
	//   una vez que se empieza a escribir el body, el código HTTP ya no
	//   puede cambiar. Esto es un detalle de Go.
	w.WriteHeader(http.StatusCreated)

	// Esto es lo que regreso de la función.
	// - Todo lo que se 'escriba' en nuestra variable 'w' será lo que se
	//   envíe DIRECTO al cliente. Esta función no retorna datos, se comunica
	//   directo con el cliente conforme lo que le escribamos en 'w'.
	// - Regresamos el Pokemon creado en formato JSON.
	err = json.NewEncoder(w).Encode(pokemon)

	// if err != nil {
	// 	http.Error(w, err.Error(), http.StatusInternalServerError)
	// }
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
