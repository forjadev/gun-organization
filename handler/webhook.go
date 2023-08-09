package handler

import (
	"log"
	"net/http"

	"github.com/forjadev/gun-organization/service"
	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Github Webhook Membership Receiver
// @Description Manage incoming membership webhook to ensure seamless integration
// @Tags Webhook
// @Accept json
// @Param request body service.GithubWebhookUserManagePayload true "Github webhook payload"
// @Produce json
// @Success 200
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Server Error"
// @Router /webhook [post]
func GitHubWebhookHandler(c *gin.Context) {
	postData := new(service.GithubWebhookUserManagePayload)
	if err := c.ShouldBindJSON(&postData); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}
	if err := service.WebhookHandler(c, postData); err != nil {
		sendError(c, http.StatusInternalServerError, err.Error())
		return
	}

	log.Printf("Webhook data processed: %+v", postData)
}
