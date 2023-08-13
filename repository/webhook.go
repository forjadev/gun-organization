package repository

import (
	"github.com/forjadev/gun-organization/handler"
	"github.com/forjadev/gun-organization/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @BasePath /api/v1

// @Summary Github Webhook Membership Receiver
// @Description Manage incoming membership webhook to ensure seamless integration
// @Tags Webhook
// @Accept json
// @Param request body services.GithubWebhookUserManagePayload true "Github webhook payload"
// @Produce json
// @Success 200
// @Failure 400 {object} ErrorResponse "Bad Request"
// @Failure 500 {object} ErrorResponse "Server Error"
// @Router /webhook [post]
func GitHubWebhookHandler(ctx *gin.Context) {
	wh := service.NewWebhookService()
	if err := ctx.ShouldBindJSON(wh); err != nil {
		handler.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	wh.GetWebhook(ctx)
}
