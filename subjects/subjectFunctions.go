package subjects

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"example.com/main/database"
	_ "github.com/go-sql-driver/mysql"
)

var dsn = database.InitDb()

type SubInfo struct {
	SubId   int    `json:"SubId" binging:"required"`
	SubName string `json:"SubName" binding:"required"`
	Grade   int    `json:"grade" binding:"required"`
	Credits int    `json:"credits" binding:"required"`
}

type ReturnMsg struct {
	Code    int
	Status  string
	Message string
}

type SubAllocations struct {
	Section          string `json:"section" biding:"required"`
	SubAmtAllowed    int    `json:"subAmt" biding:"required"`
	TotalCredAllowed int    `json:"totalCred" biding:"required"`
}

type Marks struct {
	RollNo       int    `json:"studRollNo" binding:"required"`
	SubId        int    `json:"subId" binding:"required"`
	TheoryM      int    `json:"theoryM" binding:"required"`
	PracticalM   int    `json:"practicalM" binding:"required"`
	OverallGrade string `json:"overallGrade" binding:"required"`
}

func sendJSONResponse(writer http.ResponseWriter, code int, status, message string) {
	writer.WriteHeader(code)
	response := ReturnMsg{Code: code, Status: status, Message: message}
	b, err := json.Marshal(response)
	if err != nil {
		fmt.Println("JSON Marshal error:", err)
		return
	}
	writer.Write(b)
}

func gradeCalculator(n int) string {
	if n > 90 {
		return "AA"
	} else if n > 80 && n <= 90 {
		return "AB"
	} else if n > 70 && n <= 80 {
		return "BB"
	} else if n > 60 && n <= 70 {
		return "BC"
	} else if n > 50 && n <= 60 {
		return "CC"
	} else if n > 40 && n <= 50 {
		return "CD"
	} else if n > 30 && n <= 40 {
		return "DD"
	}
	return "FF"
}

func AddSubjectInfo(writer http.ResponseWriter, reader *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("Cannot connect to DB", err)
		sendJSONResponse(writer, 500, "Internal Server Error", "ERROR CONNECTING DATABASE")
		return
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM information_schema.tables WHERE table_schema='goLearn' AND table_name='subjects'")
	if err != nil {
		fmt.Println("database not defined")
		sendJSONResponse(writer, 500, "Internal Server Error", "CANNOT FIND DATABASE")
		return
	}
	if !res.Next() {
		_, err = db.Exec("CREATE TABLE subjects (SubId int NOT NULL UNIQUE, SubName varchar(255), Grade varchar(4), Credits int, PRIMARY KEY(Tid))")
		if err != nil {
			fmt.Println("ISSUE WHILE CREATING TABLE")
			sendJSONResponse(writer, 500, "Internal Server Error", "ERROR CREATING TABLE")
			return
		}
		fmt.Println("table created")
	}
	defer res.Close()

	body, err := io.ReadAll(reader.Body)
	if err != nil {
		fmt.Println("UNABLE TO READ")
		sendJSONResponse(writer, 404, "Not FOunt", "CONTENT BODY NOT FOUND")
		return
	}
	var data SubInfo
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println("json unmarshal error")
		return
	}
	_, err = db.Exec("INSERT INTO subjects (SubId, SubName, Grade, Credits) VALUES (?,?,?,?)", data.SubId, data.SubName, data.Grade, data.Credits)
	if err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "ERROR ADDING IN DB")
		return
	}
	sendJSONResponse(writer, 200, "OK", "DATA ADDED SUCCESSFULLY")
}

func AddSubjectAllocation(writer http.ResponseWriter, reader *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println("UNABLE TO ACCESS DATABASE")
		sendJSONResponse(writer, 500, "Internal Server Error", "CANNOT FIND DATABASE")
		return
	}
	defer db.Close()

	res, err := db.Query("SELECT * FROM information_schema.tables WHERE table_schema='goLearn' AND table_name='subjectAllocation'")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !res.Next() {
		_, err = db.Exec("CREATE TABLE subjectAllocation (Section varchar(4) NOT NULL UNIQUE, SubAmtAllowed int, TotalCredAllowed int, PRIMARY KEY(Section))")
		if err != nil {
			fmt.Println("ISSUE WHILE CREATING TABLE")
			sendJSONResponse(writer, 500, "Internal Server Error", "ERROR CREATING TABLE")
			return
		}
		fmt.Println("table created")
	}
	defer res.Close()

	body, err := io.ReadAll(reader.Body)
	if err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 404, "Not Found", "BODY NOT FOUND")
		return
	}
	var data SubAllocations
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal server error", "ERROR WHILE UNMARSHALING")
		return
	}

	if _, err = db.Exec("INSERT INTO subjectAllocation (Section,SubAmtAllowed,TotalCredAllowed) VALUES (?,?,?,?)", data.Section, data.SubAmtAllowed, data.TotalCredAllowed); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal server error", "ERROR WHILE INSERTING INTO DB")
		return
	}
	sendJSONResponse(writer, 200, "Ok", "DATA ADDED SUCCESSFULLY")
	fmt.Println("commited and saved successfully")
}

func AddMarks(writer http.ResponseWriter, reader *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal server error", "CANT CONNECT TO DB")
	}
	res, err := db.Query("SELECT * FROM information_schema.tables WHERE table_schema='goLearn' AND table_name='marks'")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !res.Next() {
		_, err = db.Exec("CREATE TABLE marks (RollNo int NOT NULL, SubId int NOT NULL, TheoryM int, PracticalM int, Grade varchar(2), FOREIGN KEY (RollNo) REFERENCES students(RollNo), FOREIGN KEY (SubId) REFERENCES subjects(SubId))")
		if err != nil {
			fmt.Println("ISSUE WHILE CREATING TABLE")
			sendJSONResponse(writer, 500, "Internal Server Error", "ERROR CREATING TABLE")
			return
		}
		fmt.Println("table created")
	}
	defer res.Close()

	body, err := io.ReadAll(reader.Body)
	if err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 404, "Not Found", "BODY NOT FOUND")
		return
	}
	var data Marks

	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "ERROR WHILE UNMARSHALING")
		return
	}
	if _, err = db.Exec("INSERT INTO marks (RollNo, SubId, TheoryM, PracticalM, Grade) VALUES (?,?,?,?,?)", data.RollNo, data.SubId, data.TheoryM, data.PracticalM, gradeCalculator(data.TheoryM+data.PracticalM)); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "ERROR WHILE INSERTING INTO DATABASE")
	}
	fmt.Println("inserted successfully")
	sendJSONResponse(writer, 200, "Ok", "DATA ADDED SUCCESSFULLY")
}
