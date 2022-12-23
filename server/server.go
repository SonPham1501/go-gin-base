package server

import "com.son.server.goginbase/configs"

func Init()  {
	config := configs.GetConfig()
	r := NewRouter()
	r.Run(":" + config.GetString("server.port"))
}