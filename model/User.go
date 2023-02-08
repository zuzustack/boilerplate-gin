package model

import (
	"zuzustack/learn/boilerplate/utils"

	"gorm.io/gorm"
)

type user struct{
	Id int
	Name string
}

func User() *gorm.DB {
	return utils.DB().Model(user{})
}
