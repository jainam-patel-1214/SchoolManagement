package admin

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type DisplayConditions struct {
	ViewByStd     int    `json:"viewByStd"`
	ViewBySection string `json:"viewBySection"`
	MinPercent    int    `json:"minPercent"`
	MaxPercent    int    `json:"maxPercent"`
}

func DisplayStudents(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		fmt.Println("no token found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthirused access"})
		return
	}
	if role == "admin" {
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
}

func DisplaySubject(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		fmt.Println("no token found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthirused access"})
		return
	}
	if role == "admin" {
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
			SubId   int    `json:"subjectId"`
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
}

func Report(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthirused access"})
		return
	} else {
		type MarkJson struct {
			SubjectId     int    `json:"subId"`
			Subject       string `json:"subjectName"`
			TheoryMark    int    `json:"theoryMM"`
			PracticalMark int    `json:"practicalMM"`
			Grade         string `json:"grade"`
		}
		type Comments struct {
			TeacherId   string `json:"tId"`
			TeacherName string `json:"tName"`
			Comment     string `json:"comment"`
		}
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect to db"})
			return
		}
		defer db.Close()
		var Param struct {
			StudentGrNo int `json:"grNo" binding:"required"`
		}
		err = ctx.BindJSON(&Param)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "id not found"})
			return
		}
		temp := fmt.Sprintf("%v", Param.StudentGrNo)
		fmt.Println("temp var", temp)
		tc, err := db.Begin()
		if err != nil {
			log.Fatal(err)
			return
		}
		res1, err := db.Query("SELECT m.subId,s.subName,m.theoryM,m.practicalM,m.grade FROM marks m INNER JOIN subjects s ON s.subId = m.subId WHERE m.grNo = ?", temp)
		if err != nil {
			tc.Rollback()
			log.Fatal(err)
			return
		}
		_, err = tc.Exec("SAVEPOINT query1done")
		if err != nil {
			tc.Rollback()
			log.Fatal("Failed to create savepoint:", err)
		}
		res2, err := db.Query("SELECT r.tId,t.tName,r.comment FROM reviews r INNER JOIN teachers t ON t.tId = r.tId WHERE r.grNo = ?", temp)
		if err != nil {
			_, err = tc.Exec("ROLLBACK TO SAVEPOINT query1done")
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error processing query"})
				return
			}
		}
		if err = tc.Commit(); err != nil {
			log.Fatal("Failed to commit transaction:", err)
		}
		var otpt struct {
			MarkInfo    []MarkJson
			CommentInfo []Comments
		}
		for res1.Next() {
			var tp MarkJson
			err = res1.Scan(&tp.SubjectId, &tp.Subject, &tp.TheoryMark, &tp.PracticalMark, &tp.Grade)
			if err != nil {
				fmt.Println(err)
				// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant process query output"})
				return
			}
			otpt.MarkInfo = append(otpt.MarkInfo, tp)
		}
		for res2.Next() {
			var tp Comments
			err = res2.Scan(&tp.TeacherId, &tp.TeacherName, &tp.Comment)
			if err != nil {
				fmt.Println(err)
				// ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant process query output"})
				return
			}
			otpt.CommentInfo = append(otpt.CommentInfo, tp)
		}

		ctx.JSON(http.StatusOK, gin.H{"output": otpt})
	}
}

func AddStudent(ctx *gin.Context) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	} else {
		var studentData struct {
			GR_NO       int    `json:"grNo" binding:"required"`
			StudentPwd  string `json:"studPwd" binding:"required"`
			UserRole    string `json:"userRole" binding:"required"`
			StudentName string `json:"studName" binding:"required"`
			Std         int    `json:"std" binding:"required"`
			Section     string `json:"section" binding:"required"`
		}
		err = ctx.Bind(&studentData)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot read body"})
			return
		}
		_, err = db.Exec("INSERT INTO students (grNo, sPwd, userRole, studName, std, section) VALUES (?,?,?,?,?,?)", studentData.GR_NO, studentData.StudentPwd, studentData.UserRole, studentData.StudentName, studentData.Std, studentData.Section)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while inserting into db"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "student created"})
	}
}

func EditStud(ctx *gin.Context) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	} else {
		type EditBody struct {
			GR_No       int    `json:"grNo" binding:"required"`
			StudentPwd  string `json:"studPwd"`
			UserRole    string `json:"userRole"`
			StudentName string `json:"studName"`
			Std         int    `json:"std"`
			Section     string `json:"section"`
		}
		var editBody EditBody
		err = ctx.Bind(&editBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while processing data"})
			return
		}
		var defaultData EditBody
		err = db.QueryRow("SELECT * FROM students WHERE grNo=?", editBody.GR_No).Scan(&defaultData.GR_No, &defaultData.StudentPwd, &defaultData.UserRole, &defaultData.StudentName, &defaultData.Std, &defaultData.Section)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while processing data"})
			return
		}
		dbstr := "UPDATE students SET "
		var conditions []string
		if editBody.StudentPwd != defaultData.StudentPwd && editBody.StudentPwd != "" {
			conditions = append(conditions, ("sPwd = '" + editBody.StudentPwd + "'"))
		}
		if editBody.StudentName != defaultData.StudentName && editBody.StudentName != "" {
			conditions = append(conditions, ("studName = '" + editBody.StudentName + "'"))
		}
		if editBody.Std != defaultData.Std && editBody.Std <= 12 && editBody.Std > 0 {
			conditions = append(conditions, ("std = " + strconv.Itoa(editBody.Std)))
		}
		if editBody.Section != defaultData.Section && editBody.Section != "" {
			conditions = append(conditions, ("section = '" + editBody.Section + "'"))
		}
		for i, v := range conditions {
			dbstr += v
			if i != len(conditions)-1 {
				dbstr += ","
			}
		}
		dbstr += ("WHERE grNo = " + strconv.Itoa(editBody.GR_No))

		_, err = db.Exec(dbstr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating db"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "student updated successfully"})
		return
	}
}

