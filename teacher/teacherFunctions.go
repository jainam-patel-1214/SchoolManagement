package teacher

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

const dsn = "root:admin123@tcp(127.0.0.1:3306)/goLearn"

type TeacherSubAllocation struct {
	Tid     int
	SubId   int
	Section string
}
type ReturnMsg struct {
	Code    int
	Status  string
	Message string
}
type TeacherInfo struct {
	Tid            int    `json:"id" binding:"required"`
	Name           string `json:"name" binding:"required"`
	ClassAllocated string `json:"clasTeacher" binding:"required"`
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

func HandleSubmitFaculty(writer http.ResponseWriter, reader *http.Request) {
	// SELECT * FROM information_schema.tables WHERE table_schema='goLearn' AND table_name='moreInfo';
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error opening DB: ", err)
		sendJSONResponse(writer, 500, "Internal Server Error", "ERROR CONNECTING DATABASE")
		return
	}
	defer db.Close()
	res, err := db.Query("SELECT * FROM information_schema.tables WHERE table_schema='goLearn' AND table_name='teachers'")
	if err != nil {
		fmt.Println("database not defined")

		sendJSONResponse(writer, 500, "Internal Server Error", "CANNOT FIND DATABASE")
		return
	}
	if !res.Next() {
		_, err = db.Exec("CREATE TABLE teachers (Tid int NOT NULL UNIQUE, Name varchar(255), ClassAllocated varchar(4), PRIMARY KEY(Tid))")
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
		sendJSONResponse(writer, 404, "Not Found", "BODY UNREADABLE")
		return
	}
	var data TeacherInfo
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 404, "Not Found", "REQUIRED FIELDS EMPTY")
		return
	}
	transisiton, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = db.Exec(`INSERT INTO teachers (Tid, Name, ClassAllocated) VALUES (?,?,?)`, data.Tid, data.Name, data.ClassAllocated)
	if err != nil {
		fmt.Println(err)
		transisiton.Rollback()
	}
	err = transisiton.Commit()
	if err != nil {
		panic(err)
	}
	sendJSONResponse(writer, 200, "Ok", "DATA ADDED SUCCESSFULLY")
	fmt.Println("commited and saved successfully")
}

func AddTeacherMoreInfo(writer http.ResponseWriter, reader *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "CANNOT CONNECT TO DB")
	}
	defer db.Close()
	res, err := db.Query("SHOW TABLES LIKE 'teacherMoreInfo'")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !res.Next() {
		_, err = db.Exec("CREATE TABLE teacherMoreInfo (Tid int NOT NULL, SubId int NOT NULL, Section varchar(4), FOREIGN KEY (Tid) REFERENCES teachers(Tid), FOREIGN KEY (SubId) REFERENCES subjects(SubId) )")
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
		sendJSONResponse(writer, 500, "Internal Server Error", "UNABLE TO READ BODY")
		return
	}
	var data TeacherSubAllocation
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "UNMARSHALING ERROR")
		return
	}

	if _, err := db.Exec("INSERT INTO teacherMoreInfo (Tid, SubId, Section) VALUES (?,?,?)", data.Tid, data.SubId, data.Section); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "UNABLE TO INSERT INTO DB")
		return
	}
	sendJSONResponse(writer, 200, "Ok", "DATA SAVED SUCCESSFULLY")
}

func AddReviews(writer http.ResponseWriter, reader *http.Request) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "CANNOT CONNECT TO DB")
	}
	defer db.Close()
	res, err := db.Query("SHOW TABLES LIKE 'reviews'")
	if err != nil {
		fmt.Println(err)
		return
	}
	if !res.Next() {
		_, err = db.Exec("CREATE TABLE reviews (Tid int NOT NULL, RollNo int NOT NULL, Review varchar(255), FOREIGN KEY (Tid) REFERENCES teachers(Tid), FOREIGN KEY (RollNo) REFERENCES students(RollNo))")
		if err != nil {
			fmt.Println(err)
			sendJSONResponse(writer, 500, "Internal Server Error", "ERROR CREATING TABLE")
			return
		}
		fmt.Println("table created")
	}
	defer res.Close()
	body, err := io.ReadAll(reader.Body)
	if err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "UNABLE TO READ BODY")
		return
	}
	var data TeacherSubAllocation
	if err = json.Unmarshal(body, &data); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "UNMARSHALING ERROR")
		return
	}

	if _, err := db.Exec("INSERT INTO reviews (Tid, SubId, Section) VALUES (?,?,?)", data.Tid, data.SubId, data.Section); err != nil {
		fmt.Println(err)
		sendJSONResponse(writer, 500, "Internal Server Error", "UNABLE TO INSERT INTO DB")
		return
	}
	sendJSONResponse(writer, 200, "Ok", "DATA SAVED SUCCESSFULLY")
}
