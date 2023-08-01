package app

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/config"
	"github.com/ZhanserikKalmukhambet/Trello/internal/handler"
	"github.com/ZhanserikKalmukhambet/Trello/internal/repository"
	"github.com/ZhanserikKalmukhambet/Trello/internal/repository/pgrepo"
	"github.com/ZhanserikKalmukhambet/Trello/internal/service"
	"github.com/ZhanserikKalmukhambet/Trello/pkg/http_server"
	_ "github.com/lib/pq"
	"log"
)

func Run(cfg *config.Config) error {
	pg := pgrepo.Postgres{
		Port:     cfg.DB.Port,
		Host:     cfg.DB.Host,
		DBName:   cfg.DB.DBName,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		SSLMode:  "disable",
	}
	db, err := pgrepo.NewPostgresDB(pg)
	if err != nil {
		return err
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.New(services)

	server := http_server.New(handlers.InitRouter(),
		http_server.WithPort(cfg.HTTP.Port),
		http_server.WithReadTimeout(cfg.HTTP.ReadTimeout),
		http_server.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		http_server.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout))

	server.Start()
	log.Println("server started")

	return nil
}
