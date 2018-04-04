package main

/*
NOTE: Have commented logs for this file,
if required please un-comment them
Reference -https://github.com/montanaflynn/stats
*/

import (
	"fmt"
	_ "log"
	"math"
	"sort"
)

// Sort a slice
func Sort(slice []float64) []float64 {
	new := make([]float64, len(slice))
	copy(new, slice)
	sort.Float64s(new)

	// log.Println("Copied new slice")
	return new
}

// Average for given slice
func Average(slice []float64) (float64, error) {

	if len(slice) == 0 {
		return math.NaN(), fmt.Errorf("Empty")
	}

	var sum float64
	for _, n := range slice {
		sum += n
	}

	avg := sum / float64(len(slice))
	// log.Println("Slice average", avg)
	return math.Round(avg*100) / 100, nil
}

// Median for given slice
func Median(values []float64) (median float64, err error) {

	l := len(values)
	if l == 0 {
		return math.NaN(), fmt.Errorf("Value can't be empty slice")
	} else if l%2 == 0 {
		median, _ = Average(values[l/2-1 : l/2+1])
	} else {
		median = float64(values[l/2])
	}

	// log.Println("Slice median", median)
	return median, nil
}

// Mode for given slice
func Mode(values []float64) (mode []float64, err error) {

	l := len(values)
	if l == 1 {
		return values, nil
	} else if l == 0 {
		return nil, fmt.Errorf("Empty")
	}

	mode = make([]float64, 5)
	cnt, maxCnt := 1, 1
	for i := 1; i < l; i++ {
		switch {
		case values[i] == values[i-1]:
			cnt++
		case cnt == maxCnt && maxCnt != 1:
			mode = append(mode, values[i-1])
			cnt = 1
		case cnt > maxCnt:
			mode = append(mode[:0], values[i-1])
			maxCnt, cnt = cnt, 1
		default:
			cnt = 1
		}
	}
	switch {
	case cnt == maxCnt:
		mode = append(mode, values[l-1])
	case cnt > maxCnt:
		mode = append(mode[:0], values[l-1])
		maxCnt = cnt
	}

	if maxCnt == 1 {
		return []float64{}, nil
	}

	// log.Println("Slice mode", mode)
	return mode, nil
}
