package models

type Language struct {
	Id      uint    `gorm:"column:id;NOT NULL;AUTO_INCREMENT;" json:"id"`
	Name    string  `gorm:"column:name;NOT NULL;" json:"name"`
	Ratings float64 `gorm:"column:ratings;NOT NULL;" json:"ratings"`
	Date    string

	DateModel
}

func (*Language) TableName() string {
	return "languages"
}
