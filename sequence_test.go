package testcraft

import (
	"testing"
)

func TestNewSequencer(t *testing.T) {
	type args[T Number] struct {
		start T
	}
	type testCase[T Number] struct {
		name string
		args args[T]
		want *Sequence[T]
	}
	seqInt := NewSequencer(0)
	assertEqual(t, seqInt.Next(), 0)
	assertEqual(t, seqInt.Next(), 1)

	seqFl := NewSequencer(0.0)
	assertEqual(t, seqFl.Next(), 0.0)
	assertEqual(t, seqFl.Next(), 1.0)

	seqInc := NewSequencer(0)
	seqInc.SetIncrement(2)
	assertEqual(t, seqInc.Next(), 0)
	assertEqual(t, seqInc.Next(), 2)

	seqFlInc := NewSequencer(0.0)
	seqFlInc.SetIncrement(3.0)
	assertEqual(t, seqFlInc.Next(), 0.0)
	assertEqual(t, seqFlInc.Next(), 3.0)
}

func assertEqual(t *testing.T, got, want interface{}) {
	t.Helper()
	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
