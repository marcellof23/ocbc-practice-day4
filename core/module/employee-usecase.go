package module

import (
	"context"
	"errors"
	"fmt"

	"github.com/marcellof23/ocbc-practice-day4/core/entity"
	"github.com/marcellof23/ocbc-practice-day4/core/repository"
)

type EmployeeUsecase interface {
	GetEmployees(ctx context.Context) ([]entity.Employee, error)
}

type employeeUsecase struct {
	employeeRepo repository.EmployeeRepository
}

// NewEmployeeUsecase use for initiate new employee usecase
func NewEmployeeUsecase(repo repository.EmployeeRepository) EmployeeUsecase {
	return &employeeUsecase{repo}
}

var ErrEmployeeNotFound = errors.New("employee not found")

func (em *employeeUsecase) GetEmployees(ctx context.Context) ([]entity.Employee, error) {
	data, err := em.employeeRepo.FindAll(ctx)
	if err != nil {
		if errors.Is(err, repository.ErrRecordEmployeeNotFound) {
			return nil, fmt.Errorf("%w.", ErrEmployeeNotFound)
		}
		return nil, fmt.Errorf("%w: %v", ErrEmployeeNotFound, err)
	}
	return data, nil
}
