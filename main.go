package main

import (
	"bioskop-management-gin/configs"
	"bioskop-management-gin/databases"
	routers "bioskop-management-gin/routers"
)

func main() {
	PORT := ":8080"

	closeDB := configs.ConnectDB()
	defer closeDB()

	databases.RunMigration()
	routers.StartServer().Run(PORT)
}
