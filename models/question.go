package models

import "gorm.io/gorm"

type Question struct {
	gorm.Model
	Content    string `json:"content"`
	Difficulty uint8  `json:"difficulty" gorm:"check:difficulty >= 1 AND difficulty <= 5"`
}
