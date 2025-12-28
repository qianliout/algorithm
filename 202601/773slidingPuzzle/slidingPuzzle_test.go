package main

import (
	"testing"
)

func TestSlidingPuzzle(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  int
	}{
		{
			name:  "Example 1",
			board: [][]int{{1, 2, 3}, {4, 0, 5}},
			want:  1,
		},
		{
			name:  "Example 2",
			board: [][]int{{1, 2, 3}, {5, 4, 0}},
			want:  -1,
		},
		{
			name:  "Example 3",
			board: [][]int{{4, 1, 2}, {5, 0, 3}},
			want:  5,
		},
		{
			name:  "Solved",
			board: [][]int{{1, 2, 3}, {4, 5, 0}},
			want:  0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := slidingPuzzle(tt.board); got != tt.want {
				t.Errorf("slidingPuzzle() = %v, want %v", got, tt.want)
			}
		})
	}
}
