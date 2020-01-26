// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/kcarretto/paragon/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "Name" field.
	Name string `json:"Name,omitempty"`
	// Email holds the value of the "Email" field.
	Email string `json:"Email,omitempty"`
	// OAuthID holds the value of the "OAuthID" field.
	OAuthID string `json:"-"`
	// PhotoURL holds the value of the "PhotoURL" field.
	PhotoURL string `json:"PhotoURL,omitempty"`
	// SessionToken holds the value of the "SessionToken" field.
	SessionToken string `json:"-"`
	// Activated holds the value of the "Activated" field.
	Activated bool `json:"Activated,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues() []interface{} {
	return []interface{}{
		&sql.NullInt64{},  // id
		&sql.NullString{}, // Name
		&sql.NullString{}, // Email
		&sql.NullString{}, // OAuthID
		&sql.NullString{}, // PhotoURL
		&sql.NullString{}, // SessionToken
		&sql.NullBool{},   // Activated
	}
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(values ...interface{}) error {
	if m, n := len(values), len(user.Columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	value, ok := values[0].(*sql.NullInt64)
	if !ok {
		return fmt.Errorf("unexpected type %T for field id", value)
	}
	u.ID = int(value.Int64)
	values = values[1:]
	if value, ok := values[0].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Name", values[0])
	} else if value.Valid {
		u.Name = value.String
	}
	if value, ok := values[1].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field Email", values[1])
	} else if value.Valid {
		u.Email = value.String
	}
	if value, ok := values[2].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field OAuthID", values[2])
	} else if value.Valid {
		u.OAuthID = value.String
	}
	if value, ok := values[3].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field PhotoURL", values[3])
	} else if value.Valid {
		u.PhotoURL = value.String
	}
	if value, ok := values[4].(*sql.NullString); !ok {
		return fmt.Errorf("unexpected type %T for field SessionToken", values[4])
	} else if value.Valid {
		u.SessionToken = value.String
	}
	if value, ok := values[5].(*sql.NullBool); !ok {
		return fmt.Errorf("unexpected type %T for field Activated", values[5])
	} else if value.Valid {
		u.Activated = value.Bool
	}
	return nil
}

// Update returns a builder for updating this User.
// Note that, you need to call User.Unwrap() before calling this method, if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{u.config}).UpdateOne(u)
}

// Unwrap unwraps the entity that was returned from a transaction after it was closed,
// so that all next queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", Name=")
	builder.WriteString(u.Name)
	builder.WriteString(", Email=")
	builder.WriteString(u.Email)
	builder.WriteString(", OAuthID=<sensitive>")
	builder.WriteString(", PhotoURL=")
	builder.WriteString(u.PhotoURL)
	builder.WriteString(", SessionToken=<sensitive>")
	builder.WriteString(", Activated=")
	builder.WriteString(fmt.Sprintf("%v", u.Activated))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}