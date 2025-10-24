package student

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDisplayStudents(t *testing.T) {
	gin.SetMode(gin.TestMode)
	w := httptest.NewRecorder()
	ctx, _ := gin.CreateTestContext(w)

	jsonBody := `{
		"viewByStd":10,
		"viewBySection":"A",
		"MinPercent":0,
		"MaxPercent":100
	}`
	req, err := http.NewRequest(http.MethodPost, "/admin/display", bytes.NewBufferString(jsonBody))
	if err != nil {
		t.Fatalf("failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	ctx.Request = req

	ctx.Set("userrole", "admin")

	DisplayStudents(ctx)

	fmt.Println("Status Code:", w.Code)
	fmt.Println("Response Body:", w.Body.String())

	if w.Code != http.StatusOK {
		t.Fatalf("expected status 200, got %d", w.Code)
	}
}
