package usecases

import (
	"fmt"
	"ofa/bin/modules/utilities/zone/models"
	res "ofa/bin/pkg/response"
)

func (opt *Options) GetList(payload *models.ReqGetList) (result res.SendData, err error) {
	result.Data = payload.Param.Zone
	var coba = opt.Gorm.FindOne()
	fmt.Println(coba)

	return result, nil
}
