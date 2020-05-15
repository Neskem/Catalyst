package main

import (
	v1 "catalyst.Go/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {
	err1 := godotenv.Load()
	if err1 != nil {
		panic(err1)
	}
	app := gin.Default()
	v1.ApplyRoutes(app)
	port := os.Getenv("PORT")
	err2 := app.Run(":" + port)
	if err2 != nil {
		panic(err2)
	}
}
