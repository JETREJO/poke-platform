package pokemon

// Nuevo Struct de Pokemon creado para RECIBIR datos en la API
// Esto convierte en JSON a un STRUCT (gracias a 'json:')
type CreatePokemonRequest struct {
	Name            string  `json:"name"`
	PrimaryTypeID   int     `json:"primaryTypeId"`
	SecondaryTypeID *int    `json:"secondaryTypeId"`
	GenerationID    int     `json:"generationId"`
	Sprite          *string `json:"sprite"`
	Shiny           bool    `json:"shiny"`
}

// Se hace de esta forma porque NO queremos que el Frontend
// CONOZCA TODA nuestra arquitectura. Esto se hace para protegernos.
