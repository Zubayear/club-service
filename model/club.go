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

type ClubBuilder struct {
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

func NewClubBuilder() *ClubBuilder {
	return &ClubBuilder{}
}

func (c *ClubBuilder) SetName(name string) *ClubBuilder {
	c.Name = name
	return c
}

func (c *ClubBuilder) SetFounded(founded int32) *ClubBuilder {
	c.Founded = founded
	return c
}

func (c *ClubBuilder) SetLeagueName(leagueName string) *ClubBuilder {
	c.LeagueName = leagueName
	return c
}

func (c *ClubBuilder) SetManager(manager string) *ClubBuilder {
	c.Manager = manager
	return c
}

func (c *ClubBuilder) SetCapacity(capacity int32) *ClubBuilder {
	c.Capacity = capacity
	return c
}

func (c *ClubBuilder) SetLeaguePosition(leaguePosition int32) *ClubBuilder {
	c.LeaguePosition = leaguePosition
	return c
}

func (c *ClubBuilder) SetTimesLeagueWon(timesLeagueWon int32) *ClubBuilder {
	c.TimesLeagueWon = timesLeagueWon
	return c
}

func (c *ClubBuilder) SetLastLeagueWon(lastLeagueWon int32) *ClubBuilder {
	c.LastLeagueWon = lastLeagueWon
	return c
}

func (c *ClubBuilder) SetGround(ground string) *ClubBuilder {
	c.Ground = ground
	return c
}

func (c *ClubBuilder) BuildClub() *Club {
	return &Club{
		Name:           c.Name,
		Founded:        c.Founded,
		LeagueName:     c.LeagueName,
		Manager:        c.Manager,
		Capacity:       c.Capacity,
		LeaguePosition: c.LeaguePosition,
		TimesLeagueWon: c.TimesLeagueWon,
		LastLeagueWon:  c.LastLeagueWon,
		Ground:         c.Ground,
	}
}
