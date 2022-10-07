package repositories

import "gorm.io/gorm"

type (
	Options struct {
		DB *gorm.DB
	}

	GormInterface interface {
		FindManyCommon(zone int, id string) interface{}
	}
)