package server

import (
	"com.son.server.goginbase/configs"
	"com.son.server.goginbase/routers"
)

func Init() {
	config := configs.GetConfig()
	r := routers.NewRouter()
	r.Run(":" + config.GetString("server.port"))
}
