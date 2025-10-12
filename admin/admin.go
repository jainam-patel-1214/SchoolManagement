package admin

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"

	"example.com/main/database"
	"github.com/gin-gonic/gin"
)

var dsn = database.InitDb()

func CreatePendingReq(ctx *gin.Context) {
	var PendingDb struct {
		RoleRequested string `json:"roleReq"`
		Username      string `json:"userName"`
		Pwd           string `json:"pwd"`
	}
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect to db"})
		return
	}
	defer db.Close()

	err = ctx.Bind(&PendingDb)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "cannot read body"})
		return
	}
	_, err = db.Exec("INSERT INTO pendingApplications (username,role_requested,user_pwd) VALUES (?,?,?)", PendingDb.Username, PendingDb.RoleRequested, PendingDb.Pwd)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error processing query"})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"output": "request submitted"})
}

func AcceptPendingReq(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		return
	} else {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect to db"})
			return
		}
		defer db.Close()
		var body struct {
			UserName string `json:"uName" binding:"required"`
			UserPwd  string `json:"uPwd" binding:"required"`
			UserRole string `json:"uRole" binding:"required"`
			UserId   any    `json:"Uid"`
			Std      int    `json:"std"`
			Section  string `json:"section"`
			SubId    int    `json:"subId"`
		}
		err = ctx.Bind(&body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error reading body"})
			return
		}
		switch body.UserRole {
		case "student":
			if body.Std == 0 || body.Section == "" || body.UserId == 0 || body.UserId == "" || body.UserName == "" || body.UserPwd == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "fill userid/std/section accurately"})
				return
			}
			_, err = db.Exec("INSERT INTO students (grNo,sPwd,userRole,studName,std,section) VALUES (?,?,?,?,?,?)", body.UserId, body.UserPwd, body.UserRole, body.UserName, body.Std, body.Section)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating VAL in DB"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"output": "student created"})
			return
		case "teacher":
			if body.Std == 0 || body.Section == "" || body.UserId == 0 || body.UserName == "" || body.UserPwd == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "fill userid/std/section accurately"})
				return
			}
			_, err = db.Exec("INSERT INTO teachers (tId,tPwd,userRole,tName,subId,classAllocated) VALUES (?,?,?,?,?,?)", body.UserId, body.UserPwd, body.UserRole, body.UserName, body.SubId, strconv.Itoa(body.Std)+body.Section)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating VAL in DB"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"output": "teacher created"})
			return
		case "admin":
			if body.UserId == 0 || body.UserName == "" || body.UserPwd == "" {
				ctx.JSON(http.StatusBadRequest, gin.H{"error": "fill userid/name/pwd accurately"})
				return
			}
			_, err = db.Exec("INSERT INTO admins (admin_id,admin_name,admin_pwd) VALUES (?,?,?)", body.UserId, body.UserName, body.UserPwd)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while updating VAL in DB"})
				return
			}
			ctx.JSON(http.StatusOK, gin.H{"output": "admin created"})
			return
		default:
			ctx.JSON(http.StatusBadRequest, gin.H{"error": "invalid role provided"})
			return
		}
	}
}

func ShowPendingReq(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		return
	} else {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect to db"})
			return
		}
		defer db.Close()
		type PendingDb struct {
			RoleRequested string `json:"roleReq"`
			Username      string `json:"userName"`
			Pwd           string `json:"pwd"`
		}
		var result []PendingDb
		res, err := db.Query("SELECT * FROM pendingApplications")
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while fetching data"})
			return
		}
		for res.Next() {
			var temp PendingDb
			var x int
			err = res.Scan(&x, &temp.Username, &temp.RoleRequested, &temp.Pwd)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error scanning result of db"})
				return
			}
			result = append(result, temp)
		}
		ctx.JSON(http.StatusOK, gin.H{"output": result})
	}
}

func DeleteTeacher(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthirused access"})
		return
	}
	if role == "admin" {
		var tid struct {
			TId string `json:"teacherId" binding:"required"`
		}
		if err := ctx.Bind(&tid); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "unable to read body"})
			return
		}

		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "CANNOT CONNECT TO DB"})
			return
		}
		defer db.Close()
		if _, err = db.Exec("DELETE FROM teachers WHERE tId = ?", tid.TId); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error in DB, cant delete"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "DELETED SUCCESSFULLY"})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		return
	}
}

func RejectRequest(ctx *gin.Context) {
	role, exist := ctx.Get("userrole")
	if !exist || role != "admin" {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthorized access"})
		return
	} else if role == "admin" {
		db, err := sql.Open("mysql", dsn)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "cant connect to db"})
			return
		}
		defer db.Close()
		var body struct {
			UserName string `json:"uName" binding:"required"`
			UserPwd  string `json:"uPwd" binding:"required"`
			UserRole string `json:"uRole" binding:"required"`
			UserId   any    `json:"Uid"`
			Std      int    `json:"std"`
			Section  string `json:"section"`
			SubId    int    `json:"subId"`
		}
		err = ctx.Bind(&body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error reading body"})
			return
		}
		if _, err = db.Exec(fmt.Sprintf("DELETE FROM pendingApplications WHERE username='%s' AND user_pwd='%s' AND role_requested = '%s'", body.UserName, body.UserPwd, body.UserRole)); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "error while rejecting"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"output": "rejected successfully"})
	} else {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "unauthirused access"})
	}
}
