package student

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const dsn = "root:admin123@tcp(127.0.0.1:3306)/goLearn"

type TeacherInfo struct {
	ID      int    `json:"id" binding:"required"`
	Name    string `json:"name" binding:"required"`
	Subject string `json:"subject" binding:"required"`
	ClassID string `json:"classId" binding:"required"`
}

type ReturnMsg struct {
	Code    int    `json:"statusCode" binding:"required"`
	Status  string `json:"status" binding:"required"`
	Message string `json:"response"`
}

type StudentInfo struct {
	RollNo      int    `json:"rollNo" binding:"required"`
	Name        string `json:"name" binding:"required"`
	MathsMark   int    `json:"mathsMark" binding:"required"`
	ScienceMark int    `json:"scienceMark" binding:"required"`
	EnglishMark int    `json:"englishMark" binding:"required"`
	ClassID     string `json:"classId" binding:"required"`
}

type DisplayConditions struct {
	ViewByClass   bool   `json:"veiwByClass" binding:"required"`
	ClassID       string `json:"classId" binding:"required"`
	ViewByPercent bool   `json:"viewByPercent" binding:"required"`
	StartPercent  int    `json:"startPercent" binding:"required"`
	EndPercent    int    `json:"endPercent" binding:"required"`
	ViewByMarks   bool   `json:"viewByMark" binding:"required"`
}

func DisplayStudents(writer http.ResponseWriter, reader *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		var fail = ReturnMsg{Code: 500, Status: "Internal Server Error", Message: "CANNOT CONNECT DATABSE"}
		b, err1 := json.Marshal(fail)
		if err1 != nil {
			fmt.Println(err1)
		}
		writer.Write(b)
		return
	}
	defer db.Close()
	req, err := io.ReadAll(reader.Body)
	if err != nil {
		var fail = ReturnMsg{Code: 404, Status: "Not found", Message: "REQUIRED FIELDS EMPTY"}
		b, err1 := json.Marshal(fail)
		if err1 != nil {
			fmt.Println(err1)
		}
		writer.Write(b)
		return
	}

	type Output struct {
		Name string `json:"studentName"`
	}

	var constraints DisplayConditions
	if err = json.Unmarshal(req, &constraints); err != nil {
		fmt.Println("error while reading paramaeterrs")
		return
	}
	if constraints.ViewByClass && !constraints.ViewByPercent {

	} else if constraints.ViewByPercent && !constraints.ViewByClass {

	} else if constraints.ViewByClass && constraints.ViewByPercent {

	}
}
