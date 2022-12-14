package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/mourajj/graphql-go/data"
	"github.com/mourajj/graphql-go/graph/generated"
	"github.com/mourajj/graphql-go/graph/model"
)

// CreateDog is the resolver for the createDog field.
func (r *mutationResolver) CreateDog(ctx context.Context, input *model.NewDog) (*model.Dog, error) {
	newDog := model.Dog{
		ID:        "test",
		Name:      input.Name,
		IsGoodBoi: input.IsGoodBoi,
	}

	data.Dogs = append(data.Dogs, &newDog)
	return &newDog, nil
}

// Dogs is the resolver for the dogs field.
func (r *queryResolver) Dogs(ctx context.Context) ([]*model.Dog, error) {
	return data.Dogs, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
