package routes

import (
	"example.com/main/admin"
	"example.com/main/middleware"
	"example.com/main/student"
	"example.com/main/teacher"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()

	// CreateSession is not actually a middleware but validate session is
	r.POST("/login", middleware.CreateSession)
	r.POST("/register", admin.CreatePendingReq)
	{
		stud := r.Group("/student")
		stud.Use(middleware.ValidateSession())
		{
			stud.GET("/display", student.DisplayStudents)
			stud.GET("/displaySub", student.DisplaySubject)
			stud.GET("/report", student.Report)
		}
	}
	{
		teach := r.Group("/teacher")
		teach.Use(middleware.ValidateSession())
		{
			teach.GET("/studentreport", teacher.Report)
			teach.GET("/displayPerformance", teacher.Performance)
			teach.POST("/createStud", teacher.AddStudent)
			teach.PUT("/updateStud", teacher.EditStud)
			teach.POST("/createSub", teacher.CreateSub)
			teach.PUT("/updateSub", teacher.EditSub)
			teach.POST("/enterMarks", teacher.EnterMarks)
			teach.PUT("/updateMarks", teacher.EditMarks)
			teach.GET("/displaySub", student.DisplaySubject)
			teach.POST("/addReview", teacher.AddReviews)
			teach.DELETE("/delStudent", teacher.DelStud)
			teach.DELETE("/delSubject", teacher.DelSub)
		}
	}
	{
		admn := r.Group("/admin")
		admn.Use(middleware.ValidateSession())
		{
			admn.GET("/pendingRequest", admin.ShowPendingReq)
			admn.POST("/acceptRequest", admin.AcceptPendingReq)
			admn.DELETE("/rejectRequest", admin.RejectRequest)
			admn.DELETE("/delTeacher", admin.DeleteTeacher)
			admn.POST("/addTeacher", admin.AddTeacher)
			admn.PUT("/editTeacher", admin.EditTeacher)

			admn.GET("/display", admin.DisplayStudents)
			admn.GET("/displaySub", admin.DisplaySubject)
			admn.GET("/studentreport", admin.Report)

			admn.GET("/displayTeacherPerformance", admin.Performance)
			admn.POST("/createStud", admin.AddStudent)
			admn.PUT("/updateStud", admin.EditStud)
			admn.POST("/createSub", admin.CreateSub)
			admn.PUT("/updateSub", admin.EditSub)
			admn.POST("/enterMarks", admin.EnterMarks)
			admn.PUT("/updateMarks", admin.EditMarks)
			admn.DELETE("/delStudent", admin.DelStud)
			admn.DELETE("/delSubject", admin.DelSub)
			admn.POST("/setSubLimit", admin.SetSubLimit)
		}
	}
	return r
}
