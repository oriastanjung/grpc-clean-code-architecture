package repository

import (
	"github.com/oriastanjung/grpc_server/internal/entities"
	"github.com/segmentio/ksuid"
	"gorm.io/gorm"
)


type FinanceRepository interface {
	CreateFinance(finance *entities.Finance) error
	GetFinance(id ksuid.KSUID) (*entities.Finance, error)
	ListFinance(businessId ksuid.KSUID) ([]*entities.Finance, error)
	UpdateFinance(finance *entities.Finance) error
	DeleteFinance(id ksuid.KSUID) error
}

type financeRepository struct {
	db *gorm.DB
}

func NewFinanceRepository(db *gorm.DB) FinanceRepository {
	return &financeRepository{
		db: db,
	}
}

func (repo *financeRepository) CreateFinance(finance *entities.Finance) error {
	return repo.db.Create(finance).Error
}

func (repo *financeRepository) GetFinance(id ksuid.KSUID) (*entities.Finance, error) {
	var finance entities.Finance
	err := repo.db.Where("id = ?", id).First(&finance).Error
	return &finance, err
}

func (repo *financeRepository) ListFinance(businessId ksuid.KSUID) ([]*entities.Finance, error) {
	var finances []*entities.Finance
	err := repo.db.Where("business_id = ?", businessId).Find(&finances).Error
	return finances, err
}

func (repo *financeRepository) UpdateFinance(finance *entities.Finance) error {
	return repo.db.Save(finance).Error
}

func (repo *financeRepository) DeleteFinance(id ksuid.KSUID) error {
	return repo.db.Where("id = ?", id).Delete(&entities.Finance{}).Error
}