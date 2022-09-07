// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AdmNumbersColumns holds the columns for the "adm_numbers" table.
	AdmNumbersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "adm_start_no", Type: field.TypeString, Default: "1"},
		{Name: "adm_current_no", Type: field.TypeString, Default: "0"},
		{Name: "is_prefixed", Type: field.TypeBool, Default: false},
		{Name: "prefix_str", Type: field.TypeString, Default: "Inst"},
		{Name: "suffix_str", Type: field.TypeString, Default: "GC"},
		{Name: "separator", Type: field.TypeString, Default: "-"},
		{Name: "prefill_with_zero", Type: field.TypeBool, Default: true},
		{Name: "prefill_width", Type: field.TypeInt, Default: 5},
	}
	// AdmNumbersTable holds the schema information for the "adm_numbers" table.
	AdmNumbersTable = &schema.Table{
		Name:       "adm_numbers",
		Columns:    AdmNumbersColumns,
		PrimaryKey: []*schema.Column{AdmNumbersColumns[0]},
	}
	// MstStudentsColumns holds the columns for the "mst_students" table.
	MstStudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "first_name", Type: field.TypeString},
		{Name: "middle_name", Type: field.TypeString, Default: ""},
		{Name: "last_name", Type: field.TypeString, Default: ""},
		{Name: "std_studying", Type: field.TypeBool, Default: true},
		{Name: "std_status", Type: field.TypeString, Default: "CUR"},
		{Name: "std_sex", Type: field.TypeString, Default: ""},
		{Name: "std_reg_no", Type: field.TypeString, Default: ""},
		{Name: "std_adm_no", Type: field.TypeString, Default: ""},
		{Name: "std_doa", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "DATE"}},
		{Name: "std_fresher", Type: field.TypeBool, Default: true},
		{Name: "std_dob", Type: field.TypeTime, Nullable: true, SchemaType: map[string]string{"postgres": "DATE"}},
		{Name: "std_email", Type: field.TypeString, Unique: true},
		{Name: "std_mobile", Type: field.TypeString},
		{Name: "std_father_name", Type: field.TypeString, Default: ""},
		{Name: "std_mother_name", Type: field.TypeString, Default: ""},
	}
	// MstStudentsTable holds the schema information for the "mst_students" table.
	MstStudentsTable = &schema.Table{
		Name:       "mst_students",
		Columns:    MstStudentsColumns,
		PrimaryKey: []*schema.Column{MstStudentsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AdmNumbersTable,
		MstStudentsTable,
	}
)

func init() {
}
