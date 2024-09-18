package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent"
	"github.com/vector-ops/lifefolio/ent/recordaccess"
	"github.com/vector-ops/lifefolio/internal/models"
)

type RecordAccessRepository struct {
	db *ent.Client
}

func NewRecordAccessRepository(db *ent.Client) *RecordAccessRepository {
	return &RecordAccessRepository{
		db: db,
	}
}

func (r *RecordAccessRepository) CreateRecordAccess(ctx context.Context, instId uuid.UUID, recordId uuid.UUID) (*models.RecordAccess, error) {
	ra, err := r.db.RecordAccess.Create().
		SetInstitutionID(instId).
		SetRecordID(recordId).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return models.TransformRecordAccess(ra), nil
}

func (r *RecordAccessRepository) GetRecordAccess(ctx context.Context, recordId uuid.UUID) ([]*models.RecordAccess, error) {
	ra, err := r.db.RecordAccess.Query().
		Where(recordaccess.RecordIDEQ(recordId)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	recordAccess := make([]*models.RecordAccess, len(ra))
	for _, v := range ra {
		recordAccess = append(recordAccess, models.TransformRecordAccess(v))
	}

	return recordAccess, nil
}

func (r *RecordAccessRepository) UpdateRecordAccess(ctx context.Context, instId uuid.UUID, recordId uuid.UUID, approved bool) error {
	return r.db.RecordAccess.Update().
		Where(recordaccess.RecordIDEQ(recordId), recordaccess.InstitutionIDEQ(instId)).
		SetApproved(approved).
		SetApprovedAt(time.Now()).
		Exec(ctx)
}

func (r *RecordAccessRepository) DeleteRecordAccess(ctx context.Context, instId uuid.UUID, recordId uuid.UUID) error {
	_, err := r.db.RecordAccess.Delete().
		Where(recordaccess.RecordIDEQ(recordId), recordaccess.InstitutionIDEQ(instId)).
		Exec(ctx)
	return err
}
