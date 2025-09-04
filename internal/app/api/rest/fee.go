package rest

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/api/response"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/constant"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/model/vo"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/utils"
)

type IbcFeeController struct {
}

func (ctl *IbcFeeController) ChinFeeStatistics(c *gin.Context) {
	chain := c.Param("chain")
	if chain == "" {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("parameter chain is required"))
		return
	}
	exists, err := chainService.ChainExists(chain)
	if err != nil {
		c.JSON(http.StatusBadRequest, response.FailBadRequest(err.Error()))
		return
	}
	if !exists {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("this chain is not supported, please check or contact us by twitter(https://twitter.com/iobscan_ibc)"))
		return
	}

	var req vo.ChainFeeStatisticsReq
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.FailBadRequest(err.Error()))
		return
	}

	var startTime, endTime int64
	if req.StartDate != "" {
		startTimeStr := fmt.Sprintf("%s %s", req.StartDate, "00:00:00")
		startTimeParse, err := time.Parse(constant.DefaultTimeFormat, startTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FailBadRequest("invalid start_date"))
			return
		}
		startTime = startTimeParse.Unix()
	}

	if req.EndDate != "" {
		endTimeStr := fmt.Sprintf("%s %s", req.EndDate, "23:59:59")
		endTimeParse, err := time.Parse(constant.DefaultTimeFormat, endTimeStr)
		if err != nil {
			c.JSON(http.StatusBadRequest, response.FailBadRequest("invalid end_date"))
			return
		}
		endTime = endTimeParse.Unix()
	} else {
		_, endTime = utils.YesterdayUnix()
	}

	if startTime > endTime {
		c.JSON(http.StatusBadRequest, response.FailBadRequest("end_date must be greater than start_date"))
		return
	}

	res, e := feeService.ChainFeeStatistics(chain, startTime, endTime)
	if e != nil {
		c.JSON(response.HttpCode(e), response.FailError(e))
		return
	}

	c.JSON(http.StatusOK, response.Success(res))
}
