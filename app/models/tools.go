package models

type Tools struct {
	Id          uint   `json:"id" gorm:"column:id"`
	Name        string `json:"name" gorm:"column:name;not null"`
	Icon        string `json:"icon" gorm:"column:icon"`
	Description string `json:"description" gorm:"column:description;not null"`

	DateModel
}
