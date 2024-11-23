package auth_server

import (
	"context"
	"fmt"

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





func (server *AuthServer) SignUpUser(ctx context.Context, input *pb.SignUpRequest) (*pb.SignUpResponse, error){
	newUser,err := entities.NewUser(input.Username, input.Email, input.Password, entities.UserRole)
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error on NewUser with err : %v",err))
	}

	err = server.authService.RegisterUser(context.Background(),newUser,server.salt)
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v",err))
		
	}

	return &pb.SignUpResponse{
		Message: "SignUpUser Successfully",
	},nil

}


func (server *AuthServer) LoginUser(ctx context.Context, input *pb.LoginRequest) (*pb.LoginResponse, error){
	
	token, err := server.authService.LoginUser(context.Background(), &entities.User{
		Email: input.Email,
		Password: input.Password,
	})
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v",err))
	}

	return &pb.LoginResponse{
		Message: "Login User Successfully",
		Token: token,
	},nil

}


func (server *AuthServer) VerifyUser(ctx context.Context, input *pb.VerifyUserRequest) (*pb.VerifyUserResponse, error){
	token := input.Token
	err := server.authService.VerifyUser(ctx, token)
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v",err))
	}

	return &pb.VerifyUserResponse{
		Message: "Verify User Successfully",
	},nil
}


func (server *AuthServer) RequestForgetPassword(ctx context.Context, input *pb.RequestForgetPasswordRequest) (*pb.RequestForgetPasswordResponse, error){

	err := server.authService.RequestForgetPassword(ctx, input.Email)
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v",err))
	}

	return &pb.RequestForgetPasswordResponse{
		Message: "Request Forget Password Successfully",
	},nil
}


func (server *AuthServer) ResetPasswordByToken(ctx context.Context, input *pb.ResetPasswordByTokenRequest) (*pb.ResetPasswordByTokenResponse, error){

	err := server.authService.ResetPasswordByToken(ctx, input.Token, input.Password)
	if err != nil{
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v",err))
	}

	return &pb.ResetPasswordByTokenResponse{
		Message: "Reset Password By Token Successfully",
	},nil
}