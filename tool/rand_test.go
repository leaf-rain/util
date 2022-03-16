package tool

import (
	"fmt"
	"testing"
)

func TestRandStringRunes(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "TestRandStringRunes",
			args: args{10},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RandStringRunes(tt.args.n); got != tt.want {
				t.Errorf("RandStringRunes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRandInt64(t *testing.T) {
	num := RandInt64(0, 2)
	fmt.Println(num)
}
