package repository

import (
	"club-service/model"
	"context"

	log "go-micro.dev/v4/logger"

	"gorm.io/gorm"
)

type IRepository interface {
	Get(ctx context.Context, id interface{}) (interface{}, error)
	Save(ctx context.Context, club interface{}) (interface{}, error)
	Update(ctx context.Context, club interface{}, id interface{}) (interface{}, error)
	Delete(ctx context.Context, id interface{}) error
	GetAll(ctx context.Context) (interface{}, error)
}

func (c *Club) Get(_ context.Context, id interface{}) (interface{}, error) {
	dbID := id.(uint)
	var club *model.Club
	err := c.Db.First(&club, dbID).Error
	if err != nil {
		log.Errorf("Entity retrieve from DB error: %v", err)
		return nil, err
	}
	return club, nil
}

func (c *Club) Save(_ context.Context, club interface{}) (interface{}, error) {
	clubToSave := club.(*model.Club)
	err := c.Db.Create(&clubToSave).Error
	if err != nil {
		log.Errorf("Entity save in DB error: %v", err)
		return nil, err
	}
	return clubToSave, nil
}

func (c *Club) Update(ctx context.Context, club interface{}, id interface{}) (interface{}, error) {
	dbID := id.(uint)
	clubToUpdate := club.(*model.Club)
	var updatedClub *model.Club
	err := c.Db.First(&updatedClub, dbID).Error
	if err != nil {
		log.Errorf("Entity retrieve from DB error: %v", err)
		return nil, err
	}
	updatedClub.Founded = clubToUpdate.Founded
	updatedClub.Name = clubToUpdate.Name
	updatedClub.LeagueName = clubToUpdate.LeagueName
	updatedClub.Manager = clubToUpdate.Manager
	updatedClub.Capacity = clubToUpdate.Capacity
	updatedClub.Ground = clubToUpdate.Ground
	updatedClub.LeaguePosition = clubToUpdate.LeaguePosition
	updatedClub.TimesLeagueWon = clubToUpdate.TimesLeagueWon
	updatedClub.LastLeagueWon = clubToUpdate.LastLeagueWon
	err = c.Db.Save(&updatedClub).Error
	if err != nil {
		log.Errorf("Updated entity save in DB error: %v", err)
		return nil, err
	}
	return updatedClub, nil
}

func (c *Club) Delete(_ context.Context, id interface{}) error {
	dbID := id.(uint)
	err := c.Db.Unscoped().Delete(&model.Club{}, dbID).Error
	if err != nil {
		log.Errorf("Entity delete in DB error: %v", err)
		return err
	}
	return nil
}

func (c *Club) GetAll(ctx context.Context) (interface{}, error) {
	var clubs []*model.Club
	err := c.Db.Find(&clubs).Error
	if err != nil {
		log.Errorf("Entites retrieve from DB error: %v", err)
		return nil, err
	}
	return clubs, nil
}

type Club struct {
	Db *gorm.DB
}
