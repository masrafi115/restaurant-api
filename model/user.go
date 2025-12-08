package models

import "gorm.io/gorm"

var DB *gorm.DB

type UserInfo struct {
    gorm.Model
    Username string `json:"username"`
    Password string `json:"-"`
    Role     string `json:"role"`
}