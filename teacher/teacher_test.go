package teacher

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

// func TestAddStudentsByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid grNo",
// 			reqbody:      `{"grNo":1399999999,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid section",
// 			reqbody:      `{"grNo":13,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":5,"section":"A1"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"grNo":13,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":15,"section":"A"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid pwd",
// 			reqbody:      `{"grNo":13,"studPwd":"Asd","userRole":"student","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid student name",
// 			reqbody:      `{"grNo":13,"studPwd":"Asdf123@","userRole":"student","studName":"raj6","std":5,"section":"A"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid role",
// 			reqbody:      `{"grNo":13,"studPwd":"Asdf123@","userRole":"admin","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student already exists",
// 			reqbody:      `{"grNo":1,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"grNo": 101,"studPwd": "pass123","userRole": "student","studName": "John Doe","std": 10,"section": "A"}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":13,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":5,"section":"A"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/teacher/createStud", bytes.NewBufferString(tc.reqbody))
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

// func TestEditStudentsByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid grNo",
// 			reqbody:      `{"grNo":1399999999,"studPwd":"Asdf1234"`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid section",
// 			reqbody:      `{"grNo":13,"section":"A1"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"grNo":13,"std":15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid pwd",
// 			reqbody:      `{"grNo":13,"studPwd":"Asd"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid student name",
// 			reqbody:      `{"grNo":13,"studName":"raj6"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student doesnot exists",
// 			reqbody:      `{"grNo":10,"studPwd":"Asdf123@"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"grNo": 1,"studPwd": "pas01234"}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "old and new value same",
// 			reqbody:      `{"grNo":13"studName":"raj"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":13,"studName":"ram"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPut, "/teacher/updateStud", bytes.NewBufferString(tc.reqbody))
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

// func TestAddSubjectByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid subid",
// 			reqbody:      `{"subId":1399999999,"subName":"math","levelStd":5,"credits":5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"subId":1399999999,"subName":"math","levelStd":15,"credits":5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid credits",
// 			reqbody:      `{"subId":1399999999,"subName":"math","levelStd":5,"credits":-5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid name",
// 			reqbody:      `{"subId":9999,"subName":"mathffskbdibcisdhcksbckdsbcsdbcisdbcisdicvsdicbidscibdsicbicbidbcidbiddsckbbsvvsuvsvsbuksabdsyuc","levelStd":5,"credits":5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "unset subject limit",
// 			reqbody:      `{"subId":9999,"subName":"maths 2","levelStd":12,"credits":5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"subId":9999,"subName":"maths 2","levelStd":12,"credits":5}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":9999,"subName":"maths 2","levelStd":1,"credits":5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/teacher/createSub", bytes.NewBufferString(tc.reqbody))
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

// func TestEditSubjectsByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid subid",
// 			reqbody:      `{"subId":999999990,"subName":"math 3"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"subId":9999,"levelStd":15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid credits",
// 			reqbody:      `{"subId":9999,"credits":-5}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid name",
// 			reqbody:      `{"subId":9999,"subName":"mathffskbdibcisdhcksbckdsbcsdbcisdbcisdicvsdicbidscibdsicbicbidbcidbiddsckbbsvvsuvsvsbuksabdsyuc"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject doesnt exist",
// 			reqbody:      `{"subId":1010,"subName":"maths 2"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"subId":9999,"subName":"maths 2","levelStd":12,"credits":5}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":9999,"subName":"maths 3"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPut, "/teacher/updateSub", bytes.NewBufferString(tc.reqbody))
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

// func TestDisplaySubjectsByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid syd",
// 			reqbody:      `{"std":999999990}`,
// 			usrrole:      "teacher",
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
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case without result",
// 			reqbody:      `{"std":9}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodGet, "/teacher/displaySub", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			student.DisplaySubject(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestAddMarksByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid theory marks",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":88,"practicalMarks":15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid practical marks",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":80,"practicalMarks":-15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student not found",
// 			reqbody:      `{"subId":11,"grNo":1313,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not found",
// 			reqbody:      `{"subId":1111,"grNo":13,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "teacher",
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
// 			reqbody:      `{"subId":10,"grNo":13,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "record already present",
// 			reqbody:      `{"subId":10,"grNo":13,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusInternalServerError,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/teacher/enterMarks", bytes.NewBufferString(tc.reqbody))
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

