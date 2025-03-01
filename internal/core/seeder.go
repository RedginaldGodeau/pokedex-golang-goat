package core

import (
	"backend/ent/gen"
	"context"
	"encoding/json"
	"os"
)

type pokemonData struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Type1      string `json:"type_1"`
	Type2      string `json:"type_2"`
	Total      int    `json:"total"`
	HP         int    `json:"hp"`
	Attack     int    `json:"attack"`
	Defense    int    `json:"defense"`
	SpAttack   int    `json:"sp_attack"`
	SpDefense  int    `json:"sp_defense"`
	Speed      int    `json:"speed"`
	Generation int    `json:"generation"`
	Legendary  bool   `json:"legendary"`
}

func InitSeeder(db *gen.Client) error {
	count, err := db.Pokemon.Query().Count(context.TODO())
	if err != nil {
		return err
	}

	if count > 0 {
		return nil
	}

	content, err := os.ReadFile("data/pokemon.json")
	if err != nil {
		return err
	}

	var pokemons []pokemonData
	err = json.Unmarshal(content, &pokemons)
	if err != nil {
		return err
	}

	tx, err := db.Tx(context.TODO())
	if err != nil {
		return err
	}

	for _, pokemon := range pokemons {
		err = tx.Pokemon.Create().
			SetPokedex(pokemon.Id).
			SetName(pokemon.Name).
			SetType1(pokemon.Type1).
			SetType2(pokemon.Type2).
			SetTotal(pokemon.Total).
			SetHp(pokemon.HP).
			SetAttack(pokemon.Attack).
			SetDefense(pokemon.Defense).
			SetSpAttack(pokemon.SpAttack).
			SetSpDefense(pokemon.SpDefense).
			SetSpeed(pokemon.Speed).
			SetGeneration(pokemon.Generation).
			SetLegendary(pokemon.Legendary).Exec(context.TODO())
		if err != nil {
			tx.Rollback()
			return err
		}
	}
	return tx.Commit()
}
