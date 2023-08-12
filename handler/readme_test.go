package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/forjaDev/gun-organization/config"
	"github.com/stretchr/testify/assert"

	"github.com/gin-gonic/gin"
)

func _setupTestRouter() (route *gin.Engine) {
	route = gin.Default()
	route.GET("/readme", func(c *gin.Context) {
		c.String(200, "Readme Content")
	})
	return
}

func TestReadmeRoute(t *testing.T) {
    router := _setupTestRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/readme", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, http.StatusOK, w.Code)
    assert.Equal(t, "Readme Content", w.Body.String())
}

func _recordStats(t *testing.T, productID int64) (err error) {
	tx, err := config.InitDatabase()
	if err != nil {
		t.Fatalf("initiating database: %s", err)
	}

	defer func() {
		switch err {
		case nil:
			err = tx.Commit().Error
		default:
			tx.Rollback()
		}
	}()

	// tx.Exec("UPDATE products SET views = views + 1")
	// tx.Exec("INSERT INTO product_viewers (user_id, product_id) VALUES (?, ?)", userID, productID)
	// TODO: Populate sql mock tables with /teams and webhook info by using pgxmock
	return
}
