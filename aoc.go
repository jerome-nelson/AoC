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

type DiveInstruction int64
type DiveCoords struct {
	x int
	y int
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

// Points for Chapter 2
// 1. Starts from Zero Integers
// 2. Essentially a graph with x, y co-ords
// 3. Forward (left) increments horizontal. Up/Down increments/decrements the vertical number
func CurrentDiveLocation(instruction []string) int {

	currentLocation := DiveCoords{
		x: 0,
		y: 0,
	}

	for _, order := range instruction {
		destructure := strings.Split(order, " ")
		direction, err := getDiveInstruction(destructure[0])
		check(err)
		number, err := strconv.Atoi(destructure[1])
		check(err)

		if direction == Right {
			currentLocation.x += number
		}

		if direction == Up {
			currentLocation.y -= number
		}

		if direction == Down {
			currentLocation.y += number
		}
	}

	return currentLocation.x * currentLocation.y
}

// Points for Chapter 1
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
