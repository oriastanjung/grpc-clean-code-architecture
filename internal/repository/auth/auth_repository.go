package repository

import (
	"github.com/oriastanjung/grpc_server/internal/entities"
	"gorm.io/gorm"
)

type AuthRepository interface{
	RegisterAdmin(user *entities.User) error
}

type authRepository struct{
	db *gorm.DB
}


func NewAuthRepository(db *gorm.DB) AuthRepository{
	return &authRepository{
		db:db,
	}
}

func (repo *authRepository) RegisterAdmin(user *entities.User) error {
	return repo.db.Create(user).Error
}