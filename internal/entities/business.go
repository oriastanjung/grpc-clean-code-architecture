package entities

import (
	"time"

	"github.com/segmentio/ksuid"
)

type Business struct {
	ID   ksuid.KSUID `gorm:"primary_key;not null"`
	Name string      `gorm:"not null"`
	Description string	`gorm:"type:text"`
	UserId ksuid.KSUID `gorm:"not null"`
	User User `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Finance []Finance `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt time.Time   `gorm:"autoCreateTime"`
	UpdatedAt time.Time   `gorm:"autoCreateTime"`
}

