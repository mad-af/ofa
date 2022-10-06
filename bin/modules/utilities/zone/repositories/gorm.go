package repositories

import "fmt"

func (opt *Options) FindOne() interface{} {
	var data map[string]interface{}
	var coba = opt.DB.Table("zones").Find(&data)
	fmt.Println(coba)

	return data
}