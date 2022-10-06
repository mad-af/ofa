package hendlers

import (
	"github.com/labstack/echo/v4"

	"ofa/bin/middleware"
	"ofa/bin/modules/utilities/zone/models"
	"ofa/bin/modules/utilities/zone/repositories"
	"ofa/bin/modules/utilities/zone/usecases"
	res "ofa/bin/pkg/response"
	"ofa/bin/pkg/utils"
	db "ofa/bin/repositories"
)

var gorm repositories.GormInterface = &repositories.Options{
	DB: db.InitPostgre(),
}

var u usecases.Interface = &usecases.Options{
	Hallo: "hahah",
	Gorm: gorm,
}

func Init(g *echo.Group) {
	g.GET("/v1/zone/:zone", GetList, middleware.BasicAuth())
	g.GET("/v1/zone/:zone/:id", GetList, middleware.BasicAuth())
}

func GetList(c echo.Context) error {
	var req = new(models.ReqGetList)
	if err := utils.BindValidate(c, req); err != nil {
		return err
	}

	var data, err = u.GetList(req)
	if err != nil {
		return err
	}
	return res.Reply(&data, c)
}
