package config

// type Config struct {
// 	pgDB *PostGresConfig
// }


import (
  "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var (
	host     =  "posgres_micky"
	port     =  5432
	user     =  "postgres"
	password =  "k8kwQ8f4A2fjZk3QhyebekRYKK"
	dbname   =  "micky"
)
func Connect() (*gorm.DB, error) {
	db, err := gorm.Open("postgres", "host=" + host + " user=" +user+" dbname=" +dbname+ " password="+password +" sslmode=disable")
	if err != nil {
		return db, err
	}
	return db, nil
}