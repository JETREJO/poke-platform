package pokemon

type Service struct {
	repository *Repository
}

// Constructor
func NewService(repository *Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetAll() ([]Pokemon, error) {
	return s.repository.GetAll()
}

func (s *Service) Create(request CreatePokemonRequest) (Pokemon, error) {
	// A esta asignación se le llama: 'COMPOSITE LITERAL'
	// No es mas que una inicialización de un Struct con datos.
	// Es como si inicializáramos un objeto en Typescript poniendo explícitamente
	// su tipo y llenando cada campo con su valor.
	// *
	// NOTA: En GO los nombres de los campos siempre empiezan con MAYÚSCULA.
	// *
	pokemon := Pokemon{
		Name:            request.Name,
		PrimaryTypeID:   request.PrimaryTypeID,
		SecondaryTypeID: request.SecondaryTypeID,
		GenerationID:    request.GenerationID,
		Sprite:          request.Sprite,
		Shiny:           request.Shiny,
	}
	return s.repository.Create(pokemon)
}
