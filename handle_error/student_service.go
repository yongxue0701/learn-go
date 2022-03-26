package handle_error

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type StudentService struct {
}

func NewStudentService() *StudentService {
	return &StudentService{}
}

func (s *StudentService) GetByID(id int) (*Student, error) {
	repo := NewStudentRepository()
	student, err := repo.GetByID(id)

	switch errors.Cause(err) {
	case sql.ErrNoRows:
		errMsg := fmt.Sprintf("student with id %d is not found", id)
		return nil, errors.WithMessage(err, errMsg)
	}

	return student, nil
}
