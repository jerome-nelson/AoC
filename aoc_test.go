package aoc

import "testing"

func TestDepthTracker(t *testing.T) {
	mock := []int{1, 2, 1, 2}
	depth := DepthTracker(mock)

	if depth != 2 {
		t.Errorf("DepthTracking should be 2, instead got: %d", depth)
	}
}
