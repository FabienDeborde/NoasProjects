package entity

import "github.com/jinzhu/gorm"

type Project struct {
	gorm.Model
	Title       string `json:"title" binding:"required" gorm:"type:varchar(100);"`
	Description string `json:"description" gorm:"type:varchar(1000);"`
	Link        string `json:"link" binding:"required" gorm:"type:varchar(100);"`
}
