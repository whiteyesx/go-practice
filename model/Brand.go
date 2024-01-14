package model

import "gorm.io/gorm"

//TODO из чего рейтинг складывается?

type Brand struct {
	gorm.Model
	Name       string     `gorm:"column:name"`
	Categories []Category `gorm:"many2many:brand_categories"`
}
