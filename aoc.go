package aoc

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"strconv"
)

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
