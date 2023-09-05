package controller

import (
	"context"
	"net/http"

	"github.com/arturzhamaliyev/p2p-transfer/internal/dto"
	"github.com/arturzhamaliyev/p2p-transfer/pkg/utils"
	"github.com/gin-gonic/gin"
)

//	{
//	    "from_account_id": 1,
//	    "to_account_id": 2,
//	    "amount": 10,
//	    "currency": "KZT"
//	}
func (h *Handler) MakeTransfer(ctx *gin.Context) {
	var transfer dto.TransferRequest

	err := ctx.BindJSON(&transfer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, utils.Response{Message: err.Error()})
		return
	}

	err = h.Service.Transfer.MakeTransfer(context.Background(), transfer)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, utils.Response{Message: err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, utils.Response{Message: "OK"})
}
