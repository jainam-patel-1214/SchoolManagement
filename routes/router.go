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
			teach.GET("/displayPerformance", teacher.Performance)
			teach.POST("/createStud", teacher.AddStudent)
			teach.PUT("/updateStud", teacher.EditStud)
			teach.POST("/createSub", teacher.CreateSub)
			teach.PUT("/updateSub", teacher.EditSub)
			teach.POST("/enterMarks", teacher.EnterMarks)
			teach.PUT("/updateMarks", teacher.EditMarks)
			teach.GET("/displaySub", student.DisplayStudents)
			teach.POST("/addReview", teacher.AddReviews)
		}
	}
	{
		admn := r.Group("/admin")
		admn.Use(middleware.ValidateSession())
		{
			admn.GET("/pendingRequest", admin.ShowPendingReq)
			admn.POST("/acceptRequest", admin.AcceptPendingReq)
		}
	}
	return r
}
