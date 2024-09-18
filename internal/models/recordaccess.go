package models

import (
	"time"

	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent"
)

type RecordAccess struct {
	ID            uuid.UUID `json:"id,omitempty"`
	RecordId      uuid.UUID `json:"record_id,omitempty"`
	InstitutionId uuid.UUID `json:"institution_id,omitempty"`
	Approved      bool      `json:"approved,omitempty"`
	ApprovedAt    time.Time `json:"approved_at,omitempty"`
}

func TransformRecordAccess(ra *ent.RecordAccess) *RecordAccess {
	return &RecordAccess{
		ID:            ra.ID,
		RecordId:      ra.RecordID,
		InstitutionId: ra.InstitutionID,
		Approved:      ra.Approved,
		ApprovedAt:    *ra.ApprovedAt,
	}
}
