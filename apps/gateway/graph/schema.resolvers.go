package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gateway/di"
	"gateway/graph/generated"
	"gateway/graph/model"
	"proto"
)

// CreateTask is the resolver for the createTask field.
func (r *mutationResolver) CreateTask(ctx context.Context, input model.CreateTaskInput) (*model.Task, error) {
	return r.Query().Task(ctx, "1")
}

// Tasks is the resolver for the tasks field.
func (r *queryResolver) Tasks(ctx context.Context, ownerID string) ([]*model.Task, error) {
	var tasks []*model.Task = nil

	err := di.Invoke(func(sc *di.ServiceClient) error {
		resp, err := sc.Tasks.ListTasks(ctx, &proto.Empty{}) // TODO: implement tasks filtering and pagination
		if err != nil {
			return err
		}

		tasks = make([]*model.Task, len(resp.Tasks))

		for i, t := range resp.Tasks {
			status, err := t.Status.ToString()
			if err != nil {
				return err
			}

			tasks[i] = &model.Task{
				ID:     t.Id,
				Title:  t.Title,
				Status: model.TaskStatus(status),
				Owner:  &model.User{ID: t.OwnerId},
			}
		}

		return nil
	})

	return tasks, err
}

// Task is the resolver for the task field.
func (r *queryResolver) Task(ctx context.Context, id string) (*model.Task, error) {
	var task *model.Task = nil

	err := di.Invoke(func(sc *di.ServiceClient) error {
		resp, err := sc.Tasks.GetTask(ctx, &proto.GetTaskRequest{
			Id: id,
		})
		if err != nil {
			return err
		}

		task = &model.Task{
			ID:        resp.Id,
			Title:     resp.Title,
			Status:    model.TaskStatusActive,
			Owner:     &model.User{ID: resp.OwnerId},
			CreatedAt: resp.CreatedAt,
			UpdatedAt: resp.UpdatedAt,
		}
		return nil
	})

	return task, err
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
