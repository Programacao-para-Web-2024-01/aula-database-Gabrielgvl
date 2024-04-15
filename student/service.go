package student

type StudentService struct {
	repo *StudentRepository
}

func NewStudentService(repo *StudentRepository) *StudentService {
	return &StudentService{repo: repo}
}

func (s *StudentService) List() ([]Student, error) {
	return s.repo.List()
}

func (s *StudentService) Get(id int) (*Student, error) {
	return s.repo.Get(id)
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
