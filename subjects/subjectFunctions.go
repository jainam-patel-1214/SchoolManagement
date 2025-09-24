package subjects

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:admin123@tcp(127.0.0.1:3306)/goLearn"

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
		if !res.Next() {
			_, err = db.Exec("CREATE TABLE subjects (SubId int NOT NULL UNIQUE, SubName varchar(255), Grade varchar(4), Credits int, PRIMARY KEY(Tid))")
			if err != nil {
				fmt.Println("ISSUE WHILE CREATING TABLE")
				sendJSONResponse(writer, 500, "Internal Server Error", "ERROR CREATING TABLE")
				return
			}
			fmt.Println("table created")
		}
		sendJSONResponse(writer, 500, "Internal Server Error", "CANNOT FIND DATABASE")
		return
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
	_, err = db.Exec("INSERT INTO subjects VALUES (SubId, SubName, Grade, Credits) VALUES (?,?,?,?)", data.SubId, data.SubName, data.Grade, data.Credits)
	if err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "ERROR ADDING IN DB")
		return
	}
	sendJSONResponse(writer, 200, "OK", "DATA ADDED SUCCESSFULLY")
}
