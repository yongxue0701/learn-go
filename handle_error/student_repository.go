package handle_error

import (
	"database/sql"
	"fmt"

	"github.com/pkg/errors"
)

type StudentRepository struct {
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (s *StudentRepository) GetByID(id int) (*Student, error) {
	// Assume ErrNoRows is thrown
	err := sql.ErrNoRows

	query := fmt.Sprintf("SELECT from student where id = %d", id)
	errMsg := fmt.Sprintf("failed to execute query %s", query)

	return nil, errors.Wrap(err, errMsg)
}
