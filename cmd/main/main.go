package main

import (
	"fmt"
	"os"
	"quilink/internal/handlers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	var r *gin.Engine
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("no .env file present: %v", err)
	}

	mode := os.Getenv("GIN_MODE")
	if mode == "release" {
		gin.SetMode(gin.ReleaseMode)
		r = gin.Default()
	} else {
		gin.SetMode(gin.DebugMode)
		r = gin.Default()
	}

	// database

	// handlers
	blogHandler := handlers.NewBlogHandler()
	handlers.GroupBlogHandlers(r, blogHandler)

	r.Run()
}
