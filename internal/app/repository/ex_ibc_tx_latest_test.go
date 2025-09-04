package repository

import (
	"testing"

	"github.com/irisnet/ibc-explorer-openapi/internal/app/utils"
)

func TestAggrChainAddress(t *testing.T) {
	res, err := new(ExIbcTxRepo).AggrChainAddress(1672588800, 1672675199, true, false)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(utils.MarshalJsonIgnoreErr(res)))
}
