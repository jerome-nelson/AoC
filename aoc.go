package aoc

// Points for Chapter 1
// 1. Need a way to send input to script
// 2. Data is given as a HTML page - so need to scrape that
// 3. Need to look at the previous data when reviewing current data (Linked List type)
// 4. Need to keep a count of all increased times

func DepthTracker(listOfDepths []int) int {
	tracker := 0
	for order, currentDepth := range listOfDepths {
		var previousDepth int
		if order == 0 {
			continue
		}

		previousDepth = listOfDepths[order-1]
		if previousDepth < currentDepth {
			tracker += 1
		}

	}
	return tracker
}
