package main

import (
	"club-service/handler"
	"club-service/model"
	pb "club-service/proto"
	"club-service/repository"

	"gorm.io/gorm"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"gorm.io/driver/mysql"
)

var (
	service = "club-service"
	version = "latest"
)

func main() {
	// setup db
	db := setUpDB()

	clubService := &handler.ClubService{DB: db}

	// Create service
	srv := micro.NewService(
		micro.Name(service),
		micro.Version(version),
	)
	srv.Init()

	// Register handler
	pb.RegisterClubServiceHandler(srv.Server(), clubService)

	// p := players.NewPlayersService("players", client.NewClient())

	// rsp, err := p.Get(context.TODO(), &players.PlayerRequest{
	// 	Id: "61e2ebe89f5bed7251ddf3f3",
	// })
	// if err != nil {
	// 	log.Errorf("Client call error: %v", err)
	// }
	// fmt.Println("response", rsp.Player)
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}

func setUpDB() repository.IRepository {
	dsn := "root:root@tcp(127.0.0.1:3306)/?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Exec("CREATE DATABASE IF NOT EXISTS " + "clubdb").Exec("USE " + "clubdb")
	if err != nil {
		log.Errorf("gorm.Open error: %v", err)
		panic("failed to connect database")
	}
	err = db.AutoMigrate(&model.Club{})
	if err != nil {
		log.Errorf("db.AutoMigrate() error: %v", err)
		panic("failed to auto migrate")
	}
	return &repository.Club{Db: db}
}
