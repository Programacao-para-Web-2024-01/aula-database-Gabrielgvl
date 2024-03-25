package db

import (
	"database/sql"
	"sync"
)

type Student struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

type StudentRepository struct {
	db *sql.DB
	m  map[int]Student
	mu *sync.RWMutex
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (sr *StudentRepository) List() ([]Student, error) {
	rows, err := sr.db.Query(`SELECT id, name, age, email, phone FROM students`)
	if err != nil {
		return nil, err
	}

	var students []Student

	for rows.Next() {
		var student Student
		err = rows.Scan(&student.Id, &student.Name, &student.Age, &student.Email, &student.Phone)
		if err != nil {
			return nil, err
		}

		students = append(students, student)
	}

	rows.Close()

	return students, nil
}

func (sr *StudentRepository) Get(id int) (*Student, error) {
	row := sr.db.QueryRow(`
		SELECT id, name, age, email, phone
		FROM students
		WHERE id = ?`, id)

	var student Student
	err := row.Scan(&student.Id, &student.Name, &student.Age, &student.Email, &student.Phone)
	if err != nil {
		return nil, err
	}

	return &student, nil
}

func (sr *StudentRepository) Create(student Student) (int, error) {
	sr.db.Exec()
}

func (sr *StudentRepository) Update(id int, student Student) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	sr.m[id] = student

	return nil
}

func (sr *StudentRepository) Delete(id int) error {
	sr.mu.Lock()
	defer sr.mu.Unlock()

	delete(sr.m, id)

	return nil
}
