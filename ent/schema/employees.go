package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
)

// Employees holds the schema definition for the Employees entity.
type Employees struct {
	ent.Schema
}

// Fields of the Employees.
func (Employees) Fields() []ent.Field {
	return []ent.Field{
		field.String("gender").NotEmpty(),
	}
}
