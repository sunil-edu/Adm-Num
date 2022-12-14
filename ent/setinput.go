// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"
)

// CreateAdmNumberInput represents a mutation input for creating admnumbers.
type CreateAdmNumberInput struct {
	AdmStartNo      *string
	AdmCurrentNo    *string
	IsPrefixed      *bool
	PrefixStr       *string
	SuffixStr       *string
	Separator       *string
	PrefillWithZero *bool
	PrefillWidth    *int
}

// Mutate applies the CreateAdmNumberInput on the AdmNumberCreate builder.
func (i *CreateAdmNumberInput) Mutate(m *AdmNumberCreate) {
	if v := i.AdmStartNo; v != nil {
		m.SetAdmStartNo(*v)
	}
	if v := i.AdmCurrentNo; v != nil {
		m.SetAdmCurrentNo(*v)
	}
	if v := i.IsPrefixed; v != nil {
		m.SetIsPrefixed(*v)
	}
	if v := i.PrefixStr; v != nil {
		m.SetPrefixStr(*v)
	}
	if v := i.SuffixStr; v != nil {
		m.SetSuffixStr(*v)
	}
	if v := i.Separator; v != nil {
		m.SetSeparator(*v)
	}
	if v := i.PrefillWithZero; v != nil {
		m.SetPrefillWithZero(*v)
	}
	if v := i.PrefillWidth; v != nil {
		m.SetPrefillWidth(*v)
	}
}

// SetInput applies the change-set in the CreateAdmNumberInput on the create builder.
func (c *AdmNumberCreate) SetInput(i CreateAdmNumberInput) *AdmNumberCreate {
	i.Mutate(c)
	return c
}

// UpdateAdmNumberInput represents a mutation input for updating admnumbers.
type UpdateAdmNumberInput struct {
	AdmStartNo      *string
	AdmCurrentNo    *string
	IsPrefixed      *bool
	PrefixStr       *string
	SuffixStr       *string
	Separator       *string
	PrefillWithZero *bool
	PrefillWidth    *int
}

// Mutate applies the UpdateAdmNumberInput on the AdmNumberMutation.
func (i *UpdateAdmNumberInput) Mutate(m *AdmNumberMutation) {
	if v := i.AdmStartNo; v != nil {
		m.SetAdmStartNo(*v)
	}
	if v := i.AdmCurrentNo; v != nil {
		m.SetAdmCurrentNo(*v)
	}
	if v := i.IsPrefixed; v != nil {
		m.SetIsPrefixed(*v)
	}
	if v := i.PrefixStr; v != nil {
		m.SetPrefixStr(*v)
	}
	if v := i.SuffixStr; v != nil {
		m.SetSuffixStr(*v)
	}
	if v := i.Separator; v != nil {
		m.SetSeparator(*v)
	}
	if v := i.PrefillWithZero; v != nil {
		m.SetPrefillWithZero(*v)
	}
	if v := i.PrefillWidth; v != nil {
		m.SetPrefillWidth(*v)
	}
}

// SetInput applies the change-set in the UpdateAdmNumberInput on the update builder.
func (u *AdmNumberUpdate) SetInput(i UpdateAdmNumberInput) *AdmNumberUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateAdmNumberInput on the update-one builder.
func (u *AdmNumberUpdateOne) SetInput(i UpdateAdmNumberInput) *AdmNumberUpdateOne {
	i.Mutate(u.Mutation())
	return u
}

// CreateMstStudentInput represents a mutation input for creating mststudents.
type CreateMstStudentInput struct {
	FirstName     string
	MiddleName    *string
	LastName      *string
	StdStudying   *bool
	StdStatus     *string
	StdSex        *string
	StdRegNo      *string
	StdAdmNo      *string
	StdDoa        *time.Time
	StdFresher    *bool
	StdDob        *time.Time
	StdEmail      string
	StdMobile     string
	StdFatherName *string
	StdMotherName *string
}

