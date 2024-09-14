package repository

import (
	"context"
	"time"

	"github.com/vector-ops/lifefolio/ent"
	"github.com/vector-ops/lifefolio/ent/user"
	"github.com/vector-ops/lifefolio/internal/dto/request"
	"github.com/vector-ops/lifefolio/internal/models"
	"github.com/vector-ops/lifefolio/utils"
)

type UserRepository struct {
	db *ent.Client
}

func NewUserRepository(db *ent.Client) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(ctx context.Context, userIn request.UserRegisterIn) (*models.User, error) {

	patientId := utils.GenPatientId(userIn.FirstName, userIn.LastName)
	user, err := r.db.User.Create().
		SetPatientID(patientId).
		SetFirstName(userIn.FirstName).
		SetLastName(userIn.LastName).
		SetEmail(userIn.Email).
		SetPassword(userIn.Password).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	return models.TransformUser(user), nil
}

func (r *UserRepository) GetUserByPatientId(ctx context.Context, patientId string) (*models.User, error) {
	user, err := r.db.User.Query().
		Where(user.PatientIDEQ(patientId)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return models.TransformUser(user), nil
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	user, err := r.db.User.Query().
		Where(user.EmailEQ(email)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	return models.TransformUser(user), nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, patientId string, userIn request.UserUpdateIn) (*models.User, error) {
	err := r.db.User.Update().
		Where(user.PatientIDEQ(patientId)).
		SetNillableHeight(&userIn.Height).
		SetNillableWeight(&userIn.Weight).
		SetNillableDOB(&userIn.DOB).
		SetNillableBloodGroup(&userIn.BloodGroup).
		SetNillableFirstName(&userIn.FirstName).
		SetNillableLastName(&userIn.LastName).
		SetUpdatedAt(time.Now()).
		Exec(ctx)
	if err != nil {
		return nil, err
	}
	context, cancel := context.WithTimeout(context.Background(), time.Second*3)
	defer cancel()
	return r.GetUserByPatientId(context, patientId)
}

func (r *UserRepository) UpdateEmail(ctx context.Context, patientId string, email string) error {
	return r.db.User.Update().
		Where(user.PatientIDEQ(patientId)).
		SetEmail(email).
		SetIsVerified(false).
		Exec(ctx)
}

func (r *UserRepository) UpdateVerified(ctx context.Context, patientId string) error {
	return r.db.User.Update().
		Where(user.PatientIDEQ(patientId)).
		SetIsVerified(true).
		SetVerifiedAt(time.Now()).
		Exec(ctx)
}

func (r *UserRepository) ArchiveUser(ctx context.Context, patientId string) error {
	return r.db.User.Update().
		SetIsArchived(true).
		SetArchivedAt(time.Now()).
		Exec(ctx)
}

func (r *UserRepository) AddOTP(ctx context.Context, otp uint64) error {
	return r.db.User.Update().
		SetOtp(otp).
		Exec(ctx)
}

func (r *UserRepository) UpdateUserType(ctx context.Context, userType user.UserType) error {
	return r.db.User.Update().
		SetUserType(userType).
		Exec(ctx)
}
