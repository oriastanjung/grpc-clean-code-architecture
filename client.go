package main

import (
	"context"
	"log"
	"time"

	pbAuth "github.com/oriastanjung/grpc_server/proto/auth"
	"google.golang.org/grpc"
)

func main() {
	// Server address
	addr := "localhost:2701" // Replace with your server address and port

	// Set up a connection to the gRPC server
	conn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// Create a client for the AuthServiceRoutes service
	authClient := pbAuth.NewAuthServiceRoutesClient(conn)

	// Context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// Example login request
	loginReq := &pbAuth.LoginRequest{
		Email: "oriastan999@gmail.com",
		Password: "orias027",
	}

	// Call the Login method
	loginResp, err := authClient.LoginAdmin(ctx, loginReq)
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}

	// Handle the response
	log.Printf("Login successful: %v", loginResp)
}
