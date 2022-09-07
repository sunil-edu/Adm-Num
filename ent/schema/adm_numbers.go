package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// AdmNumber holds the schema definition for the AdmNumber entity.
type AdmNumber struct {
	ent.Schema
}

// Fields of the AdmNumber.
func (AdmNumber) Fields() []ent.Field {
	return []ent.Field{

		field.String("adm_start_no").Default("1"),
		field.String("adm_current_no").Default("0"),
		field.Bool("is_prefixed").Default(false),
		field.String("prefix_str").Default("Inst"),
		field.String("suffix_str").Default("GC"),
		field.String("separator").Default("-"),
		field.Bool("prefill_with_zero").Default(true),
		field.Int("prefill_width").Default(5).Min(0).Max(5),
	}
}
