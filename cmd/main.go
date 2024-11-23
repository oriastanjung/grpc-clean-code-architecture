package main

import (
	"log"
	"net"

	"github.com/oriastanjung/grpc_server/internal/config"
	"github.com/oriastanjung/grpc_server/internal/database"
	serverAuth "github.com/oriastanjung/grpc_server/internal/grpc/auth"
	serverFinance "github.com/oriastanjung/grpc_server/internal/grpc/finance"
	repositoryAuth "github.com/oriastanjung/grpc_server/internal/repository/auth"
	repositoryFinance "github.com/oriastanjung/grpc_server/internal/repository/finance"
	servicesAuth "github.com/oriastanjung/grpc_server/internal/services/auth"
	servicesFinance "github.com/oriastanjung/grpc_server/internal/services/finance"
	usecaseAuth "github.com/oriastanjung/grpc_server/internal/usecase/auth"
	usecaseFinance "github.com/oriastanjung/grpc_server/internal/usecase/finance"
	pbAuth "github.com/oriastanjung/grpc_server/proto/auth"
	pbFinance "github.com/oriastanjung/grpc_server/proto/finance"
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
	authRepository := repositoryAuth.NewAuthRepository(database.DB)
	authUseCase := usecaseAuth.NewAuthUseCase(authRepository)
	authService := servicesAuth.NewAuthService(authUseCase)
	authServer := serverAuth.NewAuthServer(authService,config.BcryptSalt)
	// end auth service


	// finance service
	financeRepository := repositoryFinance.NewFinanceRepository(database.DB)
	financeUseCase := usecaseFinance.NewFinanceUseCase(financeRepository)
	financeService := servicesFinance.NewFinanceService(financeUseCase)
	fincanceServer := serverFinance.NewFinanceServer(financeService)
	// end finance service
	listener,err := net.Listen("tcp",addr)
    if err != nil{
        log.Printf("Error on Listening %v\n",err)
    }

    defer listener.Close()
	log.Printf("listening on %s\n", addr)

    

    serverInstance := grpc.NewServer()

    pbAuth.RegisterAuthServiceRoutesServer(serverInstance, authServer)
	pbFinance.RegisterFinanceRoutesServiceServer(serverInstance, fincanceServer)
    reflection.Register(serverInstance)
    if err = serverInstance.Serve(listener); err != nil{
        log.Fatalf("Failed on Serve %v\n",err)
    }
}