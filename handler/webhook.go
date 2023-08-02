package handler

import (
	"log"
	"net/http"

	"github.com/forjadev/gun-organization/service"
	"github.com/gin-gonic/gin"
)

func WebhookReceiverHandle(c *gin.Context) {
	postData := new(service.GithubWebhookUserManagePayload)
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := service.WebhookHandler(c, postData); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	log.Printf("Webhook data processed: %+v", postData)
	c.JSON(http.StatusOK, gin.H{"data": postData})
}
