package daos

import "github.com/google/wire"

// DaoSet daos DI
var DaoSet = wire.NewSet(
	IoTDaoSet,
	OptionDaoSet,
)
