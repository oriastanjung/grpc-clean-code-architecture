package usecase

import (
	"github.com/oriastanjung/grpc_server/internal/entities"
	repository "github.com/oriastanjung/grpc_server/internal/repository/finance"
	"github.com/oriastanjung/grpc_server/internal/utils"
	pb "github.com/oriastanjung/grpc_server/proto/finance"
	"github.com/segmentio/ksuid"
)


type FinanceUseCase interface{
	CreateFinance(input *pb.CreateFinanceRequest) error
	GetFinance(input *pb.GetFinanceRequest) (*pb.GetFinanceResponse, error)
	ListFinance(input *pb.ListFinanceRequest) (*pb.ListFinanceResponse, error)
	UpdateFinance(input *pb.UpdateFinanceRequest) error
	DeleteFinance(input *pb.DeleteFinanceRequest) error
}

type financeUseCase struct {
	repo repository.FinanceRepository
}


func NewFinanceUseCase(repo repository.FinanceRepository) FinanceUseCase {
	return &financeUseCase{
		repo: repo,
	}
}


func (usecase *financeUseCase) CreateFinance(input *pb.CreateFinanceRequest) error {
	businessId,err := ksuid.Parse(input.BusinessId)
	if err != nil {
		return err
	}
	finance := &entities.Finance{
		ID: utils.GenerateIDbyKSUID(),
		BusinessId: businessId,
		Title: input.Title,
		Type: entities.FinanceType(input.Type),
		ProductName: input.ProductName,
		Amount: input.Amount,
		Notes: input.Notes,
	}
	return usecase.repo.CreateFinance(finance)
}

func (usecase *financeUseCase) GetFinance(input *pb.GetFinanceRequest) (*pb.GetFinanceResponse, error) {
	id,err := ksuid.Parse(input.Id)
	if err != nil {
		return nil, err
	}
	finance,err := usecase.repo.GetFinance(id)
	financeModel := &pb.FinanceModel{
		Id: finance.ID.String(),
		BusinessId: finance.BusinessId.String(),
		Title: finance.Title,
		Type: finance.Type.String(),
		ProductName: finance.ProductName,
		Amount: finance.Amount,
		Notes: finance.Notes,
	}
	result := &pb.GetFinanceResponse{
		Finance: financeModel,
	}
	return result, err
}	

func (usecase *financeUseCase) ListFinance(input *pb.ListFinanceRequest) (*pb.ListFinanceResponse, error) {		
	businessId,err := ksuid.Parse(input.BusinessId)
	if err != nil {
		return nil, err
	}
	finances,err:= usecase.repo.ListFinance(businessId)

	var financesModel []*pb.FinanceModel
	for _,finance := range finances{
		financeModel := &pb.FinanceModel{
			Id: finance.ID.String(),
			BusinessId: finance.BusinessId.String(),
			Title: finance.Title,
			Type: entities.FinanceType(finance.Type).String(),
			ProductName: finance.ProductName,
			Amount: finance.Amount,
			Notes: finance.Notes,
		}
		financesModel = append(financesModel, financeModel)
	}
	result := &pb.ListFinanceResponse{
		Finance: financesModel,
	}
	return result, err
}

func (usecase *financeUseCase) UpdateFinance(input *pb.UpdateFinanceRequest) error {
	financeId, err := ksuid.Parse(input.Id)
	if err != nil {
		return err
	}
	businessId,err := ksuid.Parse(input.BusinessId)
	if err != nil {
		return err
	}

	finance := &entities.Finance{
		ID: financeId,
		BusinessId: businessId,
		Title: input.Title,
		Type: entities.FinanceType(input.Type),
		ProductName: input.ProductName,
		Amount: input.Amount,
		Notes: input.Notes,
	}
	return usecase.repo.UpdateFinance(finance)
}

func (usecase *financeUseCase) DeleteFinance(input *pb.DeleteFinanceRequest) error {
	id,err := ksuid.Parse(input.Id)
	if err != nil {
		return err
	}
	return usecase.repo.DeleteFinance(id)
}