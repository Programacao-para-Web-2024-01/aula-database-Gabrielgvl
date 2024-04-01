package main

import (
	"aula-database/student"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
	"net/http"
)

func main() {
	if err := createServer(); err != nil {
		log.Panic(err)
	}
}

func connectDB() *sql.DB {
	config := mysql.NewConfig()
	config.User = "root"
	config.Passwd = "uniceub"
	config.DBName = "web"
	conn, err := mysql.NewConnector(config)
	if err != nil {
		panic(err)
	}
	return sql.OpenDB(conn)
}

func createServer() error {
	db := connectDB()

	studentRepository := student.NewStudentRepository(db)
	studentService := student.NewStudentService(studentRepository)
	studentController := student.NewStudentController(studentService)

	mux := http.NewServeMux()

	mux.HandleFunc("GET /students/", studentController.List)
	mux.HandleFunc("POST /students/", studentController.Create)
	mux.HandleFunc("GET /students/{id}", studentController.Get)
	mux.HandleFunc("PUT /students/{id}", studentController.Update)
	mux.HandleFunc("DELETE /students/{id}", studentController.Delete)

	return http.ListenAndServe("localhost:8080", mux)
}
