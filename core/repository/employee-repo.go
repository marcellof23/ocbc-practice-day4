package repository

import (
	"context"
	"errors"

	"github.com/marcellof23/ocbc-practice-day4/core/entity"
)

var ErrRecordEmployeeNotFound = errors.New("record not found")

type EmployeeRepository interface {
	FindAll(ctx context.Context) ([]entity.Employee, error)
}
