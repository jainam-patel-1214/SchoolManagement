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
// 			reqbody:      `{"grNo":14,"studPwd":"Asdf123@","userRole":"student","studName":"raj","std":8,"section":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":15,"studPwd":"Asdf123@","userRole":"student","studName":"raju","std":8,"section":"B"}`,
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
// 			reqbody:      `{"grNo":14,"section":"A1"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"grNo":14,"std":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid pwd",
// 			reqbody:      `{"grNo":14,"studPwd":"Asd"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid student name",
// 			reqbody:      `{"grNo":14,"studName":"raj6"}`,
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
// 			reqbody:      `{"grNo":14"studName":"raj"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"grNo":14,"studName":"ram"}`,
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
// 			reqbody:      `{"subId":100,"subName":"english","levelStd":15,"credits":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid credits",
// 			reqbody:      `{"subId":100,"subName":"english","levelStd":5,"credits":-5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid name",
// 			reqbody:      `{"subId":100,"subName":"mathffskbdibcisdhcksbckdsbcsdbcisdbcisdicvsdicbidscibdsicbicbidbcidbiddsckbbsvvsuvsvsbuksabdsyuc","levelStd":5,"credits":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "unset subject limit",
// 			reqbody:      `{"subId":100,"subName":"eng 2","levelStd":12,"credits":5}`,
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
// 			reqbody:      `{"subId":100,"subName":"english","levelStd":11,"credits":5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":101,"subName":"french","levelStd":11,"credits":5}`,
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
// 			reqbody:      `{"subId":100,"levelStd":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid credits",
// 			reqbody:      `{"subId":100,"credits":-5}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid name",
// 			reqbody:      `{"subId":100,"subName":"mathffskbdibcisdhcksbckdsbcsdbcisdbcisdicvsdicbidscibdsicbicbidbcidbiddsckbbsvvsuvsvsbuksabdsyuc"}`,
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
// 			reqbody:      `{"subId":100,"subName":"german"}`,
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
// 			reqbody:      `{"teacherId":"T1","tPwd":"Asdf1234","role":"teacher","tName":"yash","subId":9999,"stdAllocated":1,"sectionAllocated":"A1"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"teacherId":"T1","tPwd":"Asdf1234","role":"teacher","tName":"yash","subId":9999,"stdAllocated":15,"sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid pwd",
// 			reqbody:      `{"teacherId":"T1","tPwd":"1232","role":"teacher","tName":"het"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid teacher name",
// 			reqbody:      `{"teacherId":"T1","tPwd":"Asdf1232","role":"teacher","tName":"het44"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid role",
// 			reqbody:      `{"teacherId":"T1","tPwd":"Asdf1232","role":"student","tName":"het"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "authorization fail",
// 			reqbody:      `{"teacherId":"T1","tPwd":"Asdf1232","role":"teacher","tName":"het"}`,
// 			usrrole:      "teacher",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case without subject",
// 			reqbody:      `{"teacherId":"T1","tPwd":"Asdf1232","role":"teacher","tName":"enna"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case with subject",
// 			reqbody:      `{"teacherId":"T2","tPwd":"Asdf1234","role":"teacher","tName":"meena","subId":100,"stdAllocated":11,"sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case with subject",
// 			reqbody:      `{"teacherId":"T3","tPwd":"Asdf1234","role":"teacher","tName":"deeka","subId":101,"stdAllocated":11,"sectionAllocated":"B"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "too long name",
// 			reqbody:      `{"teacherId":"T6","tPwd":"Asdf1234","role":"teacher","tName":"enna meena deeka blaaaaaaaaaaa blaaaaaaaaaaa blaaaaaaaaaaa blaaaaaaaaaaa","subId":101,"stdAllocated":11,"sectionAllocated":"B"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not exist",
// 			reqbody:      `{"teacherId":"T7","tPwd":"Asdf1234","role":"teacher","tName":"deeka","subId":10100,"stdAllocated":11,"sectionAllocated":"B"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "teacher already exists",
// 			reqbody:      `{"teacherId":"t99","tPwd":"Asdf123@","role":"teacher","tName":"raj"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "missing section",
// 			reqbody:      `{"teacherId":"t5","tPwd":"Asdf123@","role":"teacher","tName":"raj","stdAllocated":1}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "missing std",
// 			reqbody:      `{"teacherId":"t5","tPwd":"Asdf123@","role":"teacher","tName":"raj","sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "sub provided but missing class",
// 			reqbody:      `{"teacherId":"t5","tPwd":"Asdf123@","role":"teacher","tName":"raj","subId":100}`,
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
// 			reqbody:      `{"teacherId":"t99","sectionAllocated":"A1"}`,
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
// 			name:         "Invalid teacher name long",
// 			reqbody:      `{"teacherId":"t100","tName":"het wjqfujbacbdvjvjvjwjqfujbacbdvjvjvjwjqfujbacbdvjvjvjwjqfujbacbdvjvjvjwjq"}`,
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
// 			reqbody:      `{"teacherId":"t2","stdAllocated":1}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "only editing section for that who hasnt allocated class",
// 			reqbody:      `{"teacherId":"t2","sectionAllocated":"A"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "only editing subject for that who hasnt allocated class",
// 			reqbody:      `{"teacherId":"t2","subId":1}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "valid case",
// 			reqbody:      `{"teacherId":"t2","tName":"god"}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "valid case",
// 			reqbody:      `{"teacherId":"t2","tName":"prabhu","subId":101,"stdAllocated":11,"sectionAllocated":"B"}`,
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

