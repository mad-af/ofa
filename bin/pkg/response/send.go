package response

import (
	"github.com/labstack/echo/v4"
)

type (
	ReplySend struct {
		Err  bool        `json:"error"`
		Data interface{} `json:"data"`
	}

	ReplyMetaSend struct {
		Err  bool        `json:"error"`
		Data interface{} `json:"data"`
		Meta Meta        `json:"meta"`
	}

	SendData struct {
		Code int
		Data interface{} `json:"data"`
		Meta Meta        `json:"meta"`
	}

	Meta struct {
		Page      int `json:"page"`
		Quantity  int `json:"quantity"`
		TotalPage int `json:"total_page"`
		TotalData int `json:"total_data"`
	}
)

func Reply(data *SendData, c echo.Context) error {
	var final interface{}

	if data.Meta == (Meta{}) {
		final = ReplySend{
			Err: false,
			Data: data.Data,
		}
	} else {
		final = ReplyMetaSend{
			Err: false,
			Data: data.Data,
			Meta: data.Meta,
		}
	}

	return c.JSON(data.Code, final)
}