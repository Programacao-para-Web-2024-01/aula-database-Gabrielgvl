package student

import (
	"database/sql"
	"encoding/json"
)

type StudentRepository struct {
	db *sql.DB
}

func NewStudentRepository(db *sql.DB) *StudentRepository {
	return &StudentRepository{
		db: db,
	}
}

func (sr *StudentRepository) List() ([]Student, error) {
	rows, err := sr.db.Query(`
		SELECT st.id,
		       st.name,
		       st.age,
		       st.email,
		       st.phone, 
		       IF(COUNT(su.id) > 0, JSON_ARRAYAGG(su.name), NULL) subjects
		FROM students st
         	LEFT JOIN students_subjects ss ON st.id = ss.student_id
         	LEFT JOIN subjects su ON ss.subject_id = su.id
		GROUP BY st.id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var students []Student

	for rows.Next() {
		var student Student
		var subjects []byte
		err = rows.Scan(
			&student.Id,
			&student.Name,
			&student.Age,
			&student.Email,
			&student.Phone,
			&subjects,
		)
		if err != nil {
			return nil, err
		}

		if subjects != nil {
			err = json.Unmarshal(subjects, &student.SubjectsName)
			if err != nil {
				return nil, err
			}
		}

		students = append(students, student)
	}

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

func (sr *StudentRepository) Create(student Student) (int64, error) {
	result, err := sr.db.Exec(`INSERT INTO students(name, age, email, phone)
					  VALUES (?, ?, ?, ?)`,
		student.Name, student.Age, student.Email, student.Phone)

	if err != nil {
		return 0, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (sr *StudentRepository) Update(id int, student Student) error {
	_, err := sr.db.Exec(`UPDATE students
						SET name=?,
						    age=?,
						    email=?,
						    phone=?
						WHERE id=?`,
		student.Name, student.Age, student.Email, student.Phone, id)

	if err != nil {
		return err
	}

	return nil
}

func (sr *StudentRepository) Delete(id int) error {
	_, err := sr.db.Exec(`DELETE
							FROM students
							WHERE id = ?`, id)

	if err != nil {
		return err
	}

	return nil
}
