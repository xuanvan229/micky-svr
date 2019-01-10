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

func DbInfo() string{
	return fmt.Sprintf("host=%s port=%d user=%s "+
    "password=%s dbname=%s sslmode=disable",
    host, port, user, password, dbname)
}

// const (
// 	host       = "mongo_micky:27017"
// 	database   = "micky_user"
// 	username   = "admin"
// 	mechanism  = "SCRAM-SHA-1"
// 	password   = "k8kwQ8f4A2fjZk3QhyebekRYKK"
// )

// func ConnectToCol(col_name string) (*mgo.Collection, error) {
// 	info_db := &mgo.DialInfo{
// 		Addrs:    []string{host},
// 		Timeout:  20 * time.Second,
// 		Username: username,
// 		Password: password,
// 		Mechanism: mechanism,
// 		Database: database,
// 	}
// 	fmt.Println("connect to section")
// 	section, err := mgo.DialWithInfo(info_db)
// 	fmt.Println("finish to section", section)
// 	if err != nil {
// 		fmt.Println(err)
// 		return nil , err
// 	}

// 	col := section.DB(database).C(col_name)
// 	return col, nil
// }