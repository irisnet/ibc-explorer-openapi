package task

import (
	"sync"

	"github.com/irisnet/ibc-explorer-openapi/internal/app/conf"
)

var (
	taskConf      conf.Task
	TaskMetricMap = new(sync.Map)
)

func LoadTaskConf(taskCfg conf.Task) {
	taskConf = taskCfg
}
