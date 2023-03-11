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
	// create a sequencer
	seq := testcraft.NewSequencer(1)
	// create factory for User with a sequence for ID
	userFactory := testcraft.NewFactory(User{}).Attr(func(u *User) error {
		u.ID = seq.Next()
		return nil
	})
	// create a random User the sequence will be applied to ID
	randUser, err := userFactory.RandomizeWithAttrs()
	if err != nil {
		log.Fatal(err)
	}
	// create a random User the sequence will be applied to ID
	randUser2, err := userFactory.RandomizeWithAttrs()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("ID:", randUser.ID, "Name:", randUser.Name, "Books:", randUser.Books)
	log.Println("ID:", randUser2.ID, "Name:", randUser2.Name, "Books:", randUser2.Books)
}
