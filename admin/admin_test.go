package admin

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	// "example.com/main/student"
	"github.com/gin-gonic/gin"
)

// func TestAddStudentsByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid grNo",
// 			reqbody:      `{"grNo":1399999999,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid section",
// 			reqbody:      `{"grNo":14,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":5,"section":"A1"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"grNo":14,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":15,"section":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid pwd",
// 			reqbody:      `{"grNo":14,"studPwd":"Asd","userRole":"student","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid student name",
// 			reqbody:      `{"grNo":14,"studPwd":"Asdf123@","userRole":"student","studName":"raj6","std":5,"section":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid role",
// 			reqbody:      `{"grNo":14,"studPwd":"Asdf123@","userRole":"admin","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"grNo": 101,"studPwd": "pass123","userRole": "student","studName": "John Doe","std": 10,"section": "A"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":14,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "student already exists",
// 			reqbody:      `{"grNo":1,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/admin/createStud", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			AddStudent(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestEditStudentsByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid grNo",
// 			reqbody:      `{"grNo":1399999999,"studPwd":"Asdf1234"`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid section",
// 			reqbody:      `{"grNo":13,"section":"A1"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"grNo":13,"std":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid pwd",
// 			reqbody:      `{"grNo":13,"studPwd":"Asd"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid student name",
// 			reqbody:      `{"grNo":13,"studName":"raj6"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student doesnot exists",
// 			reqbody:      `{"grNo":10,"studPwd":"Asdf123@"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"grNo": 1,"studPwd": "pas01234"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "old and new value same",
// 			reqbody:      `{"grNo":13"studName":"raj"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":13,"studName":"ram"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPut, "/admin/updateStud", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			EditStud(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestAddSubjectByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid subid",
// 			reqbody:      `{"subId":1399999999,"subName":"math","levelStd":5,"credits":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"subId":1399999999,"subName":"math","levelStd":15,"credits":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid credits",
// 			reqbody:      `{"subId":1399999999,"subName":"math","levelStd":5,"credits":-5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid name",
// 			reqbody:      `{"subId":9999,"subName":"mathffskbdibcisdhcksbckdsbcsdbcisdbcisdicvsdicbidscibdsicbicbidbcidbiddsckbbsvvsuvsvsbuksabdsyuc","levelStd":5,"credits":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "unset subject limit",
// 			reqbody:      `{"subId":9999,"subName":"maths 2","levelStd":12,"credits":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"subId":9999,"subName":"maths 2","levelStd":12,"credits":5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":9999,"subName":"maths 2","levelStd":1,"credits":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/admin/createSub", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			CreateSub(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestEditSubjectsByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid subid",
// 			reqbody:      `{"subId":999999990,"subName":"math 3"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"subId":9999,"levelStd":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid credits",
// 			reqbody:      `{"subId":9999,"credits":-5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid name",
// 			reqbody:      `{"subId":9999,"subName":"mathffskbdibcisdhcksbckdsbcsdbcisdbcisdicvsdicbidscibdsicbicbidbcidbiddsckbbsvvsuvsvsbuksabdsyuc"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject doesnt exist",
// 			reqbody:      `{"subId":1010,"subName":"maths 2"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"subId":9999,"subName":"maths 2","levelStd":12,"credits":5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":9999,"subName":"maths 2"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPut, "/admin/updateSub", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			EditSub(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestDisplaySubjectsByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid syd",
// 			reqbody:      `{"std":999999990}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"std":9}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid cas with result",
// 			reqbody:      `{"std":1}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case without result",
// 			reqbody:      `{"std":9}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodGet, "/admin/displaySub", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			DisplaySubject(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestAddMarksByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid theory marks",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":88,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid practical marks",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":80,"practicalMarks":-15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student not found",
// 			reqbody:      `{"subId":11,"grNo":1313,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not found",
// 			reqbody:      `{"subId":1111,"grNo":13,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "record already present",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusInternalServerError,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/admin/enterMarks", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			EnterMarks(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestEditMarksByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid theory marks",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":-88}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid practical marks",
// 			reqbody:      `{"subId":11,"grNo":13,"practicalMarks":-15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student not found",
// 			reqbody:      `{"subId":11,"grNo":1313,"theoryMarks":70}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not found",
// 			reqbody:      `{"subId":1111,"grNo":13,"practicalMarks":13}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":20}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "record not found to edit",
// 			reqbody:      `{"subId":10,"grNo":1,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPut, "/admin/updateMarks", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			EditMarks(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestDeleteStudentByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid gr no",
// 			reqbody:      `{"grNo":1399999999}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student not found valid gr no",
// 			reqbody:      `{"grNo":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"grNo":13}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":17}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodDelete, "/admin/delStudent", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			DelStud(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestDeleteSubjectByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid subject id",
// 			reqbody:      `{"subId":1199999999}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not found",
// 			reqbody:      `{"subId":111}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"subId":11}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":112}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodDelete, "/admin/delSubject", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			DelSub(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestTeacherPerformanceByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		usrrole      string
// 		reqbody      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "teacher not taking subject",
// 			usrrole:      "admin",
// 			reqbody:      `{"tid":"t2"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "valid",
// 			usrrole:      "admin",
// 			reqbody:      `{"tid":"t1"}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "authorization fail",
// 			usrrole:      "teacherzz",
// 			reqbody:      `{"tid":"t1"}`,
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodGet, "/admin/displayPerformance", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			Performance(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestStudentReportByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Valid",
// 			usrrole:      "admin",
// 			reqbody:      `{"grNo":1}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "valid but no student found",
// 			usrrole:      "admin",
// 			reqbody:      `{"grNo":15}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "valid but no result found",
// 			usrrole:      "admin",
// 			reqbody:      `{"grNo":14}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "invalid grno",
// 			usrrole:      "admin",
// 			reqbody:      `{"grNo":1599999999}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "auth fail",
// 			usrrole:      "student",
// 			reqbody:      `{"grNo":15}`,
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodGet, "/admin/studentreport", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			Report(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestAddTeacherByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid tid",
// 			reqbody:      `{"teacherId":"t9912345678","tPwd":"Asdf1234","role":"teacher","tName":"yash","subId":9999,"stdAllocated":1,"sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid section",
// 			reqbody:      `{"teacherId":"t99","tPwd":"Asdf1234","role":"teacher","tName":"yash","subId":9999,"stdAllocated":1,"sectionAllocated":"A1"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"teacherId":"t99","tPwd":"Asdf1234","role":"teacher","tName":"yash","subId":9999,"stdAllocated":15,"sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid pwd",
// 			reqbody:      `{"teacherId":"t100","tPwd":"1232","role":"teacher","tName":"het"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid teacher name",
// 			reqbody:      `{"teacherId":"t100","tPwd":"Asdf1232","role":"teacher","tName":"het44"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid role",
// 			reqbody:      `{"teacherId":"t100","tPwd":"Asdf1232","role":"student","tName":"het"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"teacherId":"t100","tPwd":"Asdf1232","role":"teacher","tName":"het"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case without subject",
// 			reqbody:      `{"teacherId":"t100","tPwd":"Asdf1232","role":"teacher","tName":"het"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case with subject",
// 			reqbody:      `{"teacherId":"t99","tPwd":"Asdf1234","role":"teacher","tName":"yash","subId":9999,"stdAllocated":1,"sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "teacher already exists",
// 			reqbody:      `{"teacherId":"t99","tPwd":"Asdf123@","role":"teacher","tName":"raj"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "missing section",
// 			reqbody:      `{"teacherId":"t99","tPwd":"Asdf123@","role":"teacher","tName":"raj","stdAllocated":1}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "missing std",
// 			reqbody:      `{"teacherId":"t99","tPwd":"Asdf123@","role":"teacher","tName":"raj","sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "sub provided but missing class",
// 			reqbody:      `{"teacherId":"t99","tPwd":"Asdf123@","role":"teacher","tName":"raj","subId":9999}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/admin/createStud", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			AddTeacher(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestEditTeacherByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid section",
// 			reqbody:      `{"teacherId":"t99""sectionAllocated":"A1"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"teacherId":"t99","stdAllocated":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid pwd",
// 			reqbody:      `{"teacherId":"t100","tPwd":"1232"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid teacher name",
// 			reqbody:      `{"teacherId":"t100","tName":"het44"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"teacherId":"t100","tPwd":"Asdf1231"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "only editing std for that who hasnt allocated class",
// 			reqbody:      `{"teacherId":"t2","tPwd":"Asdf1232","role":"teacher","tName":"het","stdAllocated":1}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "only editing section for that who hasnt allocated class",
// 			reqbody:      `{"teacherId":"t2","tPwd":"Asdf1234","role":"teacher","tName":"yash","subId":9999,"sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "only editing subject for that who hasnt allocated class",
// 			reqbody:      `{"teacherId":"t2","tPwd":"Asdf1234","role":"teacher","tName":"yash","subId":1}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "valid case",
// 			reqbody:      `{"teacherId":"t2","tName":"sallu"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/admin/createStud", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			EditTeacher(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestDeleteTeacherByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid teacher id",
// 			reqbody:      `{"teacherId":"1399999999"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "teacher not found",
// 			reqbody:      `{"teacherId":"15"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"teacherId":13}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"teacherId":"t2"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodDelete, "/admin/delTeacher", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			DeleteTeacher(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestSetSubLimitByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"std":139,"limit":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid limit",
// 			reqbody:      `{"std":11,"limit":-5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"std":12,"limit":5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"std":11,"limit":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/admin/setSubLimit", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			SetSubLimit(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestPendingRequestByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "authorization fail",
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case (auth pass)",
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodGet, "/admin/pendingRequest", bytes.NewBufferString(""))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			ShowPendingReq(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestAcceptPendingRequestByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		usrrole      string
// 		reqbody      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "authorization fail",
// 			usrrole:      "teacher",
// 			reqbody:      `{"uName":"ta","uPwd":"password","uRole":"admin","Uid":"admin123","std":0,"section":"A","subId":9999}`,
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "invalid student id",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ts","uPwd":"password","uRole":"student","Uid":121111111112,"std":5,"section":"A"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher/admin id",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"tt","uPwd":"password","uRole":"teacher","Uid":"t100000000","std":5,"section":"A","subId":9999}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid student pwd",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ts","uPwd":"pasrd","uRole":"student","Uid":1212,"std":5,"section":"A"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher/admin pwd",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ttt","uPwd":"pass","uRole":"teacher","Uid":"t10","std":5,"section":"A","subId":9999}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid student section",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ts","uPwd":"password","uRole":"student","Uid":1212,"std":5,"section":"A1"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid student std",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ts","uPwd":"password","uRole":"student","Uid":1212,"std":15,"section":"A"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher class allocation section",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"tt","uPwd":"password","uRole":"teacher","Uid":"t10","std":5,"section":"A1","subId":9999}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher class allocation std",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"tt","uPwd":"password","uRole":"teacher","Uid":"t10","std":15,"section":"A","subId":9999}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher class allocation sub",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ttt","uPwd":"password","uRole":"teacher","Uid":"t10","subId":9999}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Valid case student",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ts","uPwd":"password","uRole":"student","Uid":12121,"std":5,"section":"A"}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case teacher with subject and class",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"tt","uPwd":"password","uRole":"teacher","Uid":"t101","std":5,"section":"A","subId":9999}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case teacher without subject and class",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ttt","uPwd":"password","uRole":"teacher","Uid":"t123"}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case admin",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ta","uPwd":"password","uRole":"admin","Uid":"a123"}`,
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/admin/acceptRequest", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			AcceptPendingReq(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

func TestRejectPendingRequestByAdmin(t *testing.T) {

	testcases := []struct {
		name         string
		usrrole      string
		reqbody      string
		expectedCode int
	}{
		{
			name:         "authorization fail",
			usrrole:      "teacher",
			reqbody:      `{"uName":"temp stud","uPwd":"password","uRole":"student"}`,
			expectedCode: http.StatusUnauthorized,
		},
		{
			name:         "Valid case",
			usrrole:      "admin",
			reqbody:      `{"uName":"temp stud","uPwd":"password","uRole":"student"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "inValid case, request not found",
			usrrole:      "admin",
			reqbody:      `{"uName":"ts","uPwd":"password","uRole":"student"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, err := http.NewRequest(http.MethodDelete, "/admin/rejectRequest", bytes.NewBufferString(tc.reqbody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req
			ctx.Set("userrole", tc.usrrole)
			RejectRequest(ctx)
			if w.Code != tc.expectedCode {
				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
			}
			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
		})
	}
}
