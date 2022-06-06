package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/marcellof23/ocbc-practice-day4/core/module"
)

type EmployeeHandler struct {
	employeeUc module.EmployeeUsecase
}

func NewEmployeeHandler(employeeUc module.EmployeeUsecase) *EmployeeHandler {
	return &EmployeeHandler{
		employeeUc: employeeUc,
	}
}

func (hdl *EmployeeHandler) Get(c *gin.Context) {
	Employees, err := hdl.employeeUc.GetEmployees(c)
	if err != nil {
		c.AbortWithStatusJSON(500, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": Employees})
}