func CreateSub(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	}
	if role == "admin" {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot connect to db"})
			return
		}
		defer db.Close()
		var subInfo struct {
			SubId    int    `json:"subId" binding:"required"`
			SubName  string `json:"subName" binding:"required"`
			LevelStd int    `json:"levelStd" binding:"required"`
			Credits  int    `json:"credits" binding:"required"`
		}
		if err = ctx.Bind(&subInfo); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error processing data"})
			return
		}
		var limit int
		err = db.QueryRow("SELECT subject_limit FROM subjectAllocation WHERE std = ?", subInfo.LevelStd).Scan(&limit)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error processing data"})
			return
		}
		var count int
		err = db.QueryRow("SELECT COUNT(subId) FROM subjects WHERE levelStd = ?", subInfo.LevelStd).Scan(&count)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error processing data"})
			return
		}
		if count >= limit {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("limit of subject for standard %d reached", subInfo.LevelStd)})
			return
		}
		if _, err = db.Exec("INSERT INTO subjects (subId,subName,levelStd,credits) VALUES (?,?,?,?)", subInfo.SubId, subInfo.SubName, subInfo.LevelStd, subInfo.Credits); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting data"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "inserted successfully"})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}
}

func EditSub(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	}
	if role == "admin" {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect db"})
			return
		}
		defer db.Close()
		type EditBody struct {
			SubId    int    `json:"subId" binding:"required"`
			SubName  string `json:"subName"`
			LevelStd int    `json:"levelStd"`
			Credits  int    `json:"credits"`
		}
		var editBody EditBody
		err = ctx.Bind(&editBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while processing data"})
			return
		}
		var defaultData EditBody
		err = db.QueryRow("SELECT * FROM subjects WHERE subId=?", editBody.SubId).Scan(&defaultData.SubId, &defaultData.SubName, &defaultData.LevelStd, &defaultData.Credits)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while processing data"})
			return
		}
		dbstr := "UPDATE subjects SET "
		var conditions []string
		if editBody.SubId != defaultData.SubId && editBody.SubId != 0 {
			conditions = append(conditions, ("subId = '" + strconv.Itoa(editBody.SubId) + "'"))
		}
		if editBody.SubName != defaultData.SubName && editBody.SubName != "" {
			conditions = append(conditions, ("subName = '" + editBody.SubName + "'"))
		}
		if editBody.LevelStd != defaultData.LevelStd && editBody.LevelStd <= 12 && editBody.LevelStd > 0 {
			conditions = append(conditions, ("levelStd = " + strconv.Itoa(editBody.LevelStd)))
		}
		if editBody.Credits != defaultData.Credits && editBody.Credits != 0 {
			conditions = append(conditions, ("credits = '" + strconv.Itoa(editBody.Credits) + "'"))
		}
		for i, v := range conditions {
			dbstr += v
			if i != len(conditions)-1 {
				dbstr += ","
			}
		}
		dbstr += ("WHERE subId = " + strconv.Itoa(editBody.SubId))

		_, err = db.Exec(dbstr)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating db"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "subject updated successfully"})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}
}

