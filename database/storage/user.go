package storage

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"project/models"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) User {
	return &UserRepo{db: db}
}

func (r *UserRepo) Create(user *models.User) (string, error) {
	id := uuid.New()
	user.Id = id
	if err := r.db.Create(user).Error; err != nil {
		return "", err
	}

	return id.String(), nil
}

func (r *UserRepo) Update(user *models.User) error {
	if err := r.db.Save(user).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Delete(req models.RequestId) error {
	if err := r.db.Where("id = ?", req.Id).Delete(&models.User{}).Error; err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) Get(req models.RequestId) (*models.User, error) {
	var user models.User
	if err := r.db.Where("id = ?", req.Id).First(&user).Error; err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepo) GetAll(req models.GetAllUsersRequest) (*models.GetAllUsersResponse, error) {
	var (
		resp   models.GetAllUsersResponse
		query  = r.db.Model(&models.User{})
		offset = (req.Page - 1) * req.Limit
	)

	if req.Search != "" {
		query = query.Where("username ILIKE ?", "%"+req.Search+"%")
	}

	err := query.Offset(int(offset)).Limit(int(req.Limit)).Find(&resp.Users).Error
	if err != nil {
		return nil, err
	}

	err = query.Count(&resp.Count).Error
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

func (r *UserRepo) GetByUsername(username string) (*models.User, error) {
	var user models.User
	err := r.db.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &user, err
	}

	return &user, err
}
