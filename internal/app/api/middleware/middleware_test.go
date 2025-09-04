package middleware

import (
	"encoding/json"
	"strconv"
	"testing"
	"time"

	"github.com/irisnet/ibc-explorer-openapi/internal/app/model/vo"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/utils"
	"github.com/stretchr/testify/require"
)

func Test_Middleware(t *testing.T) {
	//for {
	apiKey := "your_api_key_here"
	apiSecret := "your_api_secret_here"
	timestamp := time.Now().Unix()
	host := "http://localhost:8000"
	uri := "/ibc/txs/BDBA021CF4939699208228E227699C98B48E9FF59C3CD517C2D88281131EA3CC?chain=irishubqa"

	signature := calculateSignature(uri, "", apiSecret, timestamp)
	header := map[string]string{
		"X-Api-Key":   apiKey,
		"X-Timestamp": strconv.FormatInt(timestamp, 10),
		"X-Signature": signature,
	}

	httpCode, bz, err := utils.HttpDo("GET", host+uri, nil, header)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bz))
	require.Equal(t, 200, httpCode)
	//}
}

func Test_MiddlewarePost(t *testing.T) {
	//for {
	apiKey := "your_api_key_here"
	apiSecret := "your_api_secret_here"
	timestamp := time.Now().Unix()
	host := "http://localhost:8000"
	uri := "/ibc/task/ibc_tx_fail_log_task"

	body := vo.TaskReq{
		StartTime:       1669803522,
		EndTime:         time.Now().Unix(),
		IsTargetHistory: false,
	}
	bz, _ := json.Marshal(body)
	signature := calculateSignature(uri, string(bz), apiSecret, timestamp)
	header := map[string]string{
		"Content-Type": "application/json; charset=utf-8",
		"X-Api-Key":    apiKey,
		"X-Timestamp":  strconv.FormatInt(timestamp, 10),
		"X-Signature":  signature,
	}

	httpCode, bz, err := utils.HttpDo("POST", host+uri, body, header)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(bz))
	require.Equal(t, 200, httpCode)
	//}
}

func TestName(t *testing.T) {
	sign := calculateSignature("ibc/txs/3115FB1C39C2156321C175974C9C7EFE9DC5009C2C7A2EF98EA2A70785E45B89", "", "your_api_secret_here", 1670229974)
	t.Log(sign)
}
