package controller

import (
	"context"
	"net/http"

	"github.com/arturzhamaliyev/p2p-transfer/internal/dto"
	"github.com/arturzhamaliyev/p2p-transfer/pkg/utils"
	"github.com/gin-gonic/gin"
)

//	{
//	    "amount": 100,
//	    "currency": "KZT"
//	}
func (h *Handler) CreateUser(ctx *gin.Context) {
	var userData dto.UserCreate

	err := ctx.BindJSON(&userData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
		return
	}

	err = h.Service.User.CreateUser(context.Background(), userData)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{Message: "OK"})
}
