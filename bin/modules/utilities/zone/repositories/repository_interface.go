package repositories

import "gorm.io/gorm"

type (
	Options struct {
		DB *gorm.DB
	}

	GormInterface interface {
		FindOne() interface{}
	}
)