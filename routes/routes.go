package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/marcellof23/ocbc-practice-day4/handler"
)

func SetupRoutes(employeeHdl handler.EmployeeHandler) *gin.Engine {
	r := gin.Default()

	r.GET("/employees", employeeHdl.Get)
	// r.POST("/employees", handler.CreateEmployee)
	// r.GET("/employees/:id", handler.FindEmployee)
	// r.PATCH("/employees/:id", handler.UpdateEmployee)
	// r.DELETE("employees/:id", handler.DeleteEmployee)
	return r
}