// func TestDisplaySubjectsByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "Invalid syd",
// 			reqbody:      `{"std":99}`,
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
// 			reqbody:      `{"std":11}`,
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
// 			reqbody:      `{"subId":100,"grNo":13,"theoryMarks":88,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid practical marks",
// 			reqbody:      `{"subId":100,"grNo":13,"theoryMarks":80,"practicalMarks":-15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student not found",
// 			reqbody:      `{"subId":100,"grNo":131311,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not found",
// 			reqbody:      `{"subId":1001,"grNo":1,"theoryMarks":80,"practicalMarks":15}`,
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
// 			reqbody:      `{"subId":100,"grNo":14,"theoryMarks":80,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":100,"grNo":15,"theoryMarks":70,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":101,"grNo":14,"theoryMarks":60,"practicalMarks":20}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case",
// 			reqbody:      `{"subId":101,"grNo":15,"theoryMarks":80,"practicalMarks":20}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "record already present",
// 			reqbody:      `{"subId":1,"grNo":1,"theoryMarks":80,"practicalMarks":15}`,
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
// 			reqbody:      `{"subId":1,"grNo":1,"theoryMarks":-88}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid practical marks",
// 			reqbody:      `{"subId":1,"grNo":1,"practicalMarks":-15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student not found",
// 			reqbody:      `{"subId":1,"grNo":131311,"theoryMarks":70}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "subject not found",
// 			reqbody:      `{"subId":11110,"grNo":1,"practicalMarks":13}`,
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
// 			reqbody:      `{"subId":1,"grNo":1,"theoryMarks":50,"practicalMarks":15}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "record not found to edit",
// 			reqbody:      `{"subId":100,"grNo":1,"theoryMarks":80,"practicalMarks":15}`,
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
// 			reqbody:      `{"tid":"t123"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "valid",
// 			usrrole:      "admin",
// 			reqbody:      `{"tid":"T2"}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "teacher not found",
// 			usrrole:      "admin",
// 			reqbody:      `{"tid":"T254"}`,
// 			expectedCode: http.StatusBadRequest,
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

