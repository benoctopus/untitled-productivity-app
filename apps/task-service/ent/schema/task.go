package schema

import (
	"fmt"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Task holds the schema definition for the Task entity.

type TaskStatus int8

func (ts *TaskStatus) GetName() (string, error) {
	switch *ts {
	case 0:
		return "ACTIVE", nil
	case 1:
		return "COMPLETE", nil
	case 2:
		return "WONT_DO", nil
	case 3:
		return "DELETED", nil
	default:
		return "", fmt.Errorf("invalid status")
	}
}

const TASK_STATUS_ACTIVE = 0
const TASK_STATUS_COMPLETE = 1
const TASK_STATUS_WONT_DO = 2
const TASK_STATUS_DELETED = 3

type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").NotEmpty(),
		field.Uint8("status").Default(0).Max(3).Min(0),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("owner", User.Type).Ref("tasks").Unique(),
	}
}
