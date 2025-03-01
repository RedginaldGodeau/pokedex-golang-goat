package pokemon

type PokemonResponse struct {
	ID         int    `json:"id"`
	Pokedex    int    `json:"pokedex"`
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
