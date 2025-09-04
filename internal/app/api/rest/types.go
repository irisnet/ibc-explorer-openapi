package rest

import (
	"github.com/irisnet/ibc-explorer-openapi/internal/app/service"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/task"
)

var (
	txService      service.ITxService      = new(service.TxService)
	chainService   service.IChainService   = new(service.ChainService)
	feeService     service.IFeeService     = new(service.FeeService)
	addressService service.IAddressService = new(service.AddressService)
	tokenService   service.ITokenService   = new(service.TokenService)
)

var (
	// task
	ibcTxFailLogTask         task.IBCTxFailLogTask
	iBCChainFeeStatisticTask task.IBCChainFeeStatisticTask
	ibcAddressStatisticTask  task.IBCAddressStatisticTask
)
