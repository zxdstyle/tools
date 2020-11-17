package models

type Roles struct {
	Id     int    `gorm:"column:id;primary_key;AUTO_INCREMENT" json:"id"`
	Name   string `gorm:"column:name;NOT NULL" json:"name"`     // 角色名称
	Slug   string `gorm:"column:slug;NOT NULL" json:"slug"`     // 角色标识
	Status int    `gorm:"column:status;NOT NULL" json:"status"` // 角色状态

	DateModel
}

func (*Roles) TableName() string {
	return "roles"
}
