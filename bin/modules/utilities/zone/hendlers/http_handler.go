package hendlers

import (
	"github.com/labstack/echo/v4"

	"ofa/bin/modules/utilities/zone/models"
	"ofa/bin/modules/utilities/zone/usecases"
	res "ofa/bin/pkg/response"
	"ofa/bin/pkg/utils"
)

var u usecases.Interface = &usecases.Context{
	Hallo: "hahah",
}

func Init(g *echo.Group) {
	g.GET("/v1/zone/:zone", GetList)
	g.GET("/v1/zone/:zone/:id", GetList)
}

func GetList(c echo.Context) error {
	var req = new(models.ReqGetList)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}

	var data, err = u.GetList(req); 
	if err != nil { return err } 
	return res.Reply(data, c)
}
