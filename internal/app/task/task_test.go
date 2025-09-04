package task

import (
	"context"
	"testing"
	"time"

	"github.com/irisnet/ibc-explorer-openapi/internal/app/conf"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/constant"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/global"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/repository"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/repository/cache"
	"github.com/sirupsen/logrus"
)

func TestMain(m *testing.M) {
	time.Local = time.UTC
	logrus.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat:   constant.DefaultTimeFormat,
		DisableHTMLEscape: true,
	})
	cache.InitRedisClient(conf.Redis{
		Addrs:    "127.0.0.1:6379",
		User:     "",
		Password: "",
		Mode:     "single",
		Db:       0,
	})
	repository.InitMgo(conf.Mongo{
		Url:      "mongodb://username:password@host:port/?connect=direct&authSource=database_name",
		Database: "iobscan-ibc_0805",
	}, context.Background())

	time.Local = time.UTC
	global.Config = &conf.Config{
		Task: conf.Task{},
	}
	m.Run()
}
