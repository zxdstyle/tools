package models

import (
	"github.com/golang-module/carbon"
	"gorm.io/gorm"
)

var DB *gorm.DB

type DateModel struct {
	CreatedAt carbon.ToDateTimeString `json:"created_at" gorm:"column:created_at"`
	UpdatedAt carbon.ToDateTimeString `json:"updated_at" gorm:"column:updated_at"`
}
