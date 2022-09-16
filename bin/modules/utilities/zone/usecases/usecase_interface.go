package usecases

import (
	"ofa/bin/modules/utilities/zone/models"
	res "ofa/bin/pkg/response"
)

type (
	Context struct {
		Hallo string
	}

	Interface interface {
		GetList(*models.ReqGetList) (*res.SendData, error)
		// COba()
	}
)
