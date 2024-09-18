package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent"
	"github.com/vector-ops/lifefolio/ent/medicalrecord"
	"github.com/vector-ops/lifefolio/ent/recordaccess"
	"github.com/vector-ops/lifefolio/internal/models"
)

type MedicalRecordRepository struct {
	db *ent.Client
}

func NewMedicalRecordRepository(db *ent.Client) *MedicalRecordRepository {
	return &MedicalRecordRepository{
		db: db,
	}
}

func (r *MedicalRecordRepository) CreateMedicalRecord(ctx context.Context, uid uuid.UUID, instId uuid.UUID, file string) (*models.MedicalRecord, error) {
	mr, err := r.db.MedicalRecord.Create().SetUserID(uid).SetFile(file).SetNillableInstitutionID(&instId).Save(ctx)
	if err != nil {
		return nil, err
	}

	return models.TransformMedicalRecord(mr), nil
}

func (r *MedicalRecordRepository) UpdateMedicalRecord(ctx context.Context, uid uuid.UUID, instId uuid.UUID, file string) error {
	return r.db.MedicalRecord.Update().Where(medicalrecord.UserIDEQ(uid)).SetNillableFile(&file).SetNillableInstitutionID(&instId).SetUpdatedAt(time.Now()).Exec(ctx)
}

func (r *MedicalRecordRepository) GetMedicalRecordsByUserId(ctx context.Context, uid uuid.UUID) ([]*models.MedicalRecord, error) {
	mrs, err := r.db.MedicalRecord.Query().Where(medicalrecord.UserIDEQ(uid)).All(ctx)
	if err != nil {
		return nil, err
	}

	transformedMR := make([]*models.MedicalRecord, len(mrs))
	for _, v := range mrs {
		transformedMR = append(transformedMR, models.TransformMedicalRecord(v))
	}

	return transformedMR, nil
}

func (r *MedicalRecordRepository) GetMedicalRecordsByInstId(ctx context.Context, instId uuid.UUID) ([]*models.MedicalRecord, error) {
	mrs, err := r.db.MedicalRecord.Query().Where(medicalrecord.And(medicalrecord.InstitutionIDEQ(instId), medicalrecord.HasRecordaccessWith(recordaccess.Approved(true)))).All(ctx)
	if err != nil {
		return nil, err
	}

	transformedMR := make([]*models.MedicalRecord, len(mrs))
	for _, v := range mrs {
		transformedMR = append(transformedMR, models.TransformMedicalRecord(v))
	}

	return transformedMR, nil
}
