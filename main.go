package main

import (
	"bioskop-management-gin/config"
	routers "bioskop-management-gin/routers"
)

func main() {
	PORT := ":8080"

	config.ConnectDB()

	routers.StartServer().Run(PORT)
}
