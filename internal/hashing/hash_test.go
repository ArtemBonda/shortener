package hashing

import "testing"

func TestHashURLAddr(t *testing.T) {
	type args struct {
		something string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{name: "Test #1", args: args{something: "https://golang.org"}, want: HashURLAddr([]byte("https://golang.org"))},
		{name: "Empty String", args: args{something: ""}, want: HashURLAddr([]byte(""))},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			if got := HashURLAddr([]byte(tt.args.something)); got != tt.want {
				t.Errorf("HashURLAddr() = %v, want %v", got, tt.want)
			}
		})
	}
}
