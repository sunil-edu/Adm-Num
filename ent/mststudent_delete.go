// Copyright 2021-present Eduate Pvt Ltd. All rights reserved.
// Code generated by entc, DO NOT EDIT.

package ent

import (
	"adm-num/ent/mststudent"
	"adm-num/ent/predicate"
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MstStudentDelete is the builder for deleting a MstStudent entity.
type MstStudentDelete struct {
	config
	hooks    []Hook
	mutation *MstStudentMutation
}

// Where appends a list predicates to the MstStudentDelete builder.
func (msd *MstStudentDelete) Where(ps ...predicate.MstStudent) *MstStudentDelete {
	msd.mutation.Where(ps...)
	return msd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (msd *MstStudentDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(msd.hooks) == 0 {
		affected, err = msd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*MstStudentMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			msd.mutation = mutation
			affected, err = msd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(msd.hooks) - 1; i >= 0; i-- {
			if msd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = msd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, msd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (msd *MstStudentDelete) ExecX(ctx context.Context) int {
	n, err := msd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (msd *MstStudentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: mststudent.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: mststudent.FieldID,
			},
		},
	}
	if ps := msd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, msd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// MstStudentDeleteOne is the builder for deleting a single MstStudent entity.
type MstStudentDeleteOne struct {
	msd *MstStudentDelete
}

// Exec executes the deletion query.
func (msdo *MstStudentDeleteOne) Exec(ctx context.Context) error {
	n, err := msdo.msd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{mststudent.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (msdo *MstStudentDeleteOne) ExecX(ctx context.Context) {
	msdo.msd.ExecX(ctx)
}
