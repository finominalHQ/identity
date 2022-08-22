package otp

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v6"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

type OtpType string

const (
	Email OtpType = "Email"
	Phone OtpType = "Phone"
)

// Otp is used by pop to map your otps database table to your go code.
type Otp struct {
	ID        uuid.UUID `json:"id" db:"id"`
	Value     string    `json:"value" db:"value"`
	Type      OtpType   `json:"type" db:"type"`
	Owner     string    `json:"owner" db:"owner"`
	UsedAt    time.Time `json:"used_at" db:"used_at"`
	ExpiredAt time.Time `json:"expired_at" db:"expired_at"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt time.Time `json:"deleted_at" db:"deleted_at"`
}

// String is not required by pop and may be deleted
func (u Otp) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Otps is not required by pop and may be deleted
type Otps []Otp

// String is not required by pop and may be deleted
func (u Otps) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *Otp) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *Otp) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *Otp) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
