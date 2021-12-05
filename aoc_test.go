package aoc

import (
	"strconv"
	"testing"
)

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

	GetFile("mocks/day-one.txt")
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
	mock := []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"}
	data := CurrentDiveLocation(mock)

	total := data.depth * data.horizontal
	if total != 900 {
		t.Errorf("CurrentDiveLocation should be 900, instead got: %d", total)
	}

	if data.aim != 10 {
		t.Errorf("Aim should be 10, instead got: %d", data.aim)
	}
}

func TestCheckPowerConsumption(t *testing.T) {
	mock := []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"}
	result := CheckPowerConsumption(mock)

	if result > 0 {
		t.Errorf("Binary Conversion was wrong should be 000 was %d", result)
	}
}

func TestConvertStringToBinary(t *testing.T) {

	mock := []string{"00100", "11110", "10110"}
	result := [][]int64{{1, 0, 0}, {1, 1, 1, 1, 0}, {1, 0, 1, 1, 0}}

	for i, m := range mock {
		num, _ := strconv.ParseInt(m, 2, 32)
		test := _SumArr(ConvertUIntToBinary(num))
		comparison := _SumArr(result[i])
		if comparison != test {
			t.Errorf("Binary Conversion was wrong should be %d was %d", comparison, test)
		}
	}
}

func _SumArr(list []int64) int64 {
	var total int64
	for _, item := range list {
		total += item
	}

	return total
}
