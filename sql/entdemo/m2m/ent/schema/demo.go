package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Demo struct {
	ent.Schema
}

func (Demo) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name"),
		field.Bool("is_delete"),
	}
}

func (Demo) Edges() []ent.Edge {
	return []ent.Edge{}
}

func (DhExpert) Demo() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "demo"}}
}
