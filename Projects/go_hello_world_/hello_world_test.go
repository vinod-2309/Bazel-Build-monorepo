package helloWorld

import "testing"

func TestHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test with empty string",
			input: "",
			want: "Empty string",
		},
		{
			name: "Test with special characters",
			input:"Vinod"
			want: "Hello, World!",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HelloWorld(tt.input); got != tt.want {
				t.Errorf("HelloWorld() = %v, want %v", got, tt.want)
			}
		})
	}
}
