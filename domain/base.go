package domain

import (
	"time"
)

type BaseModel struct {
	ID        int64     `json:"id" db:"id"`
	CreatedAt time.Time `json:"created_at" bson:"created_at" db:"created_at"`
	UpdatedAt time.Time `json:"updated_at" bson:"updated_at" db:"updated_at"`
	IsDeleted bool      `json:"is_deleted" bson:"is_deleted" db:"is_deleted" gorm:"default:false"`
}
