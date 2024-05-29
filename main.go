package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gopheramol/document-handler/bootstrap"
	"github.com/gopheramol/document-handler/configuration"
	"github.com/gopheramol/document-handler/model"
)

func main() {

	// Load configuration
	config := configuration.LoadConfigs()

	// Initialize webhook server
	go startWebhookServer()

	// Initialize routes
	router := bootstrap.InitializeRoutes(config)

	// Run the HTTP server
	if err := router.Run(":" + config.Port); err != nil {
		log.Fatalf("failed to start HTTP server: %v", err)
	}
}

func startWebhookServer() {
	// Create a new Gin router for webhook server
	webhookRouter := gin.Default()

	// Define a route to handle SNS webhook notifications
	webhookRouter.POST("/webhook/sns", handleSNSWebhook)

	// Run the webhook server on port 8081
	if err := webhookRouter.Run(":8081"); err != nil {
		log.Fatalf("failed to start webhook server: %v", err)
	}
}

// HandleSNSWebhook handles incoming SNS webhook notifications
func handleSNSWebhook(c *gin.Context) {
	log.Default().Println("Webhook called")

	// Read the request body
	byteData, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Printf("failed to read request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to read request body"})
		return
	}
	log.Default().Println(string(byteData))

	// Unmarshal the JSON payload into the SNSNotification struct
	var snsNotification model.SNSNotification
	if err := json.Unmarshal(byteData, &snsNotification); err != nil {
		log.Printf("failed to unmarshal JSON request body: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "failed to unmarshal JSON"})
		return
	}

	// Process the SNS notification payload
	log.Printf("Received SNS notification: %+v", snsNotification)

	// Respond with a success message
	c.JSON(http.StatusOK, gin.H{"Response": "SNS notification processed successfully"})
}
