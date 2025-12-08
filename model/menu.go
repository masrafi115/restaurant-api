package models

import "gorm.io/gorm"

var DB *gorm.DB


type OrderMenu struct {
    gorm.Model
    Name  string  `json:"item_name"`
    Price float64 `json:"price"`
}