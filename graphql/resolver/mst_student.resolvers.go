package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"adm-num/ent"
	"adm-num/graphql/generated"
	"context"
)

// AddStudent is the resolver for the AddStudent field.
func (r *mutationResolver) AddStudent(ctx context.Context, input ent.CreateMstStudentInput) (*ent.MstStudent, error) {
	client := ent.FromContext(ctx)
	return client.MstStudent.Create().SetInput(input).Save(ctx)
}

// GetStudents is the resolver for the GetStudents field.
func (r *queryResolver) GetStudents(ctx context.Context) ([]*ent.MstStudent, error) {
	return r.client.MstStudent.Query().All(ctx)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
