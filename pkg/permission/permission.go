package permission

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/nulls"
	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type PermissionStatus string

const (
	Active   PermissionStatus = "Active"
	Inactive PermissionStatus = "Inactive"
)

// Permission is used by pop to map your permissions database table to your go code.
type Permission struct {
	ID        uuid.UUID        `json:"id" db:"id"`
	Name      string           `json:"name" db:"name"`
	Desc      nulls.String     `json:"desc" db:"desc"`
	Status    PermissionStatus `json:"status" db:"status"`
	CreatedAt time.Time        `json:"created_at" db:"created_at"`
	UpdatedAt time.Time        `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time        `json:"deleted_at" db:"deleted_at"`
}

// String is not required by pop and may be deleted
func (u Permission) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Permissions is not required by pop and may be deleted
type Permissions []Permission

// String is not required by pop and may be deleted
func (u Permissions) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *Permission) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *Permission) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *Permission) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
