package main

import (
	"fmt"
	"github.com/trapajim/testcraft"

	"github.com/trapajim/testcraft/datagen"
)

type User struct {
	ID    int
	Name  string
	Books []string
}

func main() {
	// Create a new factory for the User struct
	userFactory := testcraft.NewFactory(User{})

	// Create a sequencer to generate unique IDs
	userSeq := testcraft.NewSequencer(1)

	// Define attributes for the User struct
	userFactory.Attr(func(u *User) error {
		u.ID = userSeq.Next()
		u.Books = testcraft.Multiple(5, func(i int) string {
			return fmt.Sprintf("book %d", i)
		})
		u.Name = datagen.AlphanumericBetween(5, 10)
		return nil
	})

	// Build a new User struct
	user1, err := userFactory.Build()
	fmt.Println("ID:", user1.ID, "Name:", user1.Name, "Books:", user1.Books)
	if err != nil {
		// Handle the error
	}
	// Build a new User struct MustBuild panics on error
	user2 := userFactory.MustBuild()
	fmt.Println("ID:", user2.ID, "Name:", user2.Name, "Books:", user2.Books)

	randUser, err := userFactory.Randomize()
	if err != nil {
		// Handle the error
	}
	fmt.Println("ID:", randUser.ID, "Name:", randUser.Name, "Books:", randUser.Books)

}
