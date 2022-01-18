package handler

import (
	players "club-service/client/proto"
	"club-service/model"
	pb "club-service/proto"
	"club-service/repository"
	"context"
	"fmt"

	"go-micro.dev/v4/client"
	log "go-micro.dev/v4/logger"
)

type ClubService struct {
	DB repository.IRepository
}

func (c *ClubService) Save(ctx context.Context, request *pb.SaveRequest, response *pb.SaveResponse) error {
	log.Infof("Received ClubService.Save request: %v", request)
	club := &model.Club{
		Name:           request.Name,
		Founded:        request.Founded,
		LeagueName:     request.LeagueName,
		Manager:        request.Manager,
		Capacity:       request.Capacity,
		Ground:         request.Ground,
		LeaguePosition: request.LeaguePosition,
		TimesLeagueWon: request.TimesLeagueWon,
		LastLeagueWon:  request.LastLeagueWon,
	}
	res, err := c.DB.Save(ctx, club)
	if err != nil {
		return err
	}
	savedClubFromDB := res.(*model.Club)
	clubToReturn := &pb.Club{}
	response.Club = copyPropDest(savedClubFromDB, clubToReturn)
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
	value := copyPropDest(clubFromDB, clubToReturn)

	p := players.NewPlayersService("players", client.DefaultClient)

	rsp, err := p.Get(context.TODO(), &players.PlayerRequest{
		Id: "61e2ebe89f5bed7251ddf3f3",
	})
	if err != nil {
		log.Errorf("Client call error: %v", err)
	}
	fmt.Println("response", rsp.Player)
	value.Player = rsp.Player
	response.Club = value
	return nil
}

func (c *ClubService) Update(ctx context.Context, request *pb.UpdateRequest, response *pb.UpdateResponse) error {
	log.Infof("Received ClubService.Update request: %v", request)
	clubToUpdate := &model.Club{
		Name:           request.Name,
		Founded:        request.Founded,
		LeagueName:     request.LeagueName,
		Manager:        request.Manager,
		Capacity:       request.Capacity,
		LeaguePosition: request.LeaguePosition,
		TimesLeagueWon: request.TimesLeagueWon,
		LastLeagueWon:  request.LastLeagueWon,
		Ground:         request.Ground,
	}
	res, err := c.DB.Update(ctx, clubToUpdate, uint(request.GetId()))
	if err != nil {
		return err
	}
	updatedClubFromDB := res.(*model.Club)
	clubToReturn := &pb.Club{}
	response.Club = copyPropDest(updatedClubFromDB, clubToReturn)
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
		clubsResponse = append(clubsResponse, copyPropDest(v, club))
	}
	response.Clubs = clubsResponse
	return nil
}

func copyPropDest(src *model.Club, dst *pb.Club) *pb.Club {
	dst.Name = src.Name
	dst.Founded = src.Founded
	dst.LeagueName = src.LeagueName
	dst.Manager = src.Manager
	dst.Capacity = src.Capacity
	dst.LeaguePosition = src.LeaguePosition
	dst.TimesLeagueWon = src.TimesLeagueWon
	dst.LastLeagueWon = src.LastLeagueWon
	dst.Ground = src.Ground
	return dst
}
