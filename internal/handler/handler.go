package handler

import "github.com/ZhanserikKalmukhambet/Trello/internal/service"

type Handler struct {
	srvs *service.Service
}

func New(srvs *service.Service) *Handler {
	return &Handler{srvs: srvs}
}
