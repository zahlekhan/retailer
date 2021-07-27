package Routes

import (
	//"awesomeProject1/LearnGo/Day3/Exercises/Exercise2/Controllers"
	"github.com/gin-gonic/gin"
)

//SetupRouter ... Configure routes
func SetupRouter() *gin.Engine {
	r := gin.Default()
	//grp1 := r.Group("/api")
	//{
	//	grp1.GET("student", Controllers.GetStudents)
	//	grp1.POST("student", Controllers.CreateStudent)
	//	grp1.GET("student/:id", Controllers.GetStudentByID)
	//	grp1.PUT("student/:id", Controllers.UpdateStudent)
	//	grp1.DELETE("student/:id", Controllers.DeleteStudent)
	//}
	return r
}
