package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent"
	"github.com/vector-ops/lifefolio/ent/user"
)

type User struct {
	ID         uuid.UUID       `json:"id,omitempty"`
	PatientId  string          `json:"patient_id,omitempty"`
	FirstName  string          `json:"first_name,omitempty"`
	LastName   string          `json:"last_name,omitempty"`
	Email      string          `json:"email,omitempty"`
	DOB        time.Time       `json:"dob,omitempty"`
	UserType   user.UserType   `json:"user_type,omitempty"`
	BloodGroup user.BloodGroup `json:"blood_group,omitempty"`
	Weight     float32         `json:"weight,omitempty"`
	Height     float32         `json:"height,omitempty"`
	IsArchived bool            `json:"is_archived,omitempty"`
	IsVerified bool            `json:"is_verified,omitempty"`
	CreatedAt  time.Time       `json:"created_at,omitempty"`
	UpdatedAt  time.Time       `json:"updated_at,omitempty"`
	ArchivedAt *time.Time      `json:"archived_at,omitempty"`
	VerifiedAt *time.Time      `json:"verified_at,omitempty"`
}

func TransformUser(user *ent.User) *User {
	return &User{
		ID:         user.ID,
		PatientId:  user.PatientID,
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		DOB:        user.DOB,
		UserType:   user.UserType,
		BloodGroup: user.BloodGroup,
		Weight:     user.Weight,
		Height:     user.Height,
		IsArchived: user.IsArchived,
		IsVerified: user.IsVerified,
		ArchivedAt: user.ArchivedAt,
		VerifiedAt: user.VerifiedAt,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
	}
}
