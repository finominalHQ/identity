package role_permission

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type RolePermissionStatus string

const (
	Active   RolePermissionStatus = "Active"
	Inactive RolePermissionStatus = "Inactive"
)

// RolePermission is used by pop to map your role_permissions database table to your go code.
type RolePermission struct {
	ID           uuid.UUID            `json:"id" db:"id"`
	RoleId       uuid.UUID            `json:"role_id" db:"role_id"`
	PermissionId uuid.UUID            `json:"permission_id" db:"permission_id"`
	Status       RolePermissionStatus `json:"status" db:"status"`
	CreatedAt    time.Time            `json:"created_at" db:"created_at"`
	UpdatedAt    time.Time            `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (u RolePermission) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// RolePermissions is not required by pop and may be deleted
type RolePermissions []RolePermission

// String is not required by pop and may be deleted
func (u RolePermissions) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *RolePermission) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *RolePermission) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *RolePermission) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
