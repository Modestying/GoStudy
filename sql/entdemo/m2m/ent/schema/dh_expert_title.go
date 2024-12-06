package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DhExpertTitle struct {
	ent.Schema
}

func (DhExpertTitle) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Comment("专家职称"),
	}
}
func (DhExpertTitle) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("expert_title", DhExpert.Type).
			Ref("title_relation"),
	}
}
func (DhExpertTitle) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "dh_expert_title"}}
}
