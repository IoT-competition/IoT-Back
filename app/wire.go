//go:generate wire
//+build wireinject

package app

import (
	"github.com/IoTCompetition/IoBack/app/basic"
	"github.com/IoTCompetition/IoBack/app/controllers"
	"github.com/IoTCompetition/IoBack/app/daos"
	"github.com/IoTCompetition/IoBack/app/models"
	"github.com/IoTCompetition/IoBack/app/utils"
	"net/http"

	"github.com/IoTCompetition/IoBack/app/config"
	"github.com/IoTCompetition/IoBack/app/routers"
	"github.com/google/wire"
)

func BuildInjector() *http.Server {
	wire.Build(
		config.NewConfig,
		basic.LoggerSet,
		basic.NewDB,
		utils.GinxSet,
		daos.DaoSet,
		models.ModelSet,
		controllers.ControllerSet,
		routers.RouterSet,
		NewApp,
	)

	return nil
}
