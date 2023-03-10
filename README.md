# TestCraft
TestCraft is a Go library for filling your structs with test data. 
You can either define attributes for your structs or randomize them. 

## Installation
To use TestCraft, you can install it using the go get command:

```bash
go get github.com/trapajim/testcraft
```


## Usage
Example of how to use TestCraft:

```go
package main

import (
	"fmt"
	"github.com/trapajim/go-fixtures/datagen"
	"github.com/trapajim/testcraft"
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
	userSeq :=  testcraft.NewSequencer(1)

	// Define attributes for the User struct
	userFactory.Attr(func(u *User) error {
		u.ID = userSeq.Next()
		u.Books =  testcraft.Multiple(5, func(i int) string {
			return fmt.Sprintf("book %d", i)
		})
		u.Name = datagen.AlphanumericBetween(5,10)
		return nil
	})

	// Build a new User struct
	user1, err := userFactory.Build()
	fmt.Println("ID:", user1.ID, "Name:", user1.Name)
	if err != nil {
		// Handle the error
	}
	// Build a new User struct MustBuild panics on error
	user2 := userFactory.MustBuild()
	fmt.Println("ID:", user2.ID, "Name:", user2.Name)
	
	
}
```

Output:
```bash
ID: 1 Name: yUKW9 Books: [book 0 book 1 book 2 book 3 book 4]
ID: 2 Name: jXZ816KH [book 0 book 1 book 2 book 3 book 4]
```

TestCraft can create random data for your structs
    
```go
randUser, err := userFactory.Randomize()
if err != nil {
// Handle the error
}
fmt.Println("ID:", randUser.ID, "Name:", randUser.Name, "Books:", randUser.Books)
```
Output:
```bash
ID: 32 Name: agree pedal Books: [cool egg fish apple advise rich]
```

The randomizer has a set of default rules for various types:

| Type | Rule |
| --- | --- |
| string | random two words |
| int | random number between 0 and 100 |
| float | random number between 0 and 100 |
| bool | random bool |
| time.Time | random time between 1970 and 2070 |



## Contributing
If you find a bug or want to suggest a new feature for testcraft, please open an issue on GitHub or submit a pull request. We welcome contributions from the community.

