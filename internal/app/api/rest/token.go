package rest

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/api/response"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/model/vo"
)

type TokenController struct {
}

func (ctl *TokenController) PopularSymbols(c *gin.Context) {
	var req vo.PopularSymbolsReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.FailBadRequest(err.Error()))
		return
	}

	minHops, err := strconv.ParseInt(req.MinHops, 10, 64)
	if err != nil || minHops < 0 {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("invalid min_hops value"))
		return
	}
	minReceiveTxs, err := strconv.ParseInt(req.MinReceiveTxs, 10, 64)
	if err != nil || minReceiveTxs < 0 {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("invalid min_receive_txs value"))
		return
	}

	res, e := tokenService.PopularSymbols(int(minHops), minReceiveTxs)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
