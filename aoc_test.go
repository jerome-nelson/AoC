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

func TestCurrentDiveLocation(t *testing.T) {
	mock := []string{"forward 1", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	data := CurrentDiveLocation(mock)

	if data != 110 {
		t.Errorf("CurrentDiveLocation should be 110, instead got: %d", data)
	}
}
