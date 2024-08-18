package main

import (
	"log"
	"os"

	"github.com/aiwanmaswood/finance-manage/config"
	routes "github.com/aiwanmaswood/finance-manage/routers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("unable to load environment file")
	}
	config.DbConnect()
	router := gin.New()
	router.Use(gin.Logger())

	routes.UserRouter(router)

	router.Run(os.Getenv("PORT"))
}
