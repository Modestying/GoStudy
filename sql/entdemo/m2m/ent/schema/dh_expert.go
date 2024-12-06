package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DhExpert struct {
	ent.Schema
}

func (DhExpert) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name"),
	}
}

func (DhExpert) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("class_relation", DhExpertClass.Type).Unique(),
		edge.To("region_relation", DhRegion.Type).Unique(),
		edge.To("title_relation", DhExpertTitle.Type),
	}
}

func (DhExpert) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "dh_expert"}}
}
