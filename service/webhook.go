package service

import (
	"errors"
	"fmt"
	"log"

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
			log.Printf("Error while deleting user: %+v", data)
			return err
		}
	case "added":
		if err := addUser(data); err != nil {
			log.Printf("Error while adding user: %+v", data)
			return err
		}
	default:
		log.Printf("unsupported action: %+v", data.Action)
		return errors.New("unsupported action")
	}
	// TODO: Issue #10 <update database on webhook service> Trigger Github Action to update README.md
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
