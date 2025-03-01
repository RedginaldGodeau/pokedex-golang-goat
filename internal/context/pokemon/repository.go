package pokemon

import (
	"backend/ent/gen"
	"backend/ent/gen/pokemon"
	"backend/pkg/application"
	"context"
	"strings"
)

type RepositoryInterface interface {
	GetByID(id int) (*gen.Pokemon, error)
	GetByPokedexID(id int) ([]*gen.Pokemon, error)
	GetByName(name string) (*gen.Pokemon, error)
	GetBySearch(search string) ([]*gen.Pokemon, error)
	GetAll() ([]*gen.Pokemon, error)
}

type Repository struct {
	db *gen.Client
}

func NewRepository(db *gen.Client) RepositoryInterface {
	return &Repository{db: db}
}
func NewRepositoryFromApp(app *application.Application) RepositoryInterface {
	return NewRepository(app.DB)
}

func (r Repository) GetByPokedexID(id int) ([]*gen.Pokemon, error) {
	_pokemons, err := r.db.Pokemon.Query().Where(pokemon.PokedexEQ(id)).All(context.TODO())
	if err != nil {
		return nil, err
	}

	return _pokemons, nil
}
func (r Repository) GetByID(id int) (*gen.Pokemon, error) {
	_pokemon, err := r.db.Pokemon.Query().Where(pokemon.ID(id)).Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return _pokemon, nil
}

func (r Repository) GetByName(name string) (*gen.Pokemon, error) {
	_pokemon, err := r.db.Pokemon.Query().Where(pokemon.Name(name)).Only(context.TODO())
	if err != nil {
		return nil, err
	}

	return _pokemon, nil
}

func (r Repository) GetBySearch(search string) ([]*gen.Pokemon, error) {
	_pokemons, err := r.db.Pokemon.Query().Where(pokemon.NameHasPrefix(strings.ToUpper(string(search[0])) + strings.ToLower(search[1:]))).All(context.TODO())
	if err != nil {
		return nil, err
	}

	return _pokemons, nil
}

func (r Repository) GetAll() ([]*gen.Pokemon, error) {
	_pokemons, err := r.db.Pokemon.Query().All(context.TODO())
	if err != nil {
		return nil, err
	}

	return _pokemons, nil
}
