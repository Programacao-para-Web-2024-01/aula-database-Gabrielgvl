package student

import (
	"encoding/json"
	"net/http"
	"strconv"
)

type StudentController struct {
	service *StudentService
}

func NewStudentController(service *StudentService) *StudentController {
	return &StudentController{service: service}
}

func (s *StudentController) List(w http.ResponseWriter, req *http.Request) {
	students, err := s.service.List()
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(students)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *StudentController) Get(w http.ResponseWriter, req *http.Request) {
	// Input
	idRaw := req.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// processamento
	student, err := s.service.Get(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// output
	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *StudentController) Create(w http.ResponseWriter, req *http.Request) {
	// Leitura do corpo (INPUT)
	var student Student
	err := json.NewDecoder(req.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	// Lógica da função/Algoritmo
	newStudent, err := s.service.Create(student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// Output / Resposta
	err = json.NewEncoder(w).Encode(newStudent)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

func (s *StudentController) Update(w http.ResponseWriter, req *http.Request) {
	// Input
	idRaw := req.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	var student Student
	err = json.NewDecoder(req.Body).Decode(&student)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	student.Id = int64(id)

	err = s.service.Update(student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	err = json.NewEncoder(w).Encode(student)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

}

func (s *StudentController) Delete(w http.ResponseWriter, req *http.Request) {
	// Input
	idRaw := req.PathValue("id")

	id, err := strconv.Atoi(idRaw)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	err = s.service.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	w.WriteHeader(204)
}
