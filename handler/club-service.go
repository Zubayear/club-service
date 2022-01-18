package handler

import (
	players "club-service/client"
	"club-service/model"
	pb "club-service/proto"
	"club-service/repository"
	"context"
	"go-micro.dev/v4/client"
	log "go-micro.dev/v4/logger"
	mapper "gopkg.in/jeevatkm/go-model.v1"
)

type ClubService struct {
	DB repository.IRepository
}

func (c *ClubService) Save(ctx context.Context, request *pb.SaveRequest, response *pb.SaveResponse) error {
	log.Infof("Received ClubService.Save request: %v", request)
	clubBuilder := model.NewClubBuilder()
	club := clubBuilder.SetName(request.Name).
		SetFounded(request.Founded).
		SetLeagueName(request.LeagueName).
		SetManager(request.Manager).
		SetCapacity(request.Capacity).
		SetGround(request.Ground).
		SetLeaguePosition(request.LeaguePosition).
		SetTimesLeagueWon(request.TimesLeagueWon).
		SetLastLeagueWon(request.LastLeagueWon).BuildClub()
	res, err := c.DB.Save(ctx, club)
	if err != nil {
		return err
	}
	savedClubFromDB := res.(*model.Club)
	clubToReturn := &pb.Club{}
	mapper.Copy(clubToReturn, savedClubFromDB)
	response.Club = clubToReturn
	return nil
}

func (c *ClubService) Get(ctx context.Context, request *pb.GetRequest, response *pb.GetResponse) error {
	log.Infof("Received ClubService.Get request: %v", request)
	res, err := c.DB.Get(ctx, uint(request.GetId()))
	if err != nil {
		return err
	}
	clubFromDB := res.(*model.Club)
	clubToReturn := &pb.Club{}
	mapper.Copy(clubToReturn, clubFromDB)

	client := players.NewPlayersService("players", client.DefaultClient)
	playersOfClub, err := client.Filter(context.Background(), &players.GetRequest{Name: clubToReturn.Name})
	if err != nil {
		log.Errorf("Players retrieve from service error: %v", err)
		return err
	}
	getResponse := &players.GetResponse{Players: playersOfClub.Players}
	clubToReturn.Players = getResponse
	response.Club = clubToReturn
	return nil
}

func (c *ClubService) Update(ctx context.Context, request *pb.UpdateRequest, response *pb.UpdateResponse) error {
	log.Infof("Received ClubService.Update request: %v", request)
	clubBuilder := model.NewClubBuilder()
	clubToUpdate := clubBuilder.SetName(request.Name).
		SetFounded(request.Founded).
		SetLeagueName(request.LeagueName).
		SetManager(request.Manager).
		SetCapacity(request.Capacity).
		SetGround(request.Ground).
		SetLeaguePosition(request.LeaguePosition).
		SetTimesLeagueWon(request.TimesLeagueWon).
		SetLastLeagueWon(request.LastLeagueWon).BuildClub()
	res, err := c.DB.Update(ctx, clubToUpdate, uint(request.GetId()))
	if err != nil {
		return err
	}
	updatedClubFromDB := res.(*model.Club)
	clubToReturn := &pb.Club{}
	mapper.Copy(clubToReturn, updatedClubFromDB)
	response.Club = clubToReturn
	return nil
}

func (c *ClubService) Delete(ctx context.Context, request *pb.DeleteRequest, response *pb.DeleteResponse) error {
	log.Infof("Received ClubService.Delete request: %v", request)
	err := c.DB.Delete(ctx, uint(request.GetId()))
	if err != nil {
		return err
	}
	return nil
}

func (c *ClubService) GetAll(ctx context.Context, request *pb.ClubsRequest, response *pb.ClubsResponse) error {
	log.Infof("Received ClubService.GetAll request: %v", request)
	clubsFromDB, err := c.DB.GetAll(ctx)
	if err != nil {
		return err
	}
	clubsToReturn := clubsFromDB.([]*model.Club)
	var clubsResponse []*pb.Club
	for _, v := range clubsToReturn {
		club := &pb.Club{}
		mapper.Copy(club, v)
		clubsResponse = append(clubsResponse, club)
	}
	response.Clubs = clubsResponse
	return nil
}
