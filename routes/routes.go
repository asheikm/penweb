package routes

import (
	"penweb/handlers"

	"github.com/gin-gonic/gin"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc gin.HandlerFunc
}

func GetRoutes(scanHandler *handlers.ScanHandler) []Route {
	routes := []Route{
		{
			Method:      "POST",
			Path:        "/scan",
			HandlerFunc: scanHandler.Scan,
		},
		{
			Method:      "GET",
			Path:        "/scan/result",
			HandlerFunc: scanHandler.GetComplatedScanResult,
		},
	}
	return routes
}
