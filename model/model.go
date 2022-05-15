package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type BaseModel struct {
	Id        uint64         `gorm:"primary_key;column:id;AUTO_INCREMENT;comment:序号;unique" `
	UUID      string         `gorm:"not null;column:uuid;unique" json:"-"`
	CreatedAt time.Time      `gorm:"column:createdAt" json:"-"`
	UpdatedAt time.Time      `gorm:"column:updatedAt" json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"column:deletedAt" sql:"index" json:"-"`
}

func (bm *BaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	bm.UUID = uuid.New().String()
	return
}

func (bm *BaseModel) CreatedAtDate() string {
	return bm.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}
