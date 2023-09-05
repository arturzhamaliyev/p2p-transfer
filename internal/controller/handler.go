package controller

import (
	"github.com/arturzhamaliyev/p2p-transfer/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/knadh/koanf/v2"
)

type Handler struct {
	Config  *koanf.Koanf
	Router  *gin.Engine
	Service service.Service
}

func NewHandler(cfg *koanf.Koanf, router *gin.Engine, service service.Service) *Handler {
	return &Handler{
		Config:  cfg,
		Router:  router,
		Service: service,
	}
}
