package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"app/graph/generated"
	"app/graph/model"
	"context"
	"time"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	task := model.Task{
		Title:     input.Title,
		Note:      input.Note,
		Completed: 0,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	r.DB.Create(&task)

	return &task, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
