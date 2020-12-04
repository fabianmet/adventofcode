package main

import "testing"

func Test_calcIndex(t *testing.T) {
	type args struct {
		ci int
		mi int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcIndex(tt.args.ci, tt.args.mi); got != tt.want {
				t.Errorf("calcIndex() = %v, want %v", got, tt.want)
			}
		})
	}
}
