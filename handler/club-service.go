package handler

import (
	"club-service/model"
	pb "club-service/proto"
	"club-service/repository"
	"context"
	log "go-micro.dev/v4/logger"
)

type ClubService struct {
	DB repository.IRepository
}

func (c *ClubService) Save(ctx context.Context, request *pb.SaveRequest, response *pb.SaveResponse) error {
	log.Infof("Received ClubService.Save request: %v", request)
	club := &model.Club{
		Name:       request.GetName(),
		Founded:    request.GetFounded(),
		LeagueName: request.GetLeagueName(),
		Manager:    request.GetManager(),
	}
	res, err := c.DB.Save(ctx, club)
	if err != nil {
		log.Errorf("c.DB.Save() error: %v", err)
		return err
	}
	savedClub := &pb.Club{
		Name:       res.Name,
		Founded:    res.Founded,
		LeagueName: res.LeagueName,
		Manager:    res.Manager,
	}
	response.Club = savedClub
	return nil
}

func (c *ClubService) Get(ctx context.Context, request *pb.GetRequest, response *pb.GetResponse) error {
	log.Infof("Received ClubService.Get request: %v", request)
	res, err := c.DB.Get(ctx, int32(request.Id))
	if err != nil {
		log.Errorf("c.DB.Get() error: %v", err)
		return err
	}
	club := &pb.Club{
		Name:       res.Name,
		Founded:    res.Founded,
		LeagueName: res.LeagueName,
		Manager:    res.Manager,
	}
	response.Club = club
	return nil
}

func (c *ClubService) Update(ctx context.Context, request *pb.UpdateRequest, response *pb.UpdateResponse) error {
	log.Infof("Received ClubService.Update request: %v", request)
	club := &model.Club{
		Name:       request.GetName(),
		Founded:    request.GetFounded(),
		LeagueName: request.GetLeagueName(),
		Manager:    request.GetManager(),
	}
	res, err := c.DB.Update(ctx, club, int64(request.Id))
	if err != nil {
		log.Errorf("c.DB.Update() error: %v", err)
		return err
	}
	updatedClub := &pb.Club{
		Name:       res.Name,
		Founded:    res.Founded,
		LeagueName: res.LeagueName,
		Manager:    res.Manager,
	}
	response.Club = updatedClub
	return nil
}

func (c *ClubService) Delete(ctx context.Context, request *pb.DeleteRequest, response *pb.DeleteResponse) error {
	log.Infof("Received ClubService.Delete request: %v", request)
	s, err := c.DB.Delete(ctx, int64(request.Id))
	if err != nil {
		log.Errorf("c.DB.Delete error: %v", err)
		return err
	}
	response.Message = s
	return nil
}

func (c *ClubService) GetAll(ctx context.Context, request *pb.ClubsRequest, response *pb.ClubsResponse) error {
	log.Infof("Received ClubService.GetAll request: %v", request)
	clubs, err := c.DB.GetAll(ctx)
	if err != nil {
		log.Errorf("Business layer GetAll error: %v", err)
		return err
	}
	var clubsResponse []*pb.Club
	for _, v := range clubs {
		club := &pb.Club{
			Name:       v.Name,
			Founded:    v.Founded,
			LeagueName: v.LeagueName,
			Manager:    v.Manager,
		}
		clubsResponse = append(clubsResponse, club)
	}
	response.Clubs = clubsResponse
	return nil
}
