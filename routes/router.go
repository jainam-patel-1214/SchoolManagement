package routes

import "github.com/gin-gonic/gin"

func InitializeRouter() *gin.Engine {
	r := gin.Default()
	{
		stud := r.Group("/student")
		stud.GET("/display")
	}
	return r
}
