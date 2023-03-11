package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/trapajim/testcraft"
)

// You can combine multiple factories to create more complex structs.
// You can also use custom types in your structs.

type Score int
type Book struct {
	Title string
	Score Score
}
type User struct {
	ID    int
	Name  string
	Books []Book
}

func main() {
	// Create a new factory for the User struct
	userFactory := testcraft.NewFactory(User{})
	// Create a new factory for the Book struct
	bookFactory := testcraft.NewFactory(Book{})
	// Create a sequencer to generate unique IDs
	userSeq := testcraft.NewSequencer(1)

	// Define attributes for the User struct
	userFactory.Attr(func(u *User) error {
		u.ID = userSeq.Next()
		// create a slice of 5 random Books
		u.Books = testcraft.Multiple(5, func(i int) Book {
			return bookFactory.MustRandomize()
		})
		u.Name = "name"
		return nil
	})

	// Build a new User struct
	user1, err := userFactory.Build()
	if err != nil {
		panic(err)
	}
	// Build a new User struct MustBuild panics on error
	user2 := userFactory.MustBuild()
	PrettyPrint(user1)
	PrettyPrint(user2)

}

func PrettyPrint(data interface{}) {
	var buffer bytes.Buffer
	enc := json.NewEncoder(&buffer)
	enc.SetIndent("", "  ")
	if err := enc.Encode(data); err != nil {
		panic(err)
	}
	fmt.Println(buffer.String())
}
