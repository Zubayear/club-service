package model

import "gorm.io/gorm"

type Club struct {
	gorm.Model
	Name           string
	Founded        int32
	LeagueName     string
	Manager        string
	Capacity       int32
	LeaguePosition int32
	TimesLeagueWon int32
	LastLeagueWon  int32
	Ground         string
}
