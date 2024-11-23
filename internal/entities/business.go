package entities

import (
	"time"

	"github.com/segmentio/ksuid"
)

type Business struct {
	ID          ksuid.KSUID `gorm:"primary_key;not null"`
	Name        string      `gorm:"not null;index"` // Index on the name for fast search
	Description string      `gorm:"type:text"`
	UserId      ksuid.KSUID `gorm:"not null;index"` // Index to optimize queries filtering by UserId
	User        User        `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Finance     []Finance   `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	CreatedAt   time.Time   `gorm:"autoCreateTime;index"`
	UpdatedAt   time.Time   `gorm:"autoCreateTime;index"`
}

