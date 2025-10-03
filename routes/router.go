package routes

import (
	"example.com/main/middleware"
	"example.com/main/student"
	"github.com/gin-gonic/gin"
)

func InitializeRouter() *gin.Engine {
	r := gin.Default()
	// CreateSession is not actually a middleware but validate session is
	r.POST("/login", middleware.CreateSession)
	{
		stud := r.Group("/student")
		stud.Use(middleware.ValidateSession())
		{
			stud.GET("/display", student.DisplayStudents)
			stud.GET("/displaySub")
			stud.GET("displayMark")
		}
	}
	{
		teacher := r.Group("/teacher")
		teacher.GET("/display")
		teacher.POST("/createStud")
		teacher.POST("/createSub")
		teacher.GET("/displaySub")
	}
	return r
}
