package models

type (
	ReqGetList struct {
		Param struct {
			Zone string `param:"zone" validate:"required"`
			ID   int    `param:"id"`
		}
	}
)
