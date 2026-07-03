package pokemon

// Nuevo Struct de Pokemon creado para RECIBIR datos en la API
//   - Esto convierte en JSON a un STRUCT (gracias a 'json:')
//   - Se hace de esta forma porque NO queremos que el Frontend
//     CONOZCA TODA nuestra arquitectura. Esto se hace para protegernos.
type CreatePokemonRequest struct {
	Name            string  `json:"name"`
	PrimaryTypeID   int     `json:"primaryTypeId"`
	SecondaryTypeID *int    `json:"secondaryTypeId"`
	GenerationID    int     `json:"generationId"`
	Sprite          *string `json:"sprite"`
	Shiny           bool    `json:"shiny"`
}

// De momento es igual al Struct de Crear, pero puede que estos
// dos Structs sean distintos. Recordemos que cada uno va a tener
// un uso distinto, por lo que es mejor tenerlos separados.
type UpdatePokemonRequest struct {
	Name            string  `json:"name"`
	PrimaryTypeID   int     `json:"primaryTypeId"`
	SecondaryTypeID *int    `json:"secondaryTypeId"`
	GenerationID    int     `json:"generationId"`
	Sprite          *string `json:"sprite"`
	Shiny           bool    `json:"shiny"`
}
