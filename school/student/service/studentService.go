package service

import (
	"errors"

	"github.com/sreeks87/school/student/domain"
)

type StudentSvc struct {
	Repo domain.Repository
}

func NewStudentSvc(r domain.Repository) domain.Service {
	return &StudentSvc{
		Repo: r,
	}
}

func (s *StudentSvc) Register(std *domain.Student) (*domain.Student, error) {
	res, e := s.Repo.Save(std)
	if e != nil {
		return nil, e
	}
	return res, nil
}

func (s *StudentSvc) Validate(std *domain.Student) error {
	if std.Age < 15 || std.Age > 25 {
		return errors.New("student not eligible to enroll")
	}
	return nil
}

func (s *StudentSvc) Fetch(id string) (*domain.Student, error) {
	return s.Repo.Get(id)
}
