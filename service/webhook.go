package service

import (
	"errors"
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

func WebhookHandler(ctx *gin.Context, data *GithubWebhookUserManagePayload) error {
	switch data.Action {
	case "removed":
		if err := deleteUser(data); err != nil {
			return err
		}
	case "added":
		if err := addUser(data); err != nil {
			return err
		}
	default:
		return errors.New("unsupported action")
	}
	return nil
}

func addUser(data *GithubWebhookUserManagePayload) error {
	// TODO: Issue #10 <update database on webhook service>
	fmt.Printf("Usuario adicionado com sucesso, \n %#v\n", data)
	return nil
}

func deleteUser(data *GithubWebhookUserManagePayload) error {
	// TODO: Issue #10 <update database on webhook service>
	fmt.Printf("Usuario deletado com sucesso, \n %#v\n", data)
	return nil
}
