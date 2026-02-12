package main

import (
	"bioskop-management-gin/configs"
	routers "bioskop-management-gin/routers"
)

func main() {
	PORT := ":8080"

	closeDB := configs.ConnectDB()
	defer closeDB()
	routers.StartServer().Run(PORT)
}
