package database

import "fmt"

var connection string

func init() {
	fmt.Println("ini dipanggil")
	connection = "Mysql"
}

func getData() string {
	return connection
}
