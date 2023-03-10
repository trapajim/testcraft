package testcraft

import (
	"fmt"
	"reflect"
	"testing"
)

func TestMultiple(t *testing.T) {
	// Test with integers
	expectedInts := []int{0, 1, 2, 3, 4}
	actualInts := Multiple(5, func(i int) int { return i })
	if !reflect.DeepEqual(actualInts, expectedInts) {
		t.Errorf("multiple num returned %v, expected %v", actualInts, expectedInts)
	}

	// Test with strings
	expectedStrings := []string{"foo_0", "foo_1", "foo_2", "foo_3", "foo_4"}
	actualStrings := Multiple(5, func(i int) string { return fmt.Sprintf("foo_%d", i) })
	if !reflect.DeepEqual(actualStrings, expectedStrings) {
		t.Errorf("Multiple strings returned %v, expected %v", actualStrings, expectedStrings)
	}

	// Test with custom struct
	type person struct {
		Name string
		Age  int
	}
	expectedPeople := []person{
		{"Person1", 25},
		{"Person2", 30},
		{"Person3", 35},
	}
	actualPeople := Multiple(3, func(i int) person {
		return person{Name: fmt.Sprintf("Person%d", i+1), Age: 25 + (i * 5)}
	})
	if !reflect.DeepEqual(actualPeople, expectedPeople) {
		t.Errorf("multiple struct returned %v, expected %v", actualPeople, expectedPeople)
	}
}
