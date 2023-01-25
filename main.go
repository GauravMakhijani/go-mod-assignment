package main

import (
	"assignment/db"
	"fmt"

	"github.com/bxcodec/faker/v4"
)

func main() {
	user := db.User{faker.FirstName(), faker.LastName(), faker.Email(), faker.Date(), faker.Phonenumber()}
	fmt.Print(user)
}
