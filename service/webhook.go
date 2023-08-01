package service

import (
	"fmt"

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

func WebhookHandler(ctx *gin.Context, data *GithubWebhookUserManagePayload) {
	switch data.Action {
	case "removed":
		deleteUser(data)
	case "added":
		addUser(data)
	}
}

func addUser(data *GithubWebhookUserManagePayload) {
	// TODO Implements add database
	fmt.Printf("Usuario adicionado com sucesso, \n %#v\n", data)
}

func deleteUser(data *GithubWebhookUserManagePayload) {
	// TODO Implements remove database
	fmt.Printf("Usuario deletado com sucesso, \n %#v\n", data)
}
