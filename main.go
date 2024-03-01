package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/parnurzeal/gorequest"
)

type ZAPRequest struct {
	URL string `json:"url"`
}

func main() {
	r := gin.Default()

	r.POST("/scan", func(c *gin.Context) {
		// Parse the JSON request body
		var zapRequest ZAPRequest
		if err := c.BindJSON(&zapRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert the struct to JSON
		zapRequestJSON, err := json.Marshal(zapRequest)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
			return
		}

		// Make a POST request to the ZAP API to start the scan
		resp, _, errs := gorequest.New().
			Post("http://localhost:8080/JSON/ascan/action/scan/?apikey=&zapapiformat=JSON&formMethod=POST").
			Send(string(zapRequestJSON)).
			End()

		if len(errs) > 0 {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error making request"})
			return
		}

		// Check if the request was successful
		if resp.StatusCode != http.StatusOK {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error starting scan"})
			return
		}

		// Return success message
		c.JSON(http.StatusOK, gin.H{"message": "Scan started successfully"})
	})

	// Run the server
	if err := r.Run(":8081"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
