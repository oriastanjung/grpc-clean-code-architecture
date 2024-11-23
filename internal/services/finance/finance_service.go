package services

import (
	"context"

	usecase "github.com/oriastanjung/grpc_server/internal/usecase/finance"
	pb "github.com/oriastanjung/grpc_server/proto/finance"
)

type FinanceService interface {
	CreateFinance(ctx context.Context, input *pb.CreateFinanceRequest) error
	GetFinance(ctx context.Context,input *pb.GetFinanceRequest) (*pb.GetFinanceResponse, error)
	UpdateFinance(ctx context.Context, input *pb.UpdateFinanceRequest) error
	ListFinance(ctx context.Context, input *pb.ListFinanceRequest) (*pb.ListFinanceResponse, error)
	DeleteFinance(ctx context.Context, input *pb.DeleteFinanceRequest) error
}

type financeService struct {
	usecase usecase.FinanceUseCase
}

func NewFinanceService(usecase usecase.FinanceUseCase) FinanceService {
	return &financeService{
		usecase: usecase,
	}
}


func (s *financeService) CreateFinance(ctx context.Context, input *pb.CreateFinanceRequest) error {
	return s.usecase.CreateFinance(input)
}

func (s *financeService) GetFinance(ctx context.Context, input *pb.GetFinanceRequest) (*pb.GetFinanceResponse, error) {
	return s.usecase.GetFinance(input)
}

func (s *financeService) UpdateFinance(ctx context.Context, input *pb.UpdateFinanceRequest) error {
	return s.usecase.UpdateFinance(input)
}

func (s *financeService) ListFinance(ctx context.Context, input *pb.ListFinanceRequest) (*pb.ListFinanceResponse, error) {
	return s.usecase.ListFinance(input)
}

func (s *financeService) DeleteFinance(ctx context.Context, input *pb.DeleteFinanceRequest) error {
	return s.usecase.DeleteFinance(input)
}