package testcraft

import (
	"errors"
	"reflect"
	"testing"
)

func TestNewFactory(t *testing.T) {
	type User struct {
		Name string
	}
	obj := User{}
	factory := NewFactory(obj)

	// Assert that the object is set correctly
	if factory.object != obj {
		t.Errorf("NewFactory returned factory with object %v, expected %v", factory.object, obj)
	}

	// Assert that the sequence map is initialized
	if factory.sequence == nil {
		t.Errorf("NewFactory returned factory with nil sequence map")
	}
}

func TestFactoryAttr(t *testing.T) {
	type User struct {
		Name string
	}
	obj := User{}
	factory := NewFactory(obj)

	// Test adding a single attribute generator

	attrGen1 := func(instance *User) error {
		return nil
	}
	factory.Attr(attrGen1)

	// Assert that the attribute generator is added
	if len(factory.attrsGen) != 1 {
		t.Errorf("Factory.Attr did not add attribute generator correctly")
	}

	factory.Attr(attrGen1)

	// Assert that both attribute generators are added in the correct order
	if len(factory.attrsGen) != 2 {
		t.Errorf("Factory.Attr did not add attribute generators correctly")
	}
}

func TestFactoryBuild(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	tests := []struct {
		obj        Person
		attrsGens  []AttrGenerator[Person]
		expected   Person
		expectFail bool
	}{
		{
			obj: Person{
				Name: "Alice",
				Age:  25,
			},
			attrsGens: []AttrGenerator[Person]{
				func(p *Person) error {
					p.Name = "Bob"
					return nil
				},
				func(p *Person) error {
					p.Age = 25
					return nil
				},
			},
			expected: Person{
				Name: "Bob",
				Age:  25,
			},
			expectFail: false,
		},
		{
			obj: Person{
				Name: "Charlie",
				Age:  30,
			},
			attrsGens: []AttrGenerator[Person]{
				func(p *Person) error {
					return errors.New("failed to generate attributes")
				},
			},
			expected:   Person{},
			expectFail: true,
		},
	}

	for i, test := range tests {

		factory := NewFactory(test.obj).Attr(test.attrsGens...)
		result, err := factory.build()
		if test.expectFail {
			if err == nil {
				t.Errorf("Test case %d: expected error, but got nil", i)
			}
		} else {
			if err != nil {
				t.Errorf("Test case %d: unexpected error: %v", i, err)
			}
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Test case %d: unexpected result: got %v, expected %v", i, result, test.expected)
			}
		}
	}
}

func TestFactory_Randomize(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}
	tests := []struct {
		obj        Person
		expectFail bool
	}{
		{
			obj:        Person{},
			expectFail: false,
		},
	}
	for i, test := range tests {
		_, err := NewFactory(test.obj).Randomize()

		if test.expectFail {
			if err == nil {
				t.Errorf("Test case %d: expected error, but got nil", i)
			}
		} else {
			if err != nil {
				t.Errorf("Test case %d: unexpected error: %v", i, err)
			}
		}
	}
}

func TestFactory_RandomizeWithAttrs(t *testing.T) {
	type Person struct {
		ID   int
		Name string
		Age  int
	}
	seq := NewSequencer(1)
	pFactory := NewFactory(Person{}).Attr(func(p *Person) error {
		p.ID = seq.Next()
		return nil
	})
	p1 := pFactory.MustRandomizeWithAttrs()
	p2 := pFactory.MustRandomizeWithAttrs()
	if p1.ID != 1 {
		t.Errorf("Expected p1.ID to be 1, got %d", p1.ID)
	}
	if p2.ID != 2 {
		t.Errorf("Expected p2.ID to be 2, got %d", p2.ID)
	}

	p3, err := pFactory.RandomizeWithAttrs()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	p4, err := pFactory.RandomizeWithAttrs()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if p3.ID != 3 {
		t.Errorf("Expected p3.ID to be 3, got %d", p3.ID)
	}
	if p4.ID != 4 {
		t.Errorf("Expected p4.ID to be 4, got %d", p4.ID)
	}
}

func TestFactory_With(t *testing.T) {
	type Person struct {
		ID   int
		Name string
		Age  int
	}
	personTraits := struct {
		Senior AttrGenerator[Person]
		Junior AttrGenerator[Person]
		Steve  AttrGenerator[Person]
	}{
		Senior: func(p *Person) error {
			p.Age = 65
			return nil
		},
		Junior: func(p *Person) error {
			p.Age = 18
			return nil
		},
		Steve: func(p *Person) error {
			p.Name = "Steve"
			return nil
		},
	}
	factory := NewFactory(Person{}).Attr(func(p *Person) error {
		p.Name = "John Doe"
		return nil
	})
	t.Run("A single trait", func(t *testing.T) {
		factory := factory.With(personTraits.Senior)
		person1 := factory.MustBuild()
		if person1.Age != 65 {
			t.Errorf("Expected person1.Age to be 65, got %d", person1.Age)
		}
		if person1.Name != "John Doe" {
			t.Errorf("Expected person1.Name to be John Doe, got %s", person1.Name)
		}
	})
	t.Run("with two traits", func(t *testing.T) {
		factory := factory.With(personTraits.Junior, personTraits.Steve)
		person1 := factory.MustBuild()
		if person1.Age != 18 {
			t.Errorf("Expected person1.Age to be 18, got %d", person1.Age)
		}
		if person1.Name != "Steve" {
			t.Errorf("Expected person1.Name to be Steve, got %s", person1.Name)
		}
	})
	t.Run("default factory should not be affected", func(t *testing.T) {
		person1 := factory.MustBuild()
		if person1.Name != "John Doe" {
			t.Errorf("Expected person1.Name to be John Doe, got %s", person1.Name)
		}
		if person1.Age != 0 {
			t.Errorf("Expected person1.Age to be 0, got %d", person1.Age)
		}
	})
}
