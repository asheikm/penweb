package handlers

import (
	"net/http"
	"penweb/models"
	"penweb/services"

	"github.com/gin-gonic/gin"
)

type ScanHandler struct {
	ScanService services.ScanService
}

func NewScanHandler(scanService services.ScanService) *ScanHandler {
	return &ScanHandler{
		ScanService: scanService,
	}
}

func (h *ScanHandler) Scan(c *gin.Context) {
	var request *models.ScanRequest
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Call the ScanService to perform the scan
	err := h.ScanService.PerformScan(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Scan Started..."})
}

func (h *ScanHandler) GetCompletedScanResult(c *gin.Context) {
	scanID := c.Param("scanID")
	result, err := h.ScanService.GetCompletedScanResult(scanID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
