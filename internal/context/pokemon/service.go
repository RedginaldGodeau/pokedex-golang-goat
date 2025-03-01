package pokemon

import "backend/pkg/application"

type ServiceInterface interface {
	GetByID(id int) (*PokemonResponse, error)
	GetByPokedex(id int) ([]PokemonResponse, error)
	GetByName(name string) (*PokemonResponse, error)
	GetBySearch(search string) ([]PokemonResponse, error)
	GetAll() ([]PokemonResponse, error)
}

type Service struct {
	repository RepositoryInterface
}

func NewService(repository RepositoryInterface) ServiceInterface {
	return &Service{
		repository: repository,
	}
}

func NewServiceFromApp(app *application.Application) ServiceInterface {
	return NewService(NewRepositoryFromApp(app))
}

func (s Service) GetByID(id int) (*PokemonResponse, error) {
	_pokemon, err := s.repository.GetByID(id)
	if err != nil {
		return nil, err
	}

	return &PokemonResponse{
		ID:         _pokemon.ID,
		Pokedex:    _pokemon.Pokedex,
		Name:       _pokemon.Name,
		Type1:      _pokemon.Type1,
		Type2:      _pokemon.Type2,
		Total:      _pokemon.Total,
		HP:         _pokemon.Hp,
		Attack:     _pokemon.Attack,
		Defense:    _pokemon.Defense,
		SpAttack:   _pokemon.SpAttack,
		SpDefense:  _pokemon.SpDefense,
		Speed:      _pokemon.Speed,
		Generation: _pokemon.Generation,
		Legendary:  _pokemon.Legendary,
	}, nil
}

func (s Service) GetByPokedex(id int) ([]PokemonResponse, error) {
	_pokemons, err := s.repository.GetByPokedexID(id)
	if err != nil {
		return nil, err
	}

	dto := make([]PokemonResponse, len(_pokemons))
	for i, _pokemon := range _pokemons {
		dto[i] = PokemonResponse{
			ID:         _pokemon.ID,
			Pokedex:    _pokemon.Pokedex,
			Name:       _pokemon.Name,
			Type1:      _pokemon.Type1,
			Type2:      _pokemon.Type2,
			Total:      _pokemon.Total,
			HP:         _pokemon.Hp,
			Attack:     _pokemon.Attack,
			Defense:    _pokemon.Defense,
			SpAttack:   _pokemon.SpAttack,
			SpDefense:  _pokemon.SpDefense,
			Speed:      _pokemon.Speed,
			Generation: _pokemon.Generation,
			Legendary:  _pokemon.Legendary,
		}
	}

	return dto, nil
}

func (s Service) GetByName(name string) (*PokemonResponse, error) {
	_pokemon, err := s.repository.GetByName(name)
	if err != nil {
		return nil, err
	}

	return &PokemonResponse{
		ID:         _pokemon.ID,
		Pokedex:    _pokemon.Pokedex,
		Name:       _pokemon.Name,
		Type1:      _pokemon.Type1,
		Type2:      _pokemon.Type2,
		Total:      _pokemon.Total,
		HP:         _pokemon.Hp,
		Attack:     _pokemon.Attack,
		Defense:    _pokemon.Defense,
		SpAttack:   _pokemon.SpAttack,
		SpDefense:  _pokemon.SpDefense,
		Speed:      _pokemon.Speed,
		Generation: _pokemon.Generation,
		Legendary:  _pokemon.Legendary,
	}, nil
}

func (s Service) GetBySearch(search string) ([]PokemonResponse, error) {
	_pokemons, err := s.repository.GetBySearch(search)
	if err != nil {
		return nil, err
	}

	dto := make([]PokemonResponse, len(_pokemons))
	for i, _pokemon := range _pokemons {
		dto[i] = PokemonResponse{
			ID:         _pokemon.ID,
			Pokedex:    _pokemon.Pokedex,
			Name:       _pokemon.Name,
			Type1:      _pokemon.Type1,
			Type2:      _pokemon.Type2,
			Total:      _pokemon.Total,
			HP:         _pokemon.Hp,
			Attack:     _pokemon.Attack,
			Defense:    _pokemon.Defense,
			SpAttack:   _pokemon.SpAttack,
			SpDefense:  _pokemon.SpDefense,
			Speed:      _pokemon.Speed,
			Generation: _pokemon.Generation,
			Legendary:  _pokemon.Legendary,
		}
	}

	return dto, nil
}

func (s Service) GetAll() ([]PokemonResponse, error) {
	_pokemons, err := s.repository.GetAll()
	if err != nil {
		return nil, err
	}

	dto := make([]PokemonResponse, len(_pokemons))
	for i, _pokemon := range _pokemons {
		dto[i] = PokemonResponse{
			ID:         _pokemon.ID,
			Pokedex:    _pokemon.Pokedex,
			Name:       _pokemon.Name,
			Type1:      _pokemon.Type1,
			Type2:      _pokemon.Type2,
			Total:      _pokemon.Total,
			HP:         _pokemon.Hp,
			Attack:     _pokemon.Attack,
			Defense:    _pokemon.Defense,
			SpAttack:   _pokemon.SpAttack,
			SpDefense:  _pokemon.SpDefense,
			Speed:      _pokemon.Speed,
			Generation: _pokemon.Generation,
			Legendary:  _pokemon.Legendary,
		}
	}

	return dto, nil
}
