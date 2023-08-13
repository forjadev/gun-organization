package service

import (
	"fmt"
	"github.com/forjadev/gun-organization/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookService struct {
	// TODO: refactor this to use the `schemas' package
	//Action string `json:"action"`
	//Scope  string `json:"scope"`
	//
	//Member schemas.Member
	//Team   schemas.Team
	//
	//Sender struct {
	//	ID    uint   `json:"id"`
	//	Login string `json:"login"`
	//	Type  string `json:"type"`
	//}
	//
	//Organization struct {
	//	ID    uint   `json:"id"`
	//	Login string `json:"login"`
	//	URL   string `json:"url"`
	//}
}

func NewWebhookService() *WebhookService {
	return &WebhookService{}
}

func (w *WebhookService) GetWebhook(ctx *gin.Context) {
	// TODO: Refactor this to make proper use of `SendError' and `gin.Context'
	//switch data.Action {
	//case "removed":
	//	if err := deleteUser(data); err != nil {
	//		handler.SendError(ctx, http.StatusInternalServerError, "Error while deleting user")
	//	}
	//case "added":
	//	if err := addUser(data); err != nil {
	//		handler.SendError(ctx, http.StatusInternalServerError, "Error while adding user")
	//		return
	//	}
	//default:
	//	handler.SendError(ctx, http.StatusForbidden, "Invalid action")
	//	return
	//}
	handler.SendError(ctx, http.StatusOK, "User managed successfully")
}

func addUser(data *WebhookService) error {
	// TODO: Issue #10 <update database on webhook services>
	fmt.Printf("Usuario adicionado com sucesso, \n %#v\n", data)
	return nil
}

func deleteUser(data *WebhookService) error {
	// TODO: Issue #10 <update database on webhook services>
	fmt.Printf("Usuario deletado com sucesso, \n %#v\n", data)
	return nil
}
