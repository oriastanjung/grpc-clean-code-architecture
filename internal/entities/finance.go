package entities

import (
	"time"

	"github.com/segmentio/ksuid"
)
type FinanceType string

const (
	Expense FinanceType = "expense"
	Income  FinanceType = "income"
)
func (ft FinanceType) String() string {
	return string(ft)
}

type Finance struct {
	ID          ksuid.KSUID `gorm:"primary_key;not null"`
	BusinessId  ksuid.KSUID `gorm:"not null"`
	Business    Business    `gorm:"foreignKey:BusinessId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Title       string      `gorm:"not null"`
	Type        FinanceType `gorm:"type:text;not null;check:type IN ('expense', 'income')"`
	ProductName string      `gorm:"not null"`
	Amount      float64     `gorm:"not null"`
	Notes       string      `gorm:"type:text"`
	CreatedAt   time.Time   `gorm:"autoCreateTime"`
	UpdatedAt   time.Time   `gorm:"autoUpdateTime"`
}