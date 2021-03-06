package gostats

import (
  "time"
)

type DurationStats struct {
	Avg time.Duration
	Max time.Duration
	Min time.Duration
}

// GetDurationStatsFunc returns the DurationStats for the previous setup
func GetDurationStatsFunc(binsize int) func(time.Duration) DurationStats {
	bins := make([]time.Duration, binsize)
	var sum time.Duration
	var count int
	pos := 0
	return func(t time.Duration) DurationStats {
		if t != 0 {
			sum -= bins[pos]
			sum += t
			bins[pos] = t
			pos++
			if count < binsize {
				count++
			}
			if pos >= binsize {
				pos = 0
			}
		}
		if count == 0 {
			return DurationStats{}
		}

		min := time.Duration(1<<63 - 1)
		var max time.Duration
		for _, a := range bins {
			if a > max {
				max = a
			}
			if a < min {
				min = a
			}
		}
		return DurationStats{Avg: time.Duration(sum / time.Duration(count)), Max: max, Min: min}
	}
}
