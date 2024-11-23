package repository

import (
	"fmt"

	"github.com/oriastanjung/grpc_server/internal/entities"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthRepository interface{
	RegisterAdmin(user *entities.User) error
	LoginAdmin(user *entities.User) error
	RegisterUser(user *entities.User) error
	LoginUser(user *entities.User) error
	VerifyUser(token string) error
	FindUserByEmail(email string) (*entities.User ,error)
	UpdateUserByEmail(Email string, dataUpdated *entities.User) (error)
	FindOneUserByKey(key string, val string) (*entities.User ,error)
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

func (repo *authRepository) LoginAdmin(user *entities.User) error {
	return repo.db.Where("email = ?", user.Email).First(user).Error
}

func (repo *authRepository) RegisterUser(user *entities.User) error {
	return repo.db.Create(user).Error
}

func (repo *authRepository) LoginUser(user *entities.User) error {
	return repo.db.Where("email = ?", user.Email).First(user).Error
}

func (repo *authRepository) VerifyUser(token string) error {
	var user entities.User
	err := repo.db.Where("verification_token = ?", token).First(&user).Error
	if err != nil{
		return status.Errorf(codes.NotFound, "User Not Found")
	}
	user.IsVerified = true
	user.VerificationToken = ""
	err = repo.db.Save(&user).Error; 
	
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error saving user: %v", err))
	}
	
	return nil
}

func (repo *authRepository) FindUserByEmail(email string) (*entities.User ,error) {
	var user entities.User
	err := repo.db.Where("email = ?", email).First(&user).Error
	if err != nil{
		return nil, status.Errorf(codes.NotFound, "User Not Found")
	}
	return &user, nil
	
}
func (repo *authRepository) UpdateUserByEmail(Email string, dataUpdated *entities.User) (error) {
	var user entities.User
	err := repo.db.Where("email = ?", Email).First(&user).Error
	if err != nil{
		return status.Errorf(codes.NotFound, "User Not Found")
	}

	user.Username = dataUpdated.Username
	user.Password = dataUpdated.Password
	user.ForgetPasswordToken = dataUpdated.ForgetPasswordToken
	user.ProfilePicture = dataUpdated.ProfilePicture
	user.ProfilePictureUrl = dataUpdated.ProfilePictureUrl
	user.Username = dataUpdated.Username

	err = repo.db.Save(&user).Error; 

	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error saving user: %v", err))
	}

	return nil
}

func(repo *authRepository) FindOneUserByKey(key string, val string) (*entities.User ,error) {
	var user entities.User
	err := repo.db.Where(fmt.Sprintf("%s = ?", key), val).First(&user).Error
	if err != nil{
		return nil, status.Errorf(codes.NotFound, "User Not Found")
	}
	return &user, nil

}