package main

import (
	"fmt"
	"time"
)

// getStartTime calculates the start time by going back N intervals of the specified aggregation in seconds,
// skipping weekends.
func getStartTime(startTime time.Time, aggregation, count int) (time.Time, int) {
	currentTime := startTime
	intervals := count
	i := 0

	for intervals > 0 {
		// Move back by one aggregation interval (in seconds)
		currentTime = currentTime.Add(-time.Duration(aggregation) * time.Second)

		// If it's a weekday, decrement the interval counter
		if currentTime.Weekday() >= time.Monday && currentTime.Weekday() <= time.Friday {
			i++
			intervals--
		} else {
			// Skip back to the previous Friday if it's a weekend
			for currentTime.Weekday() == time.Saturday || currentTime.Weekday() == time.Sunday {
				currentTime = currentTime.AddDate(0, 0, -1)
				currentTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 0, 0, currentTime.Location())
			}
		}
	}

	return currentTime, i
}

func main() {
	// Define the aggregation intervals in seconds
	aggregations := []int{
		60,
		120,
		180,
		300,
		600,
		900,
		1800,
		3600,
		7200,
		14400,
		28800,
		86400,
	}

	count := 1800
	endTime := time.Now().Add(-1 * time.Minute)

	// Generate and print the start time for each aggregation
	for _, agg := range aggregations {
		startTime, diff := getStartTime(endTime, agg, count)

		fmt.Printf("Aggregation %d, \tCounts:  %d, \tDiff: %d, \tStart time: %v,  \tEnd time: %v \n", agg, count, diff, startTime.Format("2006-01-02 15:04:05"), endTime.Format("2006-01-02 15:04:05"))
	}

	elapsed := time.Since(endTime).Seconds()
	fmt.Printf("Execution time: %v\n", elapsed)
}
