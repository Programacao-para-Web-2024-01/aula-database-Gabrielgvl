package student

import "aula-database/subject"

type StudentService struct {
	repo           *StudentRepository
	subjectService *subject.Service
}

func NewStudentService(repo *StudentRepository, subjectService *subject.Service) *StudentService {
	return &StudentService{repo: repo, subjectService: subjectService}
}

func (s *StudentService) List() ([]Student, error) {
	return s.repo.List()
}

func (s *StudentService) Get(id int) (*Student, error) {
	student, err := s.repo.Get(id)
	if err != nil {
		return nil, err
	}

	student.Subjects, err = s.subjectService.GetByStudentID(id)
	if err != nil {
		return nil, err
	}

	return student, nil
}

func (s *StudentService) Create(student Student) (*Student, error) {
	id, err := s.repo.Create(student)
	if err != nil {
		return nil, err
	}

	student.Id = id

	return &student, nil
}

func (s *StudentService) Update(student Student) error {
	id := int(student.Id)

	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Update(id, student)
}

func (s *StudentService) Delete(id int) error {
	_, err := s.Get(id)
	if err != nil {
		return err
	}

	return s.repo.Delete(id)
}
