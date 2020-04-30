package model

import (
	"github.com/jinzhu/gorm"
	uuid "github.com/satori/go.uuid"
	"time"
)

type BaseModel struct {
	PrimaryKeyID uint64    `json:"-" gorm:"primary_key;AUTO_INCREMENT;not null;unique"` //自增物理主键,不要用于逻辑处理
	ID           string    `gorm:"type:char(36);not null;unique"`                       // 逻辑主键，
	CreatedAt    time.Time `json:"-"`
	UpdatedAt    time.Time `json:"-"`
	//有DeletedAt的时候gorm就会软删除，没有就会真的删除
	DeletedAt *time.Time `sql:"index"`
}

func (model *BaseModel) BeforeCreate(scope *gorm.Scope) {
	u4 := uuid.NewV4().String()
	model.ID = u4
}
