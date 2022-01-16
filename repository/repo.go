package repository

import (
	"club-service/model"
	"context"
	log "go-micro.dev/v4/logger"

	"gorm.io/gorm"
)

type IRepository interface {
	Get(ctx context.Context, id int32) (*model.Club, error)
	Save(ctx context.Context, club *model.Club) (*model.Club, error)
	Update(ctx context.Context, club *model.Club, id int64) (*model.Club, error)
	FindByField(context.Context, string, string, string) (*model.Club, error)
	Delete(ctx context.Context, id int64) (string, error)
	GetAll(ctx context.Context) ([]*model.Club, error)
}

func (c *Club) Get(_ context.Context, id int32) (*model.Club, error) {
	club := &model.Club{}
	c.Db.First(club, id)
	return club, nil
}

func (c *Club) Save(_ context.Context, club *model.Club) (*model.Club, error) {
	err := c.Db.Create(club).Error
	if err != nil {
		log.Errorf("c.DB.Create().Error error: %v", err)
		return nil, err
	}
	return club, nil
}

func (c *Club) Update(_ context.Context, club *model.Club, id int64) (*model.Club, error) {
	updateClub := &model.Club{}
	c.Db.First(updateClub, id)
	updateClub.Founded = club.Founded
	updateClub.Name = club.Name
	updateClub.LeagueName = club.LeagueName
	updateClub.Manager = club.Manager
	c.Db.Save(updateClub)
	return updateClub, nil
}

func (c *Club) FindByField(_ context.Context, s string, s2 string, s3 string) (*model.Club, error) {
	if len(s3) == 0 {
		s3 = "*"
	}
	club := &model.Club{}
	if err := c.Db.Select(s3).Where(s+" = ?", s2).First(club).Error; err != nil {
		return nil, err
	}
	return club, nil
}

func (c *Club) Delete(_ context.Context, id int64) (string, error) {
	err := c.Db.Unscoped().Delete(&model.Club{}, id).Error
	if err != nil {
		return "Sorry but couldn't delete this entry", err
	}
	return "Entry deleted permanently", nil
}

func (c *Club) GetAll(ctx context.Context) ([]*model.Club, error) {
	var clubs []*model.Club
	err := c.Db.Find(&clubs).Error
	if err != nil {
		log.Errorf("DB Find error: %v", err)
		return nil, err
	}
	return clubs, nil
}

type Club struct {
	Db *gorm.DB
}
