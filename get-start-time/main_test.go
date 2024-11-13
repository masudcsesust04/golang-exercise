package main

import (
	"testing"
	"time"
)

func TestGetStartTime(t *testing.T) {
	endTime := time.Now().Add(-1 * time.Minute)
	agg := 60
	count := 180
	_, diff := getStartTime(endTime, agg, count)
	if count != diff {
		t.Error("expected:", count, "got:", diff)
	}

}
