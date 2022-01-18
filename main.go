package main

import (
	"club-service/handler"
	"club-service/model"
	pb "club-service/proto"
	"club-service/repository"
	"fmt"
	"github.com/asim/go-micro/plugins/config/encoder/yaml/v4"
	"go-micro.dev/v4/config"
	"go-micro.dev/v4/config/reader"
	"go-micro.dev/v4/config/reader/json"
	"go-micro.dev/v4/config/source/file"

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

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}

}

func setUpDB() repository.IRepository {
	host, err := loadConfig()
	if err != nil {
		log.Errorf("Could not load config: %v", err)
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/?charset=utf8mb4&parseTime=True&loc=Local",
		host.User, host.Password, host.Address, host.Port)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db.Exec("CREATE DATABASE IF NOT EXISTS " + host.Name).Exec("USE " + host.Name)
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

type Host struct {
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Name     string `json:"name"`
	User     string `json:"user"`
	Password string `json:"password"`
}

func loadConfig() (*Host, error) {
	enc := yaml.NewEncoder()
	c, err := config.NewConfig(config.WithReader(json.NewReader(reader.WithEncoder(enc))))
	if err != nil {
		log.Errorf("Error loading config %v", err)
		return nil, err
	}

	err = c.Load(file.NewSource(file.WithPath("./config.yaml")))
	if err != nil {
		return nil, err
	}

	host := &Host{}
	if err := c.Get("hosts", "database").Scan(host); err != nil {
		return nil, err
	}

	return host, nil
}
