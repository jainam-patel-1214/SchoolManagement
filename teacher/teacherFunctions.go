package teacher

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
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

func AddStudent(ctx *gin.Context) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	role, exist := ctx.Get("userrole")
	if !exist || role != "teacher" {
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
	if !exist || role != "teacher" {
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
		if editBody.StudentPwd != "" {
			conditions = append(conditions, ("sPwd = '" + editBody.StudentPwd + "'"))
		}
		if editBody.StudentName != "" {
			conditions = append(conditions, ("studName = '" + editBody.StudentName + "'"))
		}
		if editBody.Std != 0 && editBody.Std <= 12 && editBody.Std > 0 {
			conditions = append(conditions, ("std = " + strconv.Itoa(editBody.Std)))
		}
		if editBody.Section != "" {
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

}

func EditSub(ctx *gin.Context) {

}

func AddReviews(ctx *gin.Context) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	role, exist := ctx.Get("userrole")
	if !exist || role != "teacher" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	}
	tid, exist := ctx.Get("UiD")
	if !exist || tid == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error setting your id"})
		return
	}
	var reviewInfo struct {
		StudId  int    `json:"grNo" binding:"required"`
		Comment string `json:"comment" binding:"required"`
	}
	err = ctx.BindJSON(&reviewInfo)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read the data sent"})
	}
	if _, err := db.Exec("INSERT INTO reviews (tId, grNo, comment) VALUES (?,?,?)", tid.(string), reviewInfo.StudId, reviewInfo.Comment); err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while inserting error"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"output": "added successfully"})
	return
}

func Performance(ctx *gin.Context) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()
	role, exist := ctx.Get("userrole")
	if !exist || role != "teacher" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorizes access"})
		return
	}
	tid, exist := ctx.Get("UiD")
	if !exist || tid == nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error setting your id"})
		return
	}
	res := db.QueryRow("SELECT stdAllocated FROM teachers WHERE tId = ?", tid)
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
			res3, err := db.Query("SELECT t.tId, t.tName, t.stdAllocated, s.subName, SUM(m.theoryM) AS totalTheory, SUM(m.practicalM) AS totalPractical FROM marks m INNER JOIN students st ON st.grNo = m.grNo LEFT JOIN subjects s ON m.subId = s.subId LEFT JOIN teachers t ON s.subId = t.subId WHERE t.tId = ? AND t.stdAllocated = st.std GROUP BY t.tId, t.tName, s.subName", temp)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "something went wrong hwile fetching db"})
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
	return
}
