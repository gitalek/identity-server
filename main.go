package main

import (
	"server/models/user"
)

func main() {
	db := initDB()
	users := user.NewUsers(db, "users")

}
