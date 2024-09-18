package request

import "github.com/google/uuid"

type MedicalRecordIn struct {
	UserId        uuid.UUID
	InstitutionId uuid.UUID
	File          string
}
