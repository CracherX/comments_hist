package app

import "gorm.io/gorm"

type App struct {
	db *gorm.DB
}
