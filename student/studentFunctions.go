package student

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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
	ViewByStd     int    `json:"viewByStd"`
	ViewBySection string `json:"viewBySection"`
	MinPercent    int    `json:"minPercent"`
	MaxPercent    int    `json:"maxPercent"`
}

func DisplayStudents(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist {
		fmt.Println("no token found")
		return
	}
	if role == "student" {
		var constraints DisplayConditions
		if err := ctx.BindJSON(&constraints); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}
		fmt.Println("555,", constraints)
		dbstr := "SELECT s.studName, s.std, s.section, sub.subName, m.theoryM, m.practicalM, m.grade FROM students s RIGHT JOIN marks m ON s.grNo = m.grNo INNER JOIN subjects sub ON m.subId = sub.subId"
		count := 0
		if constraints.ViewByStd != 0 {
			if count == 0 {
				dbstr += " WHERE "
				count++
			}
			dbstr += "s.std = " + strconv.Itoa(constraints.ViewByStd)
		}
		if constraints.ViewBySection != "" {
			if count == 0 {
				dbstr += " WHERE "
				count++
			} else if count > 0 {
				dbstr += " AND "
			}
			dbstr += "s.section = " + "'" + constraints.ViewBySection + "'"
		}
		if constraints.MinPercent != 0 {
			if count == 0 {
				dbstr += " WHERE "
				count++
			} else if count > 0 {
				dbstr += " AND "
			}
			dbstr += "(m.theoryM+m.practicalM) > " + strconv.Itoa(constraints.MinPercent)
		}
		if constraints.MaxPercent != 0 {
			if count == 0 {
				dbstr += " WHERE "
				count++
			} else if count > 0 {
				dbstr += " AND "
			}
			dbstr += "(m.theoryM + m.practicalM) < " + strconv.Itoa(constraints.MaxPercent)
		}
		fmt.Println(dbstr)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "CANNOT CONNECT TO DB"})
			return
		}
		res, err := db.Query(dbstr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch from db"})
		}
		defer db.Close()
		type output struct {
			SName     string `json:"studentName"`
			SSection  string `json:"section"`
			SubName   string `json:"subject"`
			Grade     string `json:"grade"`
			SStd      int    `json:"standard"`
			Theory    int    `json:"theoryMarks"`
			Practical int    `json:"practicalMarks"`
		}
		var queryres []output
		for res.Next() {
			var record output
			if err := res.Scan(&record.SName, &record.SStd, &record.SSection, &record.SubName, &record.Theory, &record.Practical, &record.Grade); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read database results"})
				return
			}
			queryres = append(queryres, record)
		}
		ctx.JSON(http.StatusOK, gin.H{"result": queryres})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
	}
	fmt.Println("trying debug")
}

func DisplaySubject(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist {
		fmt.Println("no token found")
		return
	}
	if role == "student" {
		var constraints struct {
			Std int `json:"std" binding:"required"`
		}
		if err := ctx.BindJSON(&constraints); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "INTERNAL SERVER ERROR"})
			return
		}
		dbstr := "SELECT * FROM subjects WHERE levelStd = " + strconv.Itoa(constraints.Std)
		fmt.Println(dbstr)

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "CANNOT CONNECT TO DB"})
			return
		}
		res, err := db.Query(dbstr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot fetch from db"})
		}
		defer db.Close()
		type output struct {
			SubId   int    `json:"studentId"`
			SubName string `json:"subjectName"`
			Std     int    `json:"level"`
			Credits int    `json:"credits"`
		}
		var queryres []output
		for res.Next() {
			var record output
			if err := res.Scan(&record.SubId, &record.SubName, &record.Std, &record.Credits); err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Cannot read database results"})
				return
			}
			queryres = append(queryres, record)
		}
		ctx.JSON(http.StatusOK, gin.H{"result": queryres})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
	}
	fmt.Println("trying debug")
}
