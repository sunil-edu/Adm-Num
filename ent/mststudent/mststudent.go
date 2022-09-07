// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package mststudent

import (
	"time"

	"entgo.io/ent"
)

const (
	// Label holds the string label denoting the mststudent type in the database.
	Label = "mst_student"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldFirstName holds the string denoting the first_name field in the database.
	FieldFirstName = "first_name"
	// FieldMiddleName holds the string denoting the middle_name field in the database.
	FieldMiddleName = "middle_name"
	// FieldLastName holds the string denoting the last_name field in the database.
	FieldLastName = "last_name"
	// FieldStdStudying holds the string denoting the std_studying field in the database.
	FieldStdStudying = "std_studying"
	// FieldStdStatus holds the string denoting the std_status field in the database.
	FieldStdStatus = "std_status"
	// FieldStdSex holds the string denoting the std_sex field in the database.
	FieldStdSex = "std_sex"
	// FieldStdRegNo holds the string denoting the std_reg_no field in the database.
	FieldStdRegNo = "std_reg_no"
	// FieldStdAdmNo holds the string denoting the std_adm_no field in the database.
	FieldStdAdmNo = "std_adm_no"
	// FieldStdDoa holds the string denoting the std_doa field in the database.
	FieldStdDoa = "std_doa"
	// FieldStdFresher holds the string denoting the std_fresher field in the database.
	FieldStdFresher = "std_fresher"
	// FieldStdDob holds the string denoting the std_dob field in the database.
	FieldStdDob = "std_dob"
	// FieldStdEmail holds the string denoting the std_email field in the database.
	FieldStdEmail = "std_email"
	// FieldStdMobile holds the string denoting the std_mobile field in the database.
	FieldStdMobile = "std_mobile"
	// FieldStdFatherName holds the string denoting the std_father_name field in the database.
	FieldStdFatherName = "std_father_name"
	// FieldStdMotherName holds the string denoting the std_mother_name field in the database.
	FieldStdMotherName = "std_mother_name"
	// Table holds the table name of the mststudent in the database.
	Table = "mst_students"
)

// Columns holds all SQL columns for mststudent fields.
var Columns = []string{
	FieldID,
	FieldFirstName,
	FieldMiddleName,
	FieldLastName,
	FieldStdStudying,
	FieldStdStatus,
	FieldStdSex,
	FieldStdRegNo,
	FieldStdAdmNo,
	FieldStdDoa,
	FieldStdFresher,
	FieldStdDob,
	FieldStdEmail,
	FieldStdMobile,
	FieldStdFatherName,
	FieldStdMotherName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "adm-num/ent/runtime"
var (
	Hooks [1]ent.Hook
	// FirstNameValidator is a validator for the "first_name" field. It is called by the builders before save.
	FirstNameValidator func(string) error
	// DefaultMiddleName holds the default value on creation for the "middle_name" field.
	DefaultMiddleName string
	// DefaultLastName holds the default value on creation for the "last_name" field.
	DefaultLastName string
	// DefaultStdStudying holds the default value on creation for the "std_studying" field.
	DefaultStdStudying bool
	// DefaultStdStatus holds the default value on creation for the "std_status" field.
	DefaultStdStatus string
	// DefaultStdSex holds the default value on creation for the "std_sex" field.
	DefaultStdSex string
	// DefaultStdRegNo holds the default value on creation for the "std_reg_no" field.
	DefaultStdRegNo string
	// DefaultStdAdmNo holds the default value on creation for the "std_adm_no" field.
	DefaultStdAdmNo string
	// DefaultStdDoa holds the default value on creation for the "std_doa" field.
	DefaultStdDoa func() time.Time
	// DefaultStdFresher holds the default value on creation for the "std_fresher" field.
	DefaultStdFresher bool
	// DefaultStdDob holds the default value on creation for the "std_dob" field.
	DefaultStdDob func() time.Time
	// StdEmailValidator is a validator for the "std_email" field. It is called by the builders before save.
	StdEmailValidator func(string) error
	// StdMobileValidator is a validator for the "std_mobile" field. It is called by the builders before save.
	StdMobileValidator func(string) error
	// DefaultStdFatherName holds the default value on creation for the "std_father_name" field.
	DefaultStdFatherName string
	// DefaultStdMotherName holds the default value on creation for the "std_mother_name" field.
	DefaultStdMotherName string
)