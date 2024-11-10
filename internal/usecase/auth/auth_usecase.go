package usecase

import (
	"errors"

	"github.com/oriastanjung/grpc_server/internal/entities"
	repository "github.com/oriastanjung/grpc_server/internal/repository/auth"
	"github.com/oriastanjung/grpc_server/internal/utils"
	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface{
	RegisterAdmin(user *entities.User, passwordSalt int) error
}


type authUseCase struct{
	authRepo repository.AuthRepository
}

func NewAuthUseCase(authRepo repository.AuthRepository) AuthUseCase{
	return &authUseCase{
		authRepo: authRepo,
	}
}


func (usecase *authUseCase) RegisterAdmin(user *entities.User, passwordSalt int) error{
	// 1. Generate ID unik untuk pengguna
	user.ID = utils.GenerateIDbyKSUID()

	// 2. Set peran pengguna sebagai admin
	user.Role = string(entities.AdminRole)
	user.IsVerified = true

	// 3. Hash password untuk keamanan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), passwordSalt)
	if err != nil {
		return errors.New("Failed hashing password")
	}
	user.Password = string(hashedPassword)

	// 4. Simpan data pengguna ke database melalui repository
	return usecase.authRepo.RegisterAdmin(user)
}