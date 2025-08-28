package monitor

import (
	"context"
	"testing"

	"github.com/irisnet/ibc-explorer-openapi/internal/app/conf"
	"github.com/irisnet/ibc-explorer-openapi/internal/app/repository"
)

func TestMain(m *testing.M) {
	repository.InitMgo(conf.Mongo{
		Url:      "mongodb://username:password@host:port/?authSource=database_name",
		Database: "database_name",
	}, context.Background())
	m.Run()
}
