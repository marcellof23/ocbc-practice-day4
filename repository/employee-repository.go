package employeerepository

import (
	"context"

	"github.com/jinzhu/gorm"
	"github.com/marcellof23/ocbc-practice-day4/core/entity"
	repository_intf "github.com/marcellof23/ocbc-practice-day4/core/repository"
)

type repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) repository_intf.EmployeeRepository {
	return &repository{db}
}

func (r *repository) FindAll(ctx context.Context) ([]entity.Employee, error) {
	employees := []entity.Employee{}
	res := r.db.Find(&employees)
	print(res.RowsAffected)

	// if err != nil {
	// 	if errors.Is(err, gorm.ErrRecordNotFound) {
	// 		return nil, repository_intf.ErrRecordEmployeeNotFound
	// 	}
	// 	return nil, err
	// }

	return employees, nil
}
