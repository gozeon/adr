package model

import (
	"gorm.io/gorm"
)

type Project struct {
	gorm.Model
	Title       string `gorm:"size:50;index" binding:"required" json:"title" validate:"required,min=3,max=50"`
	Description string `gorm:"size:120" json:"description" validate:"max=120"`
}

type Record struct {
	gorm.Model
	Title       string `gorm:"size:50;index" binding:"required" json:"title" validate:"required,min=3,max=50"`
	Description string `gorm:"type:text" json:"description" validate:"required"`
	Status      uint   `gorm:"index" json:"status" desc:"提议1, 通过2, 完成3, 已弃用4, 已取代5" validate:"required,oneof=1 2 3 4 5"`
	ProjectId   uint   `json:"project_id" binding:"required" validate:"required"`
}

type Comment struct {
	gorm.Model
	Description string `gorm:"type:text" json:"description" validate:"required"`
	RecordId    uint   `json:"record_id" validate:"required"`
}
