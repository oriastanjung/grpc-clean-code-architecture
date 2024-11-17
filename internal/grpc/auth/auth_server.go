package auth_server

import (
	"context"
	"fmt"
	"log"

	"github.com/oriastanjung/grpc_server/internal/entities"
	services "github.com/oriastanjung/grpc_server/internal/services/auth"
	pb "github.com/oriastanjung/grpc_server/proto/auth"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type AuthServer struct{
	pb.AuthServiceRoutesServer
	authService services.AuthService
	salt int
}

func NewAuthServer(authService services.AuthService,salt int) *AuthServer{
	return &AuthServer{
		authService: authService,
		salt : salt,
	}
}



func (server *AuthServer) SignUpAdmin(ctx context.Context, input *pb.SignUpRequest) (*pb.SignUpResponse, error){
	log.Printf("SignUpAdmin Invoked with %v",input)

	newUser,err := entities.NewUser(input.Username, input.Email, input.Password, entities.AdminRole)
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error on NewUser with err : %v",err))
	}

	err = server.authService.RegisterAdmin(context.Background(),newUser,server.salt)
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v",err))
		
	}

	return &pb.SignUpResponse{
		Message: "SignUpAdmin Successfully",
	},nil

}


func (server *AuthServer) LoginAdmin(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error){
	log.Printf("LoginAdmin Invoked with %v",input)

	token, err := server.authService.LoginAdmin(context.Background(), &entities.User{
		Email: input.Email,
		Password: input.Password,
	})
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v",err))
	}

	return &pb.LoginResponse{
		Message: "Login Admin Successfully",
		Token: token,
	},nil

}