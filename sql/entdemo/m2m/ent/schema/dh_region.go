package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

type DhRegion struct {
	ent.Schema
}

func (DhRegion) Fields() []ent.Field {
	return []ent.Field{
		field.Uint64("id").Unique(),
		field.String("name").Comment("区域名称"),
		field.String("code"),
	}
}
func (DhRegion) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("expert_region", DhExpert.Type).
			Ref("region_relation"),
	}
}
func (DhRegion) Annotations() []schema.Annotation {
	return []schema.Annotation{entsql.Annotation{Table: "dh_region"}}
}
