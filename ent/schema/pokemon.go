package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

type Pokemon struct {
	ent.Schema
}

func (Pokemon) Fields() []ent.Field {
	return []ent.Field{
		field.Int("pokedex"),
		field.String("name"),
		field.String("type_1"),
		field.String("type_2"),
		field.Int("total"),
		field.Int("hp"),
		field.Int("attack"),
		field.Int("defense"),
		field.Int("sp_attack"),
		field.Int("sp_defense"),
		field.Int("speed"),
		field.Int("generation"),
		field.Bool("legendary"),
	}
}

func (Pokemon) Edges() []ent.Edge {
	return []ent.Edge{}
}
