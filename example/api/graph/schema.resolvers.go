package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/RiskyFeryansyahP/introduction-graphql/graph/generated"
	"github.com/RiskyFeryansyahP/introduction-graphql/graph/model"
	"github.com/google/uuid"
)

var todoItem []*model.Todo

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	todo := &model.Todo{
		ID:   id.String(),
		Text: input.Text,
		Done: false,
	}

	todoItem = append(todoItem, todo)

	return todo, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, id string) (*model.Todo, error) {
	for _, v := range todoItem {
		if v.ID == id {
			v.Done = !v.Done

			return v, nil
		}
	}

	return nil, fmt.Errorf("data not found")
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, id string) (*model.Todo, error) {
	for k, v := range todoItem {
		if v.ID == id {
			todoItem = append(todoItem[:k], todoItem[k+1:]...)

			return v, nil
		}
	}

	return nil, fmt.Errorf("data not found")
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return todoItem, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
