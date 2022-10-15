package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type DAO struct {
	Connection *sql.DB
}

func ConnectDB() (*sql.DB, error) {
	connString := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", connString)

	if err != nil {
		log.Printf("failed to connect to database: %v", err)
		return &sql.DB{}, err
	}

	return db, nil
}

func (db *DAO) GetStudents() ([]Student, error) {
	return []Student{}, nil
}

func (db *DAO) GetStudentByID(id int64) (Student, error) {
	return Student{}, nil
}

func (db *DAO) UpdateStudent(student Student) error {
	return nil
}

func (db *DAO) DeleteStudent(id int64) error {
	return nil
}