// Mutate applies the CreateMstStudentInput on the MstStudentCreate builder.
func (i *CreateMstStudentInput) Mutate(m *MstStudentCreate) {
	m.SetFirstName(i.FirstName)
	if v := i.MiddleName; v != nil {
		m.SetMiddleName(*v)
	}
	if v := i.LastName; v != nil {
		m.SetLastName(*v)
	}
	if v := i.StdStudying; v != nil {
		m.SetStdStudying(*v)
	}
	if v := i.StdStatus; v != nil {
		m.SetStdStatus(*v)
	}
	if v := i.StdSex; v != nil {
		m.SetStdSex(*v)
	}
	if v := i.StdRegNo; v != nil {
		m.SetStdRegNo(*v)
	}
	if v := i.StdAdmNo; v != nil {
		m.SetStdAdmNo(*v)
	}
	if v := i.StdDoa; v != nil {
		m.SetStdDoa(*v)
	}
	if v := i.StdFresher; v != nil {
		m.SetStdFresher(*v)
	}
	if v := i.StdDob; v != nil {
		m.SetStdDob(*v)
	}
	m.SetStdEmail(i.StdEmail)
	m.SetStdMobile(i.StdMobile)
	if v := i.StdFatherName; v != nil {
		m.SetStdFatherName(*v)
	}
	if v := i.StdMotherName; v != nil {
		m.SetStdMotherName(*v)
	}
}

// SetInput applies the change-set in the CreateMstStudentInput on the create builder.
func (c *MstStudentCreate) SetInput(i CreateMstStudentInput) *MstStudentCreate {
	i.Mutate(c)
	return c
}

// UpdateMstStudentInput represents a mutation input for updating mststudents.
type UpdateMstStudentInput struct {
	FirstName     *string
	MiddleName    *string
	LastName      *string
	StdStudying   *bool
	StdStatus     *string
	StdSex        *string
	StdRegNo      *string
	StdAdmNo      *string
	StdDoa        *time.Time
	ClearStdDoa   bool
	StdFresher    *bool
	StdDob        *time.Time
	ClearStdDob   bool
	StdEmail      *string
	StdMobile     *string
	StdFatherName *string
	StdMotherName *string
}

// Mutate applies the UpdateMstStudentInput on the MstStudentMutation.
func (i *UpdateMstStudentInput) Mutate(m *MstStudentMutation) {
	if v := i.FirstName; v != nil {
		m.SetFirstName(*v)
	}
	if v := i.MiddleName; v != nil {
		m.SetMiddleName(*v)
	}
	if v := i.LastName; v != nil {
		m.SetLastName(*v)
	}
	if v := i.StdStudying; v != nil {
		m.SetStdStudying(*v)
	}
	if v := i.StdStatus; v != nil {
		m.SetStdStatus(*v)
	}
	if v := i.StdSex; v != nil {
		m.SetStdSex(*v)
	}
	if v := i.StdRegNo; v != nil {
		m.SetStdRegNo(*v)
	}
	if v := i.StdAdmNo; v != nil {
		m.SetStdAdmNo(*v)
	}
	if i.ClearStdDoa {
		m.ClearStdDoa()
	}
	if v := i.StdDoa; v != nil {
		m.SetStdDoa(*v)
	}
	if v := i.StdFresher; v != nil {
		m.SetStdFresher(*v)
	}
	if i.ClearStdDob {
		m.ClearStdDob()
	}
	if v := i.StdDob; v != nil {
		m.SetStdDob(*v)
	}
	if v := i.StdEmail; v != nil {
		m.SetStdEmail(*v)
	}
	if v := i.StdMobile; v != nil {
		m.SetStdMobile(*v)
	}
	if v := i.StdFatherName; v != nil {
		m.SetStdFatherName(*v)
	}
	if v := i.StdMotherName; v != nil {
		m.SetStdMotherName(*v)
	}
}

// SetInput applies the change-set in the UpdateMstStudentInput on the update builder.
func (u *MstStudentUpdate) SetInput(i UpdateMstStudentInput) *MstStudentUpdate {
	i.Mutate(u.Mutation())
	return u
}

// SetInput applies the change-set in the UpdateMstStudentInput on the update-one builder.
func (u *MstStudentUpdateOne) SetInput(i UpdateMstStudentInput) *MstStudentUpdateOne {
	i.Mutate(u.Mutation())
	return u
}
