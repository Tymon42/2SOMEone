package core

import (
	"context"
)

type (
	Message struct {
		//gorm.Model
		ID        uint   `gorm:"primarykey"`
		Timestamp int64  `gorm:"default:0"`
		Phone     string `gorm:"size:14;index" json:"phone,omitempty"`
		Type      uint8  `gorm:"type:tinyint(1);default:0" json:"type,omitempty"`
		Content   string `gorm:"size:512"`
		Code      string `gorm:"size:6"`
	}

	MessageStore interface {
		Create(ctx context.Context, message *Message) error
		Save(ctx context.Context, message *Message) error
		Find(ctx context.Context, message_id uint, phone string) (*Message, error)
	}
)
