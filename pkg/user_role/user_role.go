package user_role

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type UserRoleStatus string

const (
	Active   UserRoleStatus = "Active"
	Inactive UserRoleStatus = "Inactive"
)

// UserRole is used by pop to map your user_roles database table to your go code.
type UserRole struct {
	ID        uuid.UUID      `json:"id" db:"id"`
	UserId    uuid.UUID      `json:"user_id" db:"user_id"`
	RoleId    uuid.UUID      `json:"role_id" db:"role_id"`
	Status    UserRoleStatus `json:"status" db:"status"`
	CreatedAt time.Time      `json:"created_at" db:"created_at"`
	UpdatedAt time.Time      `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (u UserRole) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UserRoles is not required by pop and may be deleted
type UserRoles []UserRole

// String is not required by pop and may be deleted
func (u UserRoles) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UserRole) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UserRole) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UserRole) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
