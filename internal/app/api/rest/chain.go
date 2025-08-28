package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/api/response"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/model/vo"
)

type ChainController struct {
}

func (ctl *ChainController) List(c *gin.Context) {
	list, e := chainService.List()
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(list))
}

func (ctl *ChainController) IbcChainsNum(c *gin.Context) {
	resp, e := chainService.IbcChainsNum()
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(resp))
}

func (ctl *ChainController) IbcChainsVolume(c *gin.Context) {
	var req vo.IbcChainsVolumeReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.FailBadRequest(err.Error()))
		return
	}

	if req.Chain != "" {
		exists, err := chainService.ChainExists(req.Chain)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FailBadRequest(err.Error()))
			return
		}
		if !exists {
			c.JSON(http.StatusBadRequest, response.FailBadRequest("this chain is not supported, please check or contact us by twitter(https://twitter.com/iobscan_ibc)"))
			return
		}
	}

	resp, e := chainService.IbcChainsVolume(req.Chain)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(resp))
}

func (ctl *ChainController) IbcChainsActive(c *gin.Context) {
	resp, e := chainService.IbcChainsActive()
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(resp))
}
