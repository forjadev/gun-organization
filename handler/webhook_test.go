package handler_test

import (
	"github.com/forjadev/gun-organization/handler"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestWebhookReceiverHandle_Success(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/webhook", handler.WebhookReceiverHandle)
	reqPayload := `{
		"action": "some_action",
		"scope": "some_scope",
		"member": {
			"id": 123,
			"login": "john_doe",
			"type": "user",
			"url": "https://example.com/user/john_doe"
		},
		"sender": {
			"id": 456,
			"login": "sender_user",
			"type": "user"
		},
		"team": {
			"id": 789,
			"name": "some_team",
			"slug": "some_team_slug"
		},
		"organization": {
			"id": 987,
			"login": "some_org",
			"url": "https://example.com/org/some_org"
		}
	}`
	rw := performRequest(r, http.MethodPost, "/api/v1/webhook", reqPayload)
	assertResponseCode(t, rw.Code, http.StatusOK)
}

func TestWebhookReceiverHandle_InvalidJSON(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.POST("/api/v1/webhook", handler.WebhookReceiverHandle)
	reqPayload := `{ invalid json payload`
	rw := performRequest(r, http.MethodPost, "/api/v1/webhook", reqPayload)

	assertResponseCode(t, rw.Code, http.StatusBadRequest)
}

func assertResponseCode(t *testing.T, got, want int) {
	t.Helper()
	assert.Equal(t, want, got, "response code is wrong")
}

func performRequest(r *gin.Engine, method, path, payload string) *httptest.ResponseRecorder {
	rw := httptest.NewRecorder()
	req, _ := http.NewRequest(method, path, strings.NewReader(payload))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(rw, req)
	return rw
}
