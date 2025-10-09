package admin

import (
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

const dsn = "root:admin123@tcp(127.0.0.1:3306)/goLearn"

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
			UserId   any    `json:"Uid" binding:"required"`
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
			if body.Std == 0 || body.Section == "" || body.UserId == 0 || body.UserName == "" || body.UserPwd == "" {
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
