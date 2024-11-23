package usecase

import (
	"errors"
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/oriastanjung/grpc_server/internal/config"
	"github.com/oriastanjung/grpc_server/internal/entities"
	repository "github.com/oriastanjung/grpc_server/internal/repository/auth"
	"github.com/oriastanjung/grpc_server/internal/utils"
	"github.com/oriastanjung/grpc_server/internal/utils/smtp"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type AuthUseCase interface{
	RegisterAdmin(user *entities.User, passwordSalt int) error
	LoginAdmin(user *entities.User) (string,error)
	RegisterUser(user *entities.User, passwordSalt int) error
	LoginUser(user *entities.User) (string,error)
	VerifyUser(token string) error
	RequestForgetPassword(token string) error
	ResetPasswordByToken(token string, password string) error
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
func (usecase *authUseCase) RegisterUser(user *entities.User, passwordSalt int) error{
	cfg := config.LoadEnv()
	// 1. Generate ID unik untuk pengguna
	user.ID = utils.GenerateIDbyKSUID()

	// 2. Set peran pengguna sebagai User
	user.Role = string(entities.UserRole)
	user.IsVerified = false
	verificationToken := uuid.New().String()
	user.VerificationToken = verificationToken
	
	// 3. Hash password untuk keamanan
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), passwordSalt)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error hashing password: %v", err))
	}
	user.Password = string(hashedPassword)

	// 3. Kirim email verifikasi// Send verification email asynchronously
	go func() {
		err := smtp.SendEmailVerification(user.Email, verificationToken, cfg.EmailVerificationLink)
		if err != nil {
			log.Printf("Error sending verification email to %s: %v", user.Email, err)
		}
	}()


	// 4. Simpan data pengguna ke database melalui repository
	return usecase.authRepo.RegisterUser(user)
}

func (usecase *authUseCase) LoginUser(user *entities.User)(string, error){
	dbUser:= &entities.User{}
	dbUser.Email = user.Email
	err:= usecase.authRepo.LoginUser(dbUser)
	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return "",status.Errorf(codes.NotFound, "User Not Found")
		}
		return "",err
	}

	if dbUser.IsVerified == false{
		return "",status.Errorf(codes.Unauthenticated, "User Not Verified")
	}

	if dbUser.Role != string(entities.UserRole){
		return "",errors.New("User Not User Role")
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
func (usecase *authUseCase) VerifyUser(token string) error {
	return usecase.authRepo.VerifyUser(token)
}
func (usecase *authUseCase) RequestForgetPassword(email string) error {
	cfg := config.LoadEnv()
	user,err := usecase.authRepo.FindUserByEmail(email)
	if err != nil{
		return status.Errorf(codes.NotFound, "User Not Found")
	}
	forgetPasswordToken := uuid.New().String()
	user.ForgetPasswordToken = forgetPasswordToken
	err = usecase.authRepo.UpdateUserByEmail(email, user)
	if err != nil{
		return status.Errorf(codes.Internal, fmt.Sprintf("Error saving user: %v", err))
	}
	go func() {
		err := smtp.SendEmailForgetPassword(email, forgetPasswordToken, cfg.EmailForgetPasswordFrontendLink)
		if err != nil {
			log.Printf("Error sending forget password email to %s: %v", email, err)
		}
	}()

	return nil
}

func (usecase *authUseCase) ResetPasswordByToken(token string, password string) (error) {
	cfg:=config.LoadEnv()
	user,err := usecase.authRepo.FindOneUserByKey("forget_password_token", token)
	if err != nil{
		return status.Errorf(codes.NotFound, "User Not Found")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), cfg.BcryptSalt)
	if err != nil {
		return status.Errorf(codes.Internal, fmt.Sprintf("Error hashing password: %v", err))
	}
	user.Password = string(hashedPassword)
	err = usecase.authRepo.UpdateUserByEmail(user.Email, user)
	if err != nil{
		return status.Errorf(codes.Internal, fmt.Sprintf("Error saving user: %v", err))
	}
	return nil
}