// func TestDisplayStudentsByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		reqbody      string
// 		usrrole      string
// 		expectedCode int
// 	}{
// 		{name: "Valid case", reqbody: `{"viewByStd":8}`, usrrole: "admin", expectedCode: http.StatusOK},
// 		{name: "Valid case", reqbody: `{"viewByStd":8,"viewBySection":"A","minPercent":20,"maxPercent":95}`, usrrole: "admin", expectedCode: http.StatusOK},
// 		{
// 			name:         "Invalid std",
// 			reqbody:      `{"viewByStd":13,"viewBySection":"A","minPercent":0,"maxPercent":100}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid section",
// 			reqbody:      `{"viewByStd":10,"viewBySection":"A2","minPercent":0,"maxPercent":100}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid min",
// 			reqbody:      `{"viewByStd":13,"viewBySection":"A","minPercent":-10,"maxPercent":100}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid combo",
// 			reqbody:      `{"viewByStd":13,"viewBySection":"A","minPercent":100,"maxPercent":80}`,
// 			usrrole:      "admin",
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Invalid role",
// 			reqbody:      `{"viewByStd":12,"viewBySection":"A","minPercent":0,"maxPercent":80}`,
// 			usrrole:      "student",
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodGet, "/student/display", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			DisplayStudents(ctx)
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
// 			reqbody:      `{"grNo":14}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "valid but no student found",
// 			usrrole:      "admin",
// 			reqbody:      `{"grNo":151}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "valid but no result found",
// 			usrrole:      "admin",
// 			reqbody:      `{"grNo":12121}`,
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
// 			reqbody:      `{"grNo":151}`,
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
// 			reqbody:      `{"grNo":14}`,
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
// 			reqbody:      `{"subId":1112}`,
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
// 			reqbody:      `{"subId":100}`,
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
// 			reqbody:      `{"std":3,"limit":-5}`,
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
// 			reqbody:      `{"std":3,"limit":5}`,
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
// 			reqbody:      `{"uName":"sanjay","uPwd":"password","uRole":"student","Uid":121111111112,"std":5,"section":"A"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid student pwd",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"sanjay","uPwd":"pasrd","uRole":"student","Uid":1001,"std":5,"section":"A"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid student section",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"sanjay","uPwd":"password","uRole":"student","Uid":1001,"std":5,"section":"A1"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid student std",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"sanjay","uPwd":"password","uRole":"student","Uid":1001,"std":15,"section":"A"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher/admin id",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"maya","uPwd":"password","uRole":"teacher","Uid":"t100000000","std":5,"section":"A","subId":9999}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher/admin pwd",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"maya","uPwd":"pass","uRole":"teacher","Uid":"t10","std":5,"section":"A","subId":101}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher class allocation section",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"maya","uPwd":"password","uRole":"teacher","Uid":"t10","std":5,"section":"A1","subId":101}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher class allocation std",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"maya","uPwd":"password","uRole":"teacher","Uid":"t10","std":15,"section":"A","subId":101}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "invalid teacher class allocation sub",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"maya","uPwd":"password","uRole":"teacher","Uid":"t10","subId":101}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "student already exist",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"sanjay","uPwd":"password","uRole":"student","Uid":1,"std":5,"section":"A"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "teacher already exist",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"saya","uPwd":"password","uRole":"teacher","Uid":"t2"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "admin already exist",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"popat","uPwd":"password","uRole":"admin","Uid":"a123"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 		{
// 			name:         "Valid case student",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"sanjay","uPwd":"password","uRole":"student","Uid":12235,"std":5,"section":"A"}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case teacher with subject and class",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"maya","uPwd":"password","uRole":"teacher","Uid":"T11","std":5,"section":"A","subId":101}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case teacher without subject and class",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"saya","uPwd":"password","uRole":"teacher","Uid":"T12"}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "Valid case admin",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"popat","uPwd":"password","uRole":"admin","Uid":"A1"}`,
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

// func TestRejectPendingRequestByAdmin(t *testing.T) {

