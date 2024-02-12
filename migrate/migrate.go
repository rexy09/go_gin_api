package main

import (
	"go_gin_api/config"
	"go_gin_api/database"
	"go_gin_api/models"
)

func init()  {
config.LoadEnv()
	database.ConnectToDB()
	
}

func main()  {
	database.DB.AutoMigrate(&models.Post{})
}