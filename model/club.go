package model

import "gorm.io/gorm"

type Club struct {
	gorm.Model
	Name       string
	Founded    int32
	LeagueName string
	Manager    string
}
