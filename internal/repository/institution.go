package repository

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/vector-ops/lifefolio/ent"
	"github.com/vector-ops/lifefolio/ent/institution"
	"github.com/vector-ops/lifefolio/internal/dto/request"
	"github.com/vector-ops/lifefolio/internal/models"
)

type InstRepository struct {
	db *ent.Client
}

func NewInstRepository(db *ent.Client) *InstRepository {
	return &InstRepository{
		db: db,
	}
}

func (r *InstRepository) CreateInstitution(ctx context.Context, instIn request.InstRegisterIn) (*models.Institution, error) {
	inst, err := r.db.Institution.Create().
		SetName(instIn.Name).
		SetEmail(instIn.Email).
		SetDescription(instIn.Description).
		SetLocation(instIn.Location).
		SetPassword(instIn.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return models.TransformInstitution(inst), nil
}

func (r *InstRepository) GetInstitutionByEmail(ctx context.Context, email string) (*models.Institution, error) {
	inst, err := r.db.Institution.Query().
		Where(institution.
			EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return models.TransformInstitution(inst), nil
}

func (r *InstRepository) GetInstitutionById(ctx context.Context, id uuid.UUID) (*models.Institution, error) {
	inst, err := r.db.Institution.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	return models.TransformInstitution(inst), nil
}

func (r *InstRepository) GetInstitutionsByName(ctx context.Context, name string) ([]*models.Institution, error) {
	inst, err := r.db.Institution.Query().
		Where(institution.NameContains(name)).
		All(ctx)
	if err != nil {
		return nil, err
	}

	institusions := make([]*models.Institution, len(inst))

	for _, i := range inst {
		institusions = append(institusions, models.TransformInstitution(i))
	}

	return institusions, nil
}

func (r *InstRepository) UpdateInstitution(ctx context.Context, instIn request.InstUpdateIn) (*models.Institution, error) {
	err := r.db.Institution.UpdateOneID(instIn.ID).
		SetNillableName(&instIn.Name).
		SetNillableDescription(&instIn.Description).
		SetNillableLocation(&instIn.Location).
		SetNillablePhone(&instIn.Phone).
		SetNillablePassword(&instIn.Password).
		SetUpdatedAt(time.Now()).
		Exec(ctx)

	if err != nil {
		return nil, err
	}

	context, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return r.GetInstitutionById(context, instIn.ID)
}

func (r *InstRepository) UpdateEmail(ctx context.Context, email string, id uuid.UUID) error {
	return r.db.Institution.UpdateOneID(id).
		SetEmail(email).
		SetIsVerified(false).
		Exec(ctx)
}

func (r *InstRepository) UpdateVerified(ctx context.Context, email string) error {
	return r.db.Institution.Update().
		Where(institution.EmailEQ(email)).
		SetVerifiedAt(time.Now()).
		SetIsVerified(true).
		Exec(ctx)
}

func (r *InstRepository) ArchiveInstitution(ctx context.Context, email string) error {
	return r.db.Institution.Update().
		Where(institution.EmailEQ(email)).
		SetArchivedAt(time.Now()).
		SetIsArchived(true).
		Exec(ctx)
}

func (r *InstRepository) GetOTP(ctx context.Context, email string) (*uint64, error) {
	inst, err := r.db.Institution.Query().
		Where(institution.EmailEQ(email)).
		Select(institution.FieldOtp).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return inst.Otp, nil
}

func (r *InstRepository) AddOTP(ctx context.Context, email string, otp uint64) error {
	return r.db.Institution.Update().
		Where(institution.EmailEQ(email)).
		SetOtp(otp).
		Exec(ctx)
}

func (r *InstRepository) AddDoctor(ctx context.Context, userIds uuid.UUIDs, instId uuid.UUID) error {
	return r.db.Institution.Update().Where(institution.IDEQ(instId)).AddDoctorIDs(userIds...).Exec(ctx)
}

func (r *InstRepository) RemoveDoctor(ctx context.Context, userIds uuid.UUIDs, instId uuid.UUID) error {
	return r.db.Institution.Update().Where(institution.IDEQ(instId)).RemoveDoctorIDs(userIds...).Exec(ctx)
}
