package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent"
)

type MedicalRecord struct {
	ID            uuid.UUID  `json:"id,omitempty"`
	UserId        uuid.UUID  `json:"user_id,omitempty"`
	InstitutionId *uuid.UUID `json:"institution_id,omitempty"`
	File          string     `json:"file,omitempty"`
	CreatedAt     time.Time  `json:"created_at,omitempty"`
	UpdatedAt     time.Time  `json:"updated_at,omitempty"`
}

func TransformMedicalRecord(mr *ent.MedicalRecord) *MedicalRecord {
	return &MedicalRecord{
		ID:            mr.ID,
		UserId:        mr.UserID,
		InstitutionId: &mr.InstitutionID,
		File:          mr.File,
		CreatedAt:     mr.CreatedAt,
		UpdatedAt:     mr.UpdatedAt,
	}
}
