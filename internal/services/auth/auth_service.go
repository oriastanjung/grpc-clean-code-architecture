package services

import (
	"context"

	"github.com/oriastanjung/grpc_server/internal/entities"
	usecase "github.com/oriastanjung/grpc_server/internal/usecase/auth"
)

type AuthService interface{
	RegisterAdmin(ctx context.Context, user *entities.User, salt int) error
}

type authService struct{
	authUseCase usecase.AuthUseCase
}

func NewAuthService(authUseCase usecase.AuthUseCase) AuthService{
	return &authService{
		authUseCase: authUseCase,
	}
}

func (service *authService) RegisterAdmin(ctx context.Context, user *entities.User,salt int) error {
	return service.authUseCase.RegisterAdmin(user,salt)
}
