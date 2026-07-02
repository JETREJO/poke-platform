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
	return s.repository.Create(request)
}
