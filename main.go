package main

import (
	"penweb/db"
	"penweb/handlers"
	"penweb/routes"
	"penweb/services"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	log.Info().Msg("Scan gin router initialized")
	r := gin.New()
	log.Info().Msg("Scan mongo db initialized")
	db.InitDB()

	scanService := services.NewDbScanService(db.DB)
	scanHandler := handlers.NewScanHandler(scanService)
	// r.Use(ginLogger())

	v1 := r.Group("/api/v1")
	log.Info().Msg("Initializing routes...")
	routes := routes.GetRoutes(scanHandler)
	for _, route := range routes {
		v1.Handle(route.Method, route.Path, route.HandlerFunc)
	}

	r.Run(":8081")
}
