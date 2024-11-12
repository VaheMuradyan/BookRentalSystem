package main

import "github.com/VaheMuradyan/BookRentalSystem/Database"

func init() {
	Database.ConnectDatabase()
	Database.DBMigrate()
}

func main() {

}
