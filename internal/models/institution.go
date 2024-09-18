package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent"
)

type Institution struct {
	ID          uuid.UUID  `json:"id,omitempty"`
	Name        string     `json:"name,omitempty"`
	Description *string    `json:"description,omitempty"`
	Location    string     `json:"location,omitempty"`
	Email       string     `json:"email,omitempty"`
	Password    string     `json:"-"`
	Phone       string     `json:"phone,omitempty"`
	IsArchived  bool       `json:"is_archived,omitempty"`
	IsVerified  bool       `json:"is_verified,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
	ArchivedAt  *time.Time `json:"archived_at,omitempty"`
	VerifiedAt  *time.Time `json:"verified_at,omitempty"`
}

func TransformInstitution(inst *ent.Institution) *Institution {
	return &Institution{
		ID:          inst.ID,
		Name:        inst.Name,
		Description: inst.Description,
		Location:    inst.Location,
		Email:       inst.Email,
		Password:    inst.Password,
		Phone:       inst.Phone,
		IsArchived:  inst.IsArchived,
		IsVerified:  inst.IsVerified,
		CreatedAt:   inst.CreatedAt,
		UpdatedAt:   inst.UpdatedAt,
		VerifiedAt:  inst.VerifiedAt,
		ArchivedAt:  inst.ArchivedAt,
	}
}
