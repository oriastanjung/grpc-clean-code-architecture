package main

import (
	"log"
	"net"

	"github.com/oriastanjung/grpc_server/internal/config"
	"github.com/oriastanjung/grpc_server/internal/database"
	auth_server "github.com/oriastanjung/grpc_server/internal/grpc/auth"
	repository "github.com/oriastanjung/grpc_server/internal/repository/auth"
	services "github.com/oriastanjung/grpc_server/internal/services/auth"
	usecase "github.com/oriastanjung/grpc_server/internal/usecase/auth"
	pbAuth "github.com/oriastanjung/grpc_server/proto/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


func main(){
	// load env
	config := config.LoadEnv()
	var port string = config.Port
	var addr string = "0.0.0.0:"+port

	// load DB
	database.InitDB()

	// Handle graceful shutdown
	database.GracefulShutdown()

	// auth service
	authRepository := repository.NewAuthRepository(database.DB)
	authUseCase := usecase.NewAuthUseCase(authRepository)
	authService := services.NewAuthService(authUseCase)
	authServer := auth_server.NewAuthServer(authService,config.BcryptSalt)
	// end auth service
	listener,err := net.Listen("tcp",addr)
    if err != nil{
        log.Printf("Error on Listening %v\n",err)
    }

    defer listener.Close()
	log.Printf("listening on %s\n", addr)

    

    serverInstance := grpc.NewServer()

    pbAuth.RegisterAuthServiceRoutesServer(serverInstance, authServer)
    reflection.Register(serverInstance)
    if err = serverInstance.Serve(listener); err != nil{
        log.Fatalf("Failed on Serve %v\n",err)
    }
}