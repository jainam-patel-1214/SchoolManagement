package student

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDisplayStudents(t *testing.T) {

	testcases := []struct {
		name         string
		reqbody      string
		usrrole      string
		expectedCode int
	}{
		{name: "Valid case", reqbody: `{"viewByStd":5,"viewBySection":"A"}`, usrrole: "student", expectedCode: http.StatusOK},
		{name: "Valid case", reqbody: `{"viewByStd":5,"viewBySection":"A","minPercent":20,"maxPercent":95}`, usrrole: "student", expectedCode: http.StatusOK},
		{
			name:         "Invalid std",
			reqbody:      `{"viewByStd":13,"viewBySection":"A","minPercent":0,"maxPercent":100}`,
			usrrole:      "student",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid section",
			reqbody:      `{"viewByStd":10,"viewBySection":"A2","minPercent":0,"maxPercent":100}`,
			usrrole:      "student",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid min",
			reqbody:      `{"viewByStd":13,"viewBySection":"A","minPercent":-10,"maxPercent":100}`,
			usrrole:      "student",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid combo",
			reqbody:      `{"viewByStd":13,"viewBySection":"A","minPercent":100,"maxPercent":80}`,
			usrrole:      "student",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid role",
			reqbody:      `{"viewByStd":12,"viewBySection":"A","minPercent":0,"maxPercent":80}`,
			usrrole:      "studentz",
			expectedCode: http.StatusUnauthorized,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, err := http.NewRequest(http.MethodGet, "/student/display", bytes.NewBufferString(tc.reqbody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req
			ctx.Set("userrole", tc.usrrole)
			DisplayStudents(ctx)
			if w.Code != tc.expectedCode {
				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
			}
			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
		})
	}
}

func TestDisplaySubject(t *testing.T) {
	testcases := []struct {
		name         string
		reqbody      string
		usrrole      string
		expectedCode int
	}{
		{
			name:         "Valid",
			reqbody:      `{"std":1}`,
			usrrole:      "student",
			expectedCode: http.StatusOK,
		},
		{
			name:         "valid but no subject found",
			reqbody:      `{"std":12}`,
			usrrole:      "student",
			expectedCode: http.StatusOK,
		},
		{
			name:         "Invalid std param",
			reqbody:      `{"std":13}`,
			usrrole:      "student",
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "Invalid role",
			reqbody:      `{"std":12}`,
			usrrole:      "studentz",
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "std 0",
			reqbody:      `{"std":0}`,
			usrrole:      "student",
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, err := http.NewRequest(http.MethodGet, "/students/displaySub", bytes.NewBufferString(tc.reqbody))
			if err != nil {
				t.Fatalf("error occured: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req
			ctx.Set("userrole", tc.usrrole)
			DisplaySubject(ctx)
			if w.Code != tc.expectedCode {
				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
			}
			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
		})
	}
}

func TestDisplayStudentReport(t *testing.T) {
	testcases := []struct {
		name         string
		usrrole      string
		usrid        int
		expectedCode int
	}{
		{
			name:         "Valid",
			usrrole:      "student",
			usrid:        1,
			expectedCode: http.StatusOK,
		},
		{
			name:         "valid but no result found",
			usrrole:      "student",
			usrid:        12121,
			expectedCode: http.StatusOK,
		},
		{
			name:         "id not found",
			usrrole:      "student",
			expectedCode: http.StatusUnauthorized,
		},
	}
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, err := http.NewRequest(http.MethodGet, "/students/displaySub", bytes.NewBufferString(""))
			if err != nil {
				t.Fatalf("error occured: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req
			ctx.Set("userrole", tc.usrrole)
			ctx.Set("UiD", tc.usrid)
			Report(ctx)
			if w.Code != tc.expectedCode {
				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
			}
			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
		})
	}
}
