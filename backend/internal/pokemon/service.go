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

/*
 * ----------------------------------------------------------
 *          GET ALL POKEMONS
 * ----------------------------------------------------------
 */

func (s *Service) GetAll() ([]Pokemon, error) {
	return s.repository.GetAll()
}

/*
 * ----------------------------------------------------------
 *          CREATE ONE POKEMON
 * ----------------------------------------------------------
 */

func (s *Service) Create(request CreatePokemonRequest) (Pokemon, error) {
	// A esta asignación se le llama: 'COMPOSITE LITERAL'
	// No es mas que una inicialización de un Struct con datos.
	// Es como si inicializáramos un objeto en Typescript poniendo explícitamente
	// su tipo y llenando cada campo con su valor.
	// *
	// NOTA: En GO los nombres de los campos siempre empiezan con MAYÚSCULA porque estos se EXPORTAN,
	// si se escribira con MINÚSCULA, entonces serían variables PRIVADOS.
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

/*
 * ----------------------------------------------------------
 *          SEARCH ONE POKEMON BY ID
 * ----------------------------------------------------------
 */

func (s *Service) GetByID(id int) (Pokemon, error) {
	return s.repository.GetByID(id)
}

/*
 * ----------------------------------------------------------
 *          UPDATE ONE POKEMON
 * ----------------------------------------------------------
 */

func (s *Service) Update(id int, request UpdatePokemonRequest) (Pokemon, error) {

	// Usamos la función de GetById() primero para obtener
	// los datos del pokemon que vamos a actualizar, esto con el
	// fin de saber si el pokemon EXISTE o NO.
	pokemon, err := s.repository.GetByID(id)
	if err != nil {
		return Pokemon{}, err
	}

	pokemon.Name = request.Name
	pokemon.PrimaryTypeID = request.PrimaryTypeID
	pokemon.SecondaryTypeID = request.SecondaryTypeID
	pokemon.GenerationID = request.GenerationID
	pokemon.Sprite = request.Sprite
	pokemon.Shiny = request.Shiny

	err = s.repository.Update(pokemon)
	if err != nil {
		return Pokemon{}, err
	}

	return pokemon, nil
}

/*
 * ----------------------------------------------------------
 *          DELETE ONE POKEMON
 * ----------------------------------------------------------
 */

func (s *Service) Delete(id int) error {

	_, err := s.repository.GetByID(id)

	if err != nil {
		return err
	}

	return s.repository.Delete(id)
}
