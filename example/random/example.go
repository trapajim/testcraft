package main

import (
	"github.com/trapajim/testcraft"
	"log"
)

type User struct {
	ID    int
	Name  string
	Books []string
}

func main() {
	// create a new factory for the User struct
	userFactory := testcraft.NewFactory(User{})
	// get a User with random values
	randUser, err := userFactory.Randomize()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ID:", randUser.ID, "Name:", randUser.Name, "Books:", randUser.Books)
}
