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
		fmt.Println("wrong role")
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
		return
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "CANNOT CONNECT TO DB"})
		return
	}

	defer db.Close()
	fmt.Println("trying debug")
}
