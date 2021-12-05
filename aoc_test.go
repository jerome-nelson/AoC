package aoc

import "testing"

func TestDepthTracker(t *testing.T) {
	mock := []string{"1", "2", "1", "2"}
	depth := DepthTracker(mock)

	if depth != 2 {
		t.Errorf("DepthTracking should be 2, instead got: %d", depth)
	}
}

func TestGetFileParse(t *testing.T) {
	defer func() {
		if recover() == nil {
			return
		}
		t.Error("Something went wrong with files")
	}()

	GetFile("mocks/chapter-one.txt")
}

func TestGetFileError(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("Error wasn't called even though path was wrong")
			return
		}
		t.Log("The call should fail - path was wrong")
	}()

	GetFile("mock/chapter-ones.txt")
}
