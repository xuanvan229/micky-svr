package db

import (
	"fmt"
)

const (
	host     = "posgres_micky"
	port     = 5432
	user     = "postgres"
	password = "k8kwQ8f4A2fjZk3QhyebekRYKK"
	dbname   = "micky"
)

func DbInfo() string {
	return fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
}