// func TestEditMarksByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid theory marks",
// 			reqbody:      `{"subId":11,"grNo":13,"theoryMarks":-88}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid practical marks",
// 			reqbody:      `{"subId":11,"grNo":13,"practicalMarks":-15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student not found",
// 			reqbody:      `{"subId":10,"grNo":1313,"theoryMarks":70}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not found",
// 			reqbody:      `{"subId":1111,"grNo":13,"practicalMarks":13}`,
// 			usrrole:      "teacher",
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
// 			reqbody:      `{"subId":10,"grNo":13,"theoryMarks":50}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "record not found to edit",
// 			reqbody:      `{"subId":10,"grNo":1,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusInternalServerError,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPut, "/teacher/updateMarks", bytes.NewBufferString(tc.reqbody))
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

// func TestDeleteStudentByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid gr no",
// 			reqbody:      `{"grNo":1399999999}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student not found valid gr no",
// 			reqbody:      `{"grNo":15}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"grNo":13}`,
// 			usrrole:      "teacherzz",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":13}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodDelete, "/teacher/delStudent", bytes.NewBufferString(tc.reqbody))
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

// func TestDeleteSubjectByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid subject id",
// 			reqbody:      `{"subId":1199999999}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not found",
// 			reqbody:      `{"subId":111}`,
// 			usrrole:      "teacher",
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
// 			reqbody:      `{"subId":11}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusOK,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodDelete, "/teacher/delSubject", bytes.NewBufferString(tc.reqbody))
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

// func TestAddReviewByTeacher(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		usrid        string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "student invalid grno",
// 			reqbody:      `{"grNo":1399999999,"comment":"sincere"}`,
// 			usrrole:      "teacher",
// 			usrid:        "t1",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "valid but student not found",
// 			reqbody:      `{"grNo":19,"comment":"sincere"}`,
// 			usrrole:      "teacher",
// 			usrid:        "t1",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"grNo":13,"comment":"sincere"}`,
// 			usrrole:      "teacherzz",
// 			usrid:        "t1",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":13,"comment":"sincere"}`,
// 			usrrole:      "teacher",
// 			usrid:        "t1",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case but review already present",
// 			reqbody:      `{"grNo":13,"comment":"sincere"}`,
// 			usrrole:      "teacher",
// 			usrid:        "t1",
// 			expectedCode: http.StatusInternalServerError,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodPost, "/teacher/addReview", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			ctx.Set("UiD", tc.usrid)
// 			AddReviews(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

// func TestTeacherPerformance(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		usrrole      string
// 		usrid        string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "teacher not taking subject",
// 			usrrole:      "teacher",
// 			usrid:        "t2",
// 			expectedCode: http.StatusNotFound,
// 		},
// 		{
// 			name:         "valid",
// 			usrrole:      "teacher",
// 			usrid:        "t1",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "authorization fail",
// 			usrrole:      "teacherzz",
// 			usrid:        "t1",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodGet, "/teacher/displayPerformance", bytes.NewBufferString(""))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			ctx.Set("UiD", tc.usrid)
// 			Performance(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

func TestStudentReportByTeacher(t *testing.T) {

	testcases := []struct {
		name         string
		reqbody      string
		usrrole      string
		expectedCode int
	}{
		{
			name:         "Valid",
			usrrole:      "teacher",
			reqbody:      `{"grNo":1}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "valid but no student found",
			usrrole:      "teacher",
			reqbody:      `{"grNo":15}`,
			expectedCode: http.StatusInternalServerError,
		},
		{
			name:         "valid but no result found",
			usrrole:      "teacher",
			reqbody:      `{"grNo":13}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid grno",
			usrrole:      "teacher",
			reqbody:      `{"grNo":1599999999}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, err := http.NewRequest(http.MethodGet, "/teacher/studentreport", bytes.NewBufferString(tc.reqbody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req
			ctx.Set("userrole", tc.usrrole)
			Report(ctx)
			if w.Code != tc.expectedCode {
				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
			}
			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
		})
	}
}
