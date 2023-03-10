package testcraft

import (
	"database/sql"
	"testing"
	"time"
)

func Test_randomize(t *testing.T) {
	type Book struct {
		Title string
	}
	type alias string
	type User struct {
		Name2     string
		Name      *string
		Books     []string
		BookPtr   []*string
		Books2    []Book
		Books2Ptr []*Book
		Books3    Book
		Books3Ptr *Book
		Alias     alias
		SqlVal    sql.NullBool
		t         time.Time
		UpdatedAt *time.Time
	}
	type args[T any] struct {
		f      T
		valuer Valuer
	}
	type testCase[T any] struct {
		name    string
		args    args[T]
		wantErr bool
	}
	tests := []testCase[User]{
		{name: "string", args: args[User]{f: User{}, valuer: defaultValuer()}, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := randomize[User](tt.args.f, tt.args.valuer)
			if (err != nil) != tt.wantErr {
				t.Errorf("randomize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}
