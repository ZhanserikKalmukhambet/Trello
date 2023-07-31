package app

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/config"
	"github.com/ZhanserikKalmukhambet/Trello/internal/handler"
	"github.com/ZhanserikKalmukhambet/Trello/internal/repository"
	"github.com/ZhanserikKalmukhambet/Trello/internal/service"
	"github.com/ZhanserikKalmukhambet/Trello/pkg/http_server"
	"log"
)

func Run(cfg *config.Config) error {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handler := handler.New(services)

	server := http_server.New(handler.InitRouter(),
		http_server.WithPort(cfg.HTTP.Port),
		http_server.WithReadTimeout(cfg.HTTP.ReadTimeout),
		http_server.WithWriteTimeout(cfg.HTTP.WriteTimeout),
		http_server.WithShutdownTimeout(cfg.HTTP.ShutdownTimeout))

	log.Println("server started")
	server.Start()

	return nil
}
