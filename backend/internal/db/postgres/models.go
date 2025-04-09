package postgres

import (
	"time"

	"github.com/google/uuid"
)

type ApiKey struct {
	Uid            uuid.UUID    `gorm:"type:uuid;primary_key"`
	ApiKey         string       `gorm:"size:64;not null;uniqueIndex"`
	TotalCallCount int64        `gorm:"not null;default:0"`
	TotalCallLimit int64        `gorm:"not null"`
	PerMinuteLimit int          `gorm:"not null"`
	CreatedAt      time.Time    `gorm:"type:timestamp;not null"`
	UpdatedAt      time.Time    `gorm:"type:timestamp;not null"`
	Paths          []APIKeyPath `gorm:"foreignKey:ApiKeyID;references:Uid;constraint:OnDelete:CASCADE"`
}

func (ApiKey) TableName() string {
	return "api_key"
}

type APIKeyPath struct {
	Uid       uuid.UUID `gorm:"type:uuid;primary_key"`
	ApiKeyID  uuid.UUID `gorm:"type:uuid;not null;index"`
	GlobPath  string    `gorm:"size:255;not null"`
	CreatedAt time.Time `gorm:"type:timestamp;not null"`
	UpdatedAt time.Time `gorm:"type:timestamp;not null"`
}

func (APIKeyPath) TableName() string {
	return "api_key_path"
}