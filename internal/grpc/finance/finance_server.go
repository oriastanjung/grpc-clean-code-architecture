package finance_server

import (
	"context"
	"fmt"

	services "github.com/oriastanjung/grpc_server/internal/services/finance"
	pb "github.com/oriastanjung/grpc_server/proto/finance"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type FinanceServer struct {
	FinanceService services.FinanceService
	pb.FinanceRoutesServiceServer
}

func NewFinanceServer(financeService services.FinanceService) pb.FinanceRoutesServiceServer {
	return &FinanceServer{
		FinanceService: financeService,
	}
}



func (s *FinanceServer) CreateFinance(ctx context.Context, req *pb.CreateFinanceRequest) (*pb.CreateFinanceResponse, error) {
	finance := req	
	err := s.FinanceService.CreateFinance(ctx, finance)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}
	return &pb.CreateFinanceResponse{
		Message: "Create Finance Successfully",
	}, nil
}



func (s *FinanceServer) GetFinance(ctx context.Context, input *pb.GetFinanceRequest) (*pb.GetFinanceResponse, error){
	response,err := s.FinanceService.GetFinance(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}
	return response, nil
}


func (s *FinanceServer) ListFinance(ctx context.Context, input *pb.ListFinanceRequest) (*pb.ListFinanceResponse, error){
	response,err := s.FinanceService.ListFinance(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}
	return response, nil
}
func (s *FinanceServer) UpdateFinance(ctx context.Context, input *pb.UpdateFinanceRequest) (*pb.UpdateFinanceResponse, error){
	err := s.FinanceService.UpdateFinance(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}
	return &pb.UpdateFinanceResponse{
		Message: "Update Finance Successfully",
	}, nil
}
func (s *FinanceServer) DeleteFinance(ctx context.Context, input *pb.DeleteFinanceRequest) (*pb.DeleteFinanceResponse, error){
	err := s.FinanceService.DeleteFinance(ctx, input)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Error : %v", err))
	}
	return &pb.DeleteFinanceResponse{
		Message: "Delete Finance Successfully",
	}, nil
}