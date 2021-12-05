package aoc

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"strconv"
	"strings"
)

// Day 3
// 1. Need two vars to hold gamma and epsilon data
// 		1a. Check the bit length of first string (NICE TO HAVE)
//	 		i. Calculate 2n x 2n x ... to get decimal max
//			ii. If more than 255 (uint16) else uint8 (could expand on this but only to cover sample and input given)
//			iii. Create array with this max added
// 2. Read each line as string and store them in array
// 		i. Can we optimise this quickly? (alot of records)
//      ii. Instead of storing as a string of bits -> can convert directly to decimal/bit and store like that?
//		iii. Each string character in Go is natively UTF-8 so approx around 1byte per character (minimum)
//		iv. Actually takes around 48bytes per character (https://go.dev/play/p/sWxtrSQbxmi.go?download=true)
// 3. Bitwise NOT on gamma to get epsilon and convert to get decimal
// 4. Multiply result
func CheckPowerConsumption(input []string) int {
	var (
		linemax int
		count   []int64
		gamma   []int64
	)

	for _, line := range input {

		if linemax == 0 {
			linemax = len(line)
		}

		num, err := strconv.ParseInt(line, 2, 32)
		check(err)
		binary := ConvertUIntToBinary(num)
		remainder := len(line) - len(binary)

		if remainder != 0 {
			prepend := make([]int64, remainder)
			binary = append(prepend, binary...)
		}

		count = append(count, binary...)
	}

	gamma = make([]int64, linemax)

	for i := range gamma {
		result := make(map[int]int)
		result[0] = 0
		result[1] = 0
		offset := linemax - i

		for j := 0; j < len(count); j += offset {
			if count[j] == 0 {
				result[0] += 1
				continue
			}
			result[1] += 1
		}

		if result[0] > result[1] {
			gamma[i] = 0
			continue
		}

		gamma[i] = 1
	}

	return 0
}

func ConvertUIntToBinary(num int64) []int64 {
	var (
		output []int64
		binary []int64
	)

	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}

	for i := len(binary) - 1; i >= 0; i-- {
		output = append(output, binary[i])
	}
	return output
}

type DiveInstruction int8
type DiveCoords struct {
	horizontal int
	depth      int
	aim        int
}

const (
	Undefined DiveInstruction = iota
	Up        DiveInstruction = iota
	Down      DiveInstruction = iota
	Right     DiveInstruction = iota
)

func getDiveInstruction(direction string) (DiveInstruction, error) {
	var err error

	switch direction {
	case "up":
		return Up, err
	case "down":
		return Down, err
	case "forward":
		return Right, err
	}
	return Undefined, errors.New("dive instruction given is not correct")
}

// Points for Day 2
// 1. Starts from Zero Integers
// 2. Essentially a graph with x, y co-ords
// 3. Forward (left) increments horizontal. Up/Down increments/decrements the vertical number

// Part 2
// 1. Aim should increment/decrement on Down/Up
// 2. On forward it should multiply current aim by number of places forward
func CurrentDiveLocation(instruction []string) DiveCoords {

	currentLocation := DiveCoords{
		horizontal: 0,
		depth:      0,
		aim:        0,
	}

	for _, order := range instruction {
		destructure := strings.Split(order, " ")
		direction, err := getDiveInstruction(destructure[0])
		check(err)
		number, err := strconv.Atoi(destructure[1])
		check(err)

		if direction == Right {
			addedDepth := currentLocation.aim * number
			currentLocation.horizontal += number
			currentLocation.depth += addedDepth
		}

		if direction == Up {
			currentLocation.aim -= number
		}

		if direction == Down {
			currentLocation.aim += number
		}
	}

	return currentLocation
}

// Points for Day 1
// 1. Need a way to send input to script
// 2. Data is given as a Txt File - so need to scrape that
// 3. Need to keep a count of all increased times
func DepthTracker(listOfDepths []string) int {
	tracker := 0
	for order, depth := range listOfDepths {
		if order == 0 {
			continue
		}

		currentDepth, err := strconv.Atoi(depth)
		check(err)
		previousDepth, err := strconv.Atoi(listOfDepths[order-1])
		check(err)

		if previousDepth < currentDepth {
			tracker += 1
		}

	}
	return tracker
}

func GetFile(path string) []string {

	var (
		data   *os.File
		err    error
		part   []byte
		prefix bool
		lines  []string
	)
	data, err = os.Open(path)
	check(err)
	defer data.Close()

	reader := bufio.NewReader(data)
	buffer := bytes.NewBuffer(make([]byte, 0))

	for {
		if part, prefix, err = reader.ReadLine(); err != nil {

			if err == io.EOF {
				return lines
			}

			check(err)
		}

		buffer.Write(part)
		if !prefix {
			lines = append(lines, buffer.String())
			buffer.Reset()
		}

	}
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
