package impl

import (
	"context"
	"fmt"
	"proto"
	"strconv"
	"task-service/ent"
	"time"
	"util"
)

type TasksServer struct {
	proto.UnimplementedTaskServiceServer
	*ent.Client
}

func (s *TasksServer) GetTask(ctx context.Context, in *proto.GetTaskRequest) (*proto.Task, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	task, err := s.Client.Task.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	t := &proto.Task{}
	entTaskToProto(task, t)

	return t, nil
}

func (s *TasksServer) CreateTask(ctx context.Context, in *proto.CreateTaskRequest) (*proto.Task, error) {
	tx, err := s.Client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	user, err := ensureUser(ctx, tx.User, in.OwnerId)
	if err != nil {
		return nil, err
	}

	task, err := tx.Task.
		Create().
		SetTitle(in.Title).
		SetOwner(user).
		Save(ctx)

	if err != nil {
		return nil, err
	}

	err = tx.Commit()

	t := &proto.Task{}
	entTaskToProto(task, t)

	return t, err
}

func (s *TasksServer) UpdateTask(ctx context.Context, in *proto.UpdateTaskRequest) (*proto.Task, error) {
	if is, _ := util.IsNillStruct(in, "Id"); is { // this error should be impossible unless protobuffs are broken
		return nil, fmt.Errorf("no fields to update")
	}

	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	b := s.Client.Task.Update()

	if in.GetTitle() != "" || in.Title != nil {
		b = b.SetTitle(in.GetTitle())
	}

	_, err = b.Save(ctx)
	if err != nil {
		return nil, err
	}

	task, err := s.Client.Task.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	t := &proto.Task{}
	entTaskToProto(task, t)

	return t, nil
}

func (s *TasksServer) DeleteTask(ctx context.Context, in *proto.GetTaskRequest) (*proto.Empty, error) {
	id, err := strconv.Atoi(in.Id)
	if err != nil {
		return nil, err
	}

	return &proto.Empty{}, s.Client.Task.DeleteOneID(id).Exec(ctx)
}

func entTaskToProto(s *ent.Task, t *proto.Task) {
	t.Title = s.Title
	t.Id = fmt.Sprintf("%d", s.ID)
	t.CreatedAt = s.CreatedAt.Format(time.RFC3339)
	t.UpdatedAt = s.UpdatedAt.Format(time.RFC3339)
}

func ensureUser(ctx context.Context, client *ent.UserClient, userID string) (*ent.User, error) {
	user, err := client.Get(ctx, userID)
	if err != nil {
		return nil, err
	}

	if user == nil {
		user, err = client.Create().SetID(userID).Save(ctx)
	}

	return user, err
}
