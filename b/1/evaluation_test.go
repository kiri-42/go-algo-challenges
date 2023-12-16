package main

import (
	"testing"
)

// TestCalculateOptimalPathByChatGPT tests the CalculateOptimalPath function.
func TestCalculateOptimalPathByChatGPT(t *testing.T) {
	var tests = []struct {
		name      string
		grid      [][]int
		startX    int
		startY    int
		endX      int
		endY      int
		wantCost  int
		wantError bool
	}{
		{
			name:      "Normal Pathfinding",
			grid:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			startX:    0,
			startY:    0,
			endX:      2,
			endY:      2,
			wantCost:  20,
			wantError: false,
		},
		{
			name:      "Shortest Path",
			grid:      [][]int{{1, 1, 1}, {100, 100, 100}, {1, 1, 1}},
			startX:    0,
			startY:    0,
			endX:      2,
			endY:      2,
			wantCost:  103,
			wantError: false,
		},
		{
			name:      "Invalid Grid Value (Negative)",
			grid:      [][]int{{-1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			startX:    0,
			startY:    0,
			endX:      2,
			endY:      2,
			wantCost:  0,
			wantError: true,
		},
		{
			name:      "Invalid Start Coordinates (Out of Bounds)",
			grid:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			startX:    -1,
			startY:    0,
			endX:      2,
			endY:      2,
			wantCost:  0,
			wantError: true,
		},
		{
			name:      "Invalid End Coordinates (Out of Bounds)",
			grid:      [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}},
			startX:    0,
			startY:    0,
			endX:      3,
			endY:      3,
			wantCost:  0,
			wantError: true,
		},
		{
			name:      "Non-Square Grid",
			grid:      [][]int{{1, 2, 3}, {4, 5, 6}},
			startX:    0,
			startY:    0,
			endX:      1,
			endY:      2,
			wantCost:  0,
			wantError: true,
		},
		{
			name:      "Empty Grid",
			grid:      [][]int{},
			startX:    0,
			startY:    0,
			endX:      0,
			endY:      0,
			wantCost:  0,
			wantError: true,
		},
		// 他に考慮すべきテストケースがあればここに追加
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCost, err := CalculateOptimalPath(tt.grid, tt.startX, tt.startY, tt.endX, tt.endY)
			if (err != nil) != tt.wantError {
				t.Errorf("CalculateOptimalPath() error = %v, wantError %v", err, tt.wantError)
				return
			}
			if gotCost != tt.wantCost {
				t.Errorf("CalculateOptimalPath() = %v, want %v", gotCost, tt.wantCost)
			}
		})
	}
}
