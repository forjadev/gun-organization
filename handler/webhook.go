package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type GithubWebhookUserManagePayload struct {
	Action string `json:"action"`
	Scope  string `json:"scope"`

	Member struct {
		ID    uint   `json:"id"`
		Login string `json:"login"`
		Type  string `json:"type"`
		Url   string `json:"html_url"`
	}

	Sender struct {
		ID    uint   `json:"id"`
		Login string `json:"login"`
		Type  string `json:"type"`
	}

	Team struct {
		ID   uint   `json:"id"`
		Name string `json:"name"`
		Slug string `json:"slug"`
	}

	Organization struct {
		ID    uint   `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	}
}

func WebhookReceiverHandle(c *gin.Context) {
	var postData GithubWebhookUserManagePayload
	if err := c.ShouldBindJSON(&postData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": postData})
}
