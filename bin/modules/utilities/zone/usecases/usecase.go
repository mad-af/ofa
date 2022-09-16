package usecases

import (
	res "ofa/bin/pkg/response"
	"ofa/bin/modules/utilities/zone/models"
)

func (o *Context) GetList(payload *models.ReqGetList) (result *res.SendData, err error) {
	result = &res.SendData{}
	result.Data = payload.Param.Zone

	return result, nil
}
