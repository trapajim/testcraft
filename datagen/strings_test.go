package datagen

import (
	"strings"
	"testing"
)

func TestAlphanumeric(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{"expect string to have a length of 1", args{1}},
		{"expect string to have a length of 100", args{100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Alphanumeric(tt.args.length); len(got) != tt.args.length {
				t.Errorf("Alphanumeric() = %v, want %v", len(got), tt.args.length)
			}
		})
	}
}

func TestWords(t *testing.T) {
	type args struct {
		length int
	}
	tests := []struct {
		name string
		args args
	}{
		{"expect string to have 1 word", args{1}},
		{"expect string to have 100 words", args{100}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Words(tt.args.length); len(strings.Fields(got)) != tt.args.length {
				t.Errorf("Words() = %v, want %v", len(strings.Fields(got)), tt.args.length)
			}
		})
	}
}
