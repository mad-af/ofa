package usecases

import (
	"ofa/bin/modules/utilities/zone/models"
	res "ofa/bin/pkg/response"
)

func (o *Context) GetList(payload *models.ReqGetList) (result res.SendData, err error) {
	result.Data = payload.Param.Zone

	return result, nil
}
