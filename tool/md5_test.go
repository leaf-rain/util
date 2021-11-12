package tool

import "testing"

func TestGetMD5Encode(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "TestGetMD5Encode",
			args: args{"123456"},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetMD5Encode(tt.args.data); got != tt.want {
				t.Errorf("GetMD5Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}
