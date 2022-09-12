// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"golang-api/ent/employees"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Employees is the model entity for the Employees schema.
type Employees struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Gender holds the value of the "gender" field.
	Gender string `json:"gender,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Employees) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case employees.FieldID:
			values[i] = new(sql.NullInt64)
		case employees.FieldGender:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Employees", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Employees fields.
func (e *Employees) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case employees.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			e.ID = int(value.Int64)
		case employees.FieldGender:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field gender", values[i])
			} else if value.Valid {
				e.Gender = value.String
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Employees.
// Note that you need to call Employees.Unwrap() before calling this method if this Employees
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Employees) Update() *EmployeesUpdateOne {
	return (&EmployeesClient{config: e.config}).UpdateOne(e)
}

// Unwrap unwraps the Employees entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Employees) Unwrap() *Employees {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("ent: Employees is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Employees) String() string {
	var builder strings.Builder
	builder.WriteString("Employees(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("gender=")
	builder.WriteString(e.Gender)
	builder.WriteByte(')')
	return builder.String()
}

// EmployeesSlice is a parsable slice of Employees.
type EmployeesSlice []*Employees

func (e EmployeesSlice) config(cfg config) {
	for _i := range e {
		e[_i].config = cfg
	}
}
