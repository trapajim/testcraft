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
