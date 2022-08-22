package user

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type IdType string

const (
	Email IdType = "Email"
	Phone IdType = "Phone"
)

type UserType string

const (
	Person  UserType = "Person"
	Machine UserType = "Machine"
)

type UserStatus string

const (
	Active   UserStatus = "Active"
	Inactive UserStatus = "Inactive"
)

// User is used by pop to map your users database table to your go code.
type User struct {
	ID          uuid.UUID    `json:"id" db:"id"`
	IdNumber    string       `json:"id_number" db:"id_number"`
	IdType      IdType       `json:"id_type" db:"id_type"`
	Password    string       `json:"password" db:"password"`
	Type        UserType     `json:"type" db:"type"`
	Source      nulls.String `json:"source" db:"source"`
	Status      UserStatus   `json:"status" db:"status"`
	SuspendedAt time.Time    `json:"suspended_at" db:"suspended_at"`
	ClosedAt    time.Time    `json:"closed_at" db:"closed_at"`
	CreatedAt   time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
