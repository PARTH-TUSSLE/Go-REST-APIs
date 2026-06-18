package main

import (
	"REST-api/internal/config"
	"REST-api/internal/database"
	"REST-api/internal/handlers"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {

	var cfg *config.Config
	var err error

	cfg, err = config.Load()

	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	var pool *pgxpool.Pool

	pool, err = database.Connect(cfg.DatabaseURL)

	if err != nil {
		log.Fatal("Failed to connect to DB:", err)
	}

	defer pool.Close()

	var router *gin.Engine = gin.Default()
	
	router.SetTrustedProxies(nil)

	router.GET("/", func( ctx *gin.Context ) {
		ctx.JSON(200, gin.H{
			"message": "API server is running",
			"status": "SUCCESS !",
			"database": "connected !",
		})
	} )

	router.POST("/todos", handlers.CreateTodoHandler(pool))

	router.Run(":"+ cfg.Port)
}