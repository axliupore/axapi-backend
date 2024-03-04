package initialize

import (
	"fmt"
	"github.com/axliupore/axapi/axapi-backend/global"
	"github.com/axliupore/axapi/axapi-backend/router"
)

func InitRouter() {
	routers := router.Router()
	address := fmt.Sprintf(":%d", global.Config.Server.Port)
	routers.Run(address)
}