// 	testcases := []struct {
// 		name         string
// 		usrrole      string
// 		reqbody      string
// 		expectedCode int
// 	}{
// 		{
// 			name:         "authorization fail",
// 			usrrole:      "teacher",
// 			reqbody:      `{"uName":"temp stud","uPwd":"password","uRole":"student"}`,
// 			expectedCode: http.StatusUnauthorized,
// 		},
// 		{
// 			name:         "Valid case",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"temp","uPwd":"password","uRole":"admin"}`,
// 			expectedCode: http.StatusOK,
// 		},
// 		{
// 			name:         "inValid case, request not found",
// 			usrrole:      "admin",
// 			reqbody:      `{"uName":"ts","uPwd":"password","uRole":"student"}`,
// 			expectedCode: http.StatusBadRequest,
// 		},
// 	}

// 	for _, tc := range testcases {
// 		t.Run(tc.name, func(t *testing.T) {
// 			w := httptest.NewRecorder()
// 			ctx, _ := gin.CreateTestContext(w)
// 			req, err := http.NewRequest(http.MethodDelete, "/admin/rejectRequest", bytes.NewBufferString(tc.reqbody))
// 			if err != nil {
// 				t.Fatalf("failed to create request: %v", err)
// 			}
// 			req.Header.Set("Content-Type", "application/json")
// 			ctx.Request = req
// 			ctx.Set("userrole", tc.usrrole)
// 			RejectRequest(ctx)
// 			if w.Code != tc.expectedCode {
// 				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
// 			}
// 			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
// 		})
// 	}
// }

func TestRegister(t *testing.T) {

	testcases := []struct {
		name         string
		reqbody      string
		expectedCode int
	}{
		{
			name:         "Valid case student",
			reqbody:      `{"yourName":"tempstud","password":"password","roleReq":"student","secretK":""}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Valid case student secretk",
			reqbody:      `{"yourName":"studd","password":"password","roleReq":"student","secretK":"$2a$15$NXTb8AxndfnaA82JWAxr2.apFmJkU.S1ROK10HmFBf69KxSCtW7S"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Valid case teacher",
			reqbody:      `{"yourName":"teachp","password":"password","roleReq":"teacher","secretK":""}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Valid case teacher secretk",
			reqbody:      `{"yourName":"teachpro","password":"password","roleReq":"teacher","secretK":"$2a$15$NXTb8AxndfnaA82JWAxr2.apFmJkU.S1ROK10HmFBf69KxSCtW7S"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Valid case admin",
			reqbody:      `{"yourName":"tempadm","password":"password","roleReq":"admin"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "Valid case admin secretk",
			reqbody:      `{"yourName":"addmmin","password":"password","roleReq":"admin","secretK":"$2a$15$NXTb8AxndfnaA82JWAxr2.apFmJkU.S1ROK10HmFBf69KxSCtW7S"}`,
			expectedCode: http.StatusOK,
		},
		{
			name:         "invalid name",
			reqbody:      `{"yourName":"studd55","password":"password","roleReq":"student","secretK":"$2a$15$NXTb8AxndfnaA82JWAxr2.apFmJkU.S1ROK10HmFBf69KxSCtW7S"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid pwd",
			reqbody:      `{"yourName":"studd","password":"psword","roleReq":"teacher"}`,
			expectedCode: http.StatusBadRequest,
		},
		{
			name:         "invalid role",
			reqbody:      `{"yourName":"studd","password":"password","roleReq":"sweeper","secretK":"$2a$15$NXTb8AxndfnaA82JWAxr2.apFmJkU.S1ROK10HmFBf69KxSCtW7S"}`,
			expectedCode: http.StatusBadRequest,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(w)
			req, err := http.NewRequest(http.MethodPost, "/register", bytes.NewBufferString(tc.reqbody))
			if err != nil {
				t.Fatalf("failed to create request: %v", err)
			}
			req.Header.Set("Content-Type", "application/json")
			ctx.Request = req
			CreatePendingReq(ctx)
			if w.Code != tc.expectedCode {
				t.Errorf("%s in this test - expected status %d, got %d", tc.name, tc.expectedCode, w.Code)
			}
			t.Logf("%s - testname, Response = %s", tc.name, w.Body.String())
		})
	}
}
