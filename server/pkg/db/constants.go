package db

import "os"

const (
	port   = 5432
	dbname = "bookings"
)

var (
	host     = os.Getenv("DB_HOST")
	user     = os.Getenv("DB_USERNAME")
	password = os.Getenv("DB_PASSWORD")
)
