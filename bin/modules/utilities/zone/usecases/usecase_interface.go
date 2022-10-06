package usecases

import (
	"ofa/bin/modules/utilities/zone/models"
	res "ofa/bin/pkg/response"

	"ofa/bin/modules/utilities/zone/repositories"
)

type (
	Options struct {
		Hallo string
		Gorm  repositories.GormInterface
	}

	Interface interface {
		GetList(*models.ReqGetList) (res.SendData, error)
	}
)
