package main

import (
	"fmt"
	"log"
	"os"
	"quilink/internal/handlers"
	m "quilink/internal/models"
	"quilink/internal/repository"

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
	db, err := repository.ConnectToDB()
	if err != nil {
		log.Fatalf("error connecting to database: %v", err)
	}

	// auto migrate
	db.AutoMigrate(&m.User{}, &m.Blog{}, &m.Link{}, &m.Profile{})

	// handlers
	blogHandler := handlers.NewBlogHandler()
	handlers.GroupBlogHandlers(r, blogHandler)

	r.Run()
}
