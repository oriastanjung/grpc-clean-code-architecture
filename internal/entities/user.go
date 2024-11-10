package entities

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"github.com/segmentio/ksuid"
)

type Role string

const (
	AdminRole  Role = "admin"
	UserRole   Role = "user"
)

type KSUIDArray []ksuid.KSUID

// Implement GORM's Valuer interface for KSUIDArray
func (a KSUIDArray) Value() (driver.Value, error) {
	return json.Marshal(a)
}

// Implement GORM's Scanner interface for KSUIDArray
func (a *KSUIDArray) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}
	return json.Unmarshal(bytes, a)
}

type User struct {
	ID                  ksuid.KSUID `gorm:"primary_key;not null" json:"id"`
	Username            string      `json:"username" gorm:"not null"`
	Email               string      `json:"email" gorm:"unique;not null"`
	Password            string      `json:"password" gorm:"not null"`
	Role                string      `json:"role" gorm:"type:text;not null;check:role IN ('admin', 'user')"` // Use CHECK constraint
	ProfilePicture      string      `json:"profile_picture" gorm:"default:''"`
	ProfilePictureUrl   string      `json:"profile_picture_url" gorm:"default:''"`
	IsVerified          bool        `json:"is_verified" gorm:"default:false"`
	VerificationToken   string      `json:"verification_token" gorm:"default:''"`
	ForgetPasswordToken string      `json:"forget_password_token" gorm:"default:''"`
	CreatedAt           time.Time   `json:"created_at"`
	UpdatedAt           time.Time   `json:"updated_at"`
}

func NewUser(username, email, password string, role Role) (*User, error) {
	if role != AdminRole && role != UserRole {
		return nil, errors.New("invalid role")
	}

	user := &User{
		ID:        ksuid.New(),
		Username:  username,
		Email:     email,
		Password:  password,
		Role:      string(role),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return user, nil
}