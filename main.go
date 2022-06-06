package main

import (
	"github.com/marcellof23/ocbc-practice-day4/config"
	"github.com/marcellof23/ocbc-practice-day4/core/entity"
	"github.com/marcellof23/ocbc-practice-day4/core/module"
	"github.com/marcellof23/ocbc-practice-day4/handler"
	employeerepository "github.com/marcellof23/ocbc-practice-day4/repository"
	"github.com/marcellof23/ocbc-practice-day4/routes"
)

func main() {

	db := config.Init()
	db.AutoMigrate(&entity.Employee{})

	employeeRepo := employeerepository.New(db)
	employeeUc := module.NewEmployeeUsecase(employeeRepo)
	employeeHdl := handler.NewEmployeeHandler(employeeUc)

	defer db.Close()
	r := routes.SetupRoutes(*employeeHdl)
	r.Run(":8081")
}
