package main

import "testing"

func Test_main(t *testing.T) {
	tests := []struct {
		name string
	}{
		// Just so we see the example code works
		{"Example"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			main()
		})
	}
}