func EnterMarks(ctx *gin.Context) {

	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	}
	if role == "admin" {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cannot connect to db"})
			return
		}
		defer db.Close()
		var marks struct {
			GrNo           int `json:"grNo" binding:"required"`
			SubId          int `json:"subId" binding:"required"`
			TheoryMarks    int `json:"theoryMarks" binding:"required"`
			PracticalMarks int `json:"practicalMarks" binding:"required"`
		}
		if err = ctx.Bind(&marks); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		var amount int
		err = db.QueryRow("SELECT COUNT(grNo) FROM marks WHERE grNo = ? AND subId = ?", marks.GrNo, marks.SubId).Scan(&amount)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		if amount > 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "record already present please try updating it"})
			return
		}
		if _, err = db.Exec("INSERT INTO marks (grNo,subId,theoryM,practicalM,grade) VALUES (?,?,?,?,?)", marks.GrNo, marks.SubId, marks.TheoryMarks, marks.PracticalMarks, gradeCalculator(marks.TheoryMarks+marks.PracticalMarks)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error inserting data"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "inserted successfully"})
		return
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

}
func EditMarks(ctx *gin.Context) {

	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	}
	if role == "admin" {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect db"})
			return
		}
		defer db.Close()
		type marks struct {
			GrNo           int `json:"grNo" binding:"required"`
			SubId          int `json:"subId" binding:"required"`
			TheoryMarks    int `json:"theoryMarks"`
			PracticalMarks int `json:"practicalMarks"`
		}
		var editBody marks
		var tempgrade string
		err = ctx.Bind(&editBody)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while processing data"})
			return
		}
		var defaultData marks
		err = db.QueryRow("SELECT * FROM marks WHERE grNo=?", editBody.GrNo).Scan(&defaultData.GrNo, &defaultData.SubId, &defaultData.TheoryMarks, &defaultData.PracticalMarks, &tempgrade)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while processing data"})
			return
		}
		dbstr := "UPDATE marks SET "
		changeOccur := false
		var conditions []string
		if editBody.TheoryMarks != defaultData.TheoryMarks && editBody.TheoryMarks <= 80 && editBody.TheoryMarks >= 0 {
			conditions = append(conditions, ("theoryM = " + strconv.Itoa(editBody.TheoryMarks)))
			changeOccur = true
		}
		if editBody.PracticalMarks != defaultData.PracticalMarks && editBody.PracticalMarks <= 20 && editBody.PracticalMarks >= 0 {
			conditions = append(conditions, ("practicalM = " + strconv.Itoa(editBody.PracticalMarks)))
			changeOccur = true
		}
		needComma := false
		for i, v := range conditions {
			dbstr += v
			needComma = true
			if i != len(conditions)-1 {
				dbstr += ","
			}
		}
		if needComma {
			dbstr += ","
		}
		if changeOccur {
			dbstr += fmt.Sprintf(" grade = '%s' ", gradeCalculator(editBody.TheoryMarks+editBody.PracticalMarks))
		}
		dbstr += fmt.Sprintf("WHERE grNo = %s AND subId = %s", strconv.Itoa(editBody.GrNo), strconv.Itoa(editBody.SubId))
		fmt.Println(dbstr)
		if changeOccur {
			_, err = db.Exec(dbstr)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating db"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"output": "student updated successfully"})
			return
		} else {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "no changes specified"})
			return
		}
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized user"})
		return
	}

}

func Performance(ctx *gin.Context) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	}
	var TeacherId struct {
		Tid string `json:"tid" binding:"required"`
	}
	err = ctx.BindJSON(&TeacherId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "teacher id not found"})
		return
	}
	res := db.QueryRow("SELECT stdAllocated FROM teachers WHERE tId = ?", TeacherId.Tid)
	var std int
	if err = res.Scan(&std); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong"})
		return
	}
	type Teachers struct {
		Tid                 string
		TName               string
		StdAllocated        int
		SubName             string
		TotalTheoryMarks    int
		TotalPracticalMarks int
	}
	var result []Teachers
	res2, err := db.Query("SELECT tId FROM teachers WHERE stdAllocated=? AND subId IS NOT NULL", std)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong hwile fetching db"})
		return
	}
	var tempres Teachers
	for res2.Next() {
		var temp any
		err = res2.Scan(&temp)
		if err != nil {
			fmt.Println("cannot scan", err)
		} else {
			res3, err := db.Query("SELECT t.tId, t.tName, t.stdAllocated, s.subName, SUM(m.theoryM) AS totalTheory, SUM(m.practicalM) AS totalPractical FROM marks m LEFT JOIN teachers t ON m.subId = t.subId INNER JOIN subjects s ON t.subId = s.subId WHERE t.tId = ? GROUP BY t.tId, t.tName, t.stdAllocated, s.subName", temp)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
				return
			}
			if res3.Next() {
				err = res3.Scan(&tempres.Tid, &tempres.TName, &tempres.StdAllocated, &tempres.SubName, &tempres.TotalTheoryMarks, &tempres.TotalPracticalMarks)
				if err != nil {
					fmt.Println("error finding data", err)
					return
				}
				result = append(result, tempres)
			}
		}
	}
	fmt.Println(result)
	ctx.JSON(http.StatusOK, gin.H{"output": result})
}

func DelStud(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		fmt.Println("no token found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthirused access"})
		return
	}
	if role == "admin" {
		var stdGrno struct {
			GRno int `json:"grNo" binding:"required"`
		}
		if err := ctx.Bind(&stdGrno); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read body"})
			return
		}

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "CANNOT CONNECT TO DB"})
			return
		}
		defer db.Close()
		if _, err = db.Exec("DELETE FROM students WHERE grNo = ?", stdGrno.GRno); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error in DB, cant delete"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "DELETED SUCCESSFULLY"})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		return
	}
}

func DelSub(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		fmt.Println("no token found")
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthirused access"})
		return
	}
	if role == "admin" {
		var subid struct {
			SubId int `json:"subId" binding:"required"`
		}
		if err := ctx.BindJSON(&subid); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read body"})
			return
		}
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "CANNOT CONNECT TO DB"})
			return
		}
		defer db.Close()
		if _, err = db.Exec("DELETE FROM subjects WHERE subId = ?", subid.SubId); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "DELETED SUCCESSFULLY"})

	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		return
	}
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
