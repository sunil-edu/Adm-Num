package schema

import (
	"adm-num/ent/hook"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema/field"
)

// MstStudent holds the schema definition for the MstStudent entity.
type MstStudent struct {
	ent.Schema
}

// Fields of the MstStudent.
func (MstStudent) Fields() []ent.Field {
	return []ent.Field{

		field.String("first_name").
			NotEmpty(),

		field.String("middle_name").Default(""),

		field.String("last_name").Default(""),

		field.Bool("std_studying").Default(true),

		field.String("std_status").Default("CUR"),

		field.String("std_sex").Default(""),

		field.String("std_reg_no").Default(""),

		field.String("std_adm_no").Default(""),

		field.Time("std_doa").Default(func() time.Time { return time.Time{} }).Optional().Nillable().
			SchemaType(map[string]string{dialect.Postgres: "DATE"}),

		field.Bool("std_fresher").Default(true),

		field.Time("std_dob").Default(func() time.Time { return time.Time{} }).Optional().Nillable().
			SchemaType(map[string]string{dialect.Postgres: "DATE"}),

		field.String("std_email").Unique().NotEmpty(),

		field.String("std_mobile").NotEmpty(),

		field.String("std_father_name").Default(""),

		field.String("std_mother_name").Default(""),
	}
}

// Edges of the MstStudent.
func (MstStudent) Edges() []ent.Edge {
	return nil
}

// Hooks of the MstStudent.
func (MstStudent) Hooks() []ent.Hook {
	return []ent.Hook{

		hook.On(MstStudentGetAdmissionNumber, ent.OpCreate),
	}
}
