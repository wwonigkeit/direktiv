// Code generated by entc, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/vorteil/direktiv/ent/subroutine"
)

// Subroutine is the model entity for the Subroutine schema.
type Subroutine struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CallerID holds the value of the "callerID" field.
	CallerID string `json:"callerID,omitempty"`
	// Semaphore holds the value of the "semaphore" field.
	Semaphore int `json:"semaphore,omitempty"`
	// Memory holds the value of the "memory" field.
	Memory string `json:"memory,omitempty"`
	// SubroutineIDs holds the value of the "subroutineIDs" field.
	SubroutineIDs []string `json:"subroutineIDs,omitempty"`
	// SubroutineResponses holds the value of the "subroutineResponses" field.
	SubroutineResponses []string `json:"subroutineResponses,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Subroutine) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case subroutine.FieldSubroutineIDs, subroutine.FieldSubroutineResponses:
			values[i] = &[]byte{}
		case subroutine.FieldID, subroutine.FieldSemaphore:
			values[i] = &sql.NullInt64{}
		case subroutine.FieldCallerID, subroutine.FieldMemory:
			values[i] = &sql.NullString{}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Subroutine", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Subroutine fields.
func (s *Subroutine) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case subroutine.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case subroutine.FieldCallerID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field callerID", values[i])
			} else if value.Valid {
				s.CallerID = value.String
			}
		case subroutine.FieldSemaphore:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field semaphore", values[i])
			} else if value.Valid {
				s.Semaphore = int(value.Int64)
			}
		case subroutine.FieldMemory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field memory", values[i])
			} else if value.Valid {
				s.Memory = value.String
			}
		case subroutine.FieldSubroutineIDs:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field subroutineIDs", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.SubroutineIDs); err != nil {
					return fmt.Errorf("unmarshal field subroutineIDs: %v", err)
				}
			}
		case subroutine.FieldSubroutineResponses:

			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field subroutineResponses", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &s.SubroutineResponses); err != nil {
					return fmt.Errorf("unmarshal field subroutineResponses: %v", err)
				}
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Subroutine.
// Note that you need to call Subroutine.Unwrap() before calling this method if this Subroutine
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Subroutine) Update() *SubroutineUpdateOne {
	return (&SubroutineClient{config: s.config}).UpdateOne(s)
}

// Unwrap unwraps the Subroutine entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Subroutine) Unwrap() *Subroutine {
	tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Subroutine is not a transactional entity")
	}
	s.config.driver = tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Subroutine) String() string {
	var builder strings.Builder
	builder.WriteString("Subroutine(")
	builder.WriteString(fmt.Sprintf("id=%v", s.ID))
	builder.WriteString(", callerID=")
	builder.WriteString(s.CallerID)
	builder.WriteString(", semaphore=")
	builder.WriteString(fmt.Sprintf("%v", s.Semaphore))
	builder.WriteString(", memory=")
	builder.WriteString(s.Memory)
	builder.WriteString(", subroutineIDs=")
	builder.WriteString(fmt.Sprintf("%v", s.SubroutineIDs))
	builder.WriteString(", subroutineResponses=")
	builder.WriteString(fmt.Sprintf("%v", s.SubroutineResponses))
	builder.WriteByte(')')
	return builder.String()
}

// Subroutines is a parsable slice of Subroutine.
type Subroutines []*Subroutine

func (s Subroutines) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}