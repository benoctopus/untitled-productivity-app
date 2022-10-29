package schema

import (
	"fmt"
	"regexp"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Tag holds the schema definition for the Tag entity.
type Tag struct {
	ent.Schema
}

var tagNamePattern = "^[A-Za-z0-9][A-Za-z0-9_-]*[A-Za-z0-9]$"
var tagNameExpression = regexp.MustCompile(tagNamePattern)

// ValidateTagSlug validates a string against `tagNamePattern`.
func ValidateTagSlug(slug string) error {
	if tagNameExpression.MatchString(slug) {
		return nil
	}

	return fmt.Errorf("invalid tag name: %s - must match %s", slug, tagNamePattern)
}

// Fields of the Tag.
func (Tag) Fields() []ent.Field {
	return []ent.Field{
		field.String("slug").
			NotEmpty().
			Validate(ValidateTagSlug),
		field.String("description").NotEmpty().Optional(),
		field.String("color").MaxLen(6).MinLen(6).Default("000000"),
		field.Time("created_at").Default(time.Now),
		field.Time("updated_at").Default(time.Now).UpdateDefault(time.Now),
	}
}

// Edges of the Tag.
func (Tag) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tasks", Task.Type),
		edge.To("owner", User.Type).Unique(),
	}
}
