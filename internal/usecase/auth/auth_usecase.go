package usecase

import (
	"errors"
	"fmt"

	"github.com/oriastanjung/grpc_server/internal/entities"
	repository "github.com/oriastanjung/grpc_server/internal/repository/auth"
	"github.com/oriastanjung/grpc_server/internal/utils"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthUseCase interface{
	RegisterAdmin(user *entities.User, passwordSalt int) error
	LoginAdmin(user *entities.User) (string,error)
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
		return status.Errorf(codes.Internal, fmt.Sprintf("Error hashing password: %v", err))
	}
	user.Password = string(hashedPassword)

	// 4. Simpan data pengguna ke database melalui repository
	return usecase.authRepo.RegisterAdmin(user)
}

func (usecase *authUseCase) LoginAdmin(user *entities.User)(string, error){
	dbUser:= &entities.User{}
	dbUser.Email = user.Email
	err:= usecase.authRepo.LoginAdmin(dbUser)
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return "",status.Errorf(codes.NotFound, "User Not Found")
		}
		return "",err
	}

	if dbUser.IsVerified == false{
		return "",status.Errorf(codes.Unauthenticated, "User Not Verified")
	}

	if dbUser.Role != string(entities.AdminRole){
		return "",errors.New("User Not Admin")
	}

	err = bcrypt.CompareHashAndPassword([]byte(dbUser.Password), []byte(user.Password))
	if err != nil{
		return "",status.Errorf(codes.Unauthenticated, "Invalid Password")
	}

	token,err := utils.GenerateTokenJWT(*user)
	if err != nil{
		return "", status.Errorf(codes.Internal, fmt.Sprintf("Error : %v",err))
	}
	return token,nil
}

