package usecases

import (
	"ofa/bin/modules/utilities/zone/helpers"
	"ofa/bin/modules/utilities/zone/models"
	res "ofa/bin/pkg/response"
)

func (opt *Options) GetList(payload *models.ReqGetList) (result res.SendData, err error) {
	result.Data = opt.Gorm.FindManyCommon(helpers.IDLength[payload.Param.Zone], payload.Param.ID)

	return result, nil
}
