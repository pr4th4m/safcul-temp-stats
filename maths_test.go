package main

import (
	"reflect"
	"sort"
	"testing"
)

func TestMean(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{1, 2, 3, 4, 5}, 3.0},
		{[]float64{1, 2, 3, 4, 5, 6}, 3.5},
		{[]float64{1}, 1.0},
	} {
		got, _ := Average(c.in)
		if got != c.out {
			t.Errorf("Mean(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Average([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestMedian(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out float64
	}{
		{[]float64{5, 3, 4, 2, 1}, 4.0},
		{[]float64{6, 3, 2, 4, 5, 1}, 3.0},
		{[]float64{1}, 1.0},
	} {
		got, _ := Median(c.in)
		if got != c.out {
			t.Errorf("Median(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Median([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}

func TestMode(t *testing.T) {
	for _, c := range []struct {
		in  []float64
		out []float64
	}{
		{[]float64{5, 3, 4, 2, 1}, []float64{}},
		{[]float64{5, 5, 3, 4, 2, 1}, []float64{5}},
		{[]float64{5, 5, 3, 3, 4, 2, 1}, []float64{3, 5}},
		{[]float64{1}, []float64{1}},
		{[]float64{-50, -46.325, -46.325, -.87, 1, 2.1122, 3.20, 5, 15, 15, 15.0001}, []float64{-46.325, 15}},
		{[]float64{1, 2, 3, 4, 4, 4, 4, 4, 5, 3, 6, 7, 5, 0, 8, 8, 7, 6, 9, 9}, []float64{4}},
		{[]float64{76, 76, 110, 76, 76, 76, 76, 119, 76, 76, 76, 76, 31, 31, 31, 31, 83, 83, 83, 78, 78, 78, 78, 78, 78, 78, 78}, []float64{78}},
	} {
		got, err := Mode(c.in)
		if err != nil {
			t.Errorf("Returned an error")
		}
		sort.Float64s(got)
		if !reflect.DeepEqual(c.out, got) {
			t.Errorf("Mode(%.1f) => %.1f != %.1f", c.in, got, c.out)
		}
	}
	_, err := Mode([]float64{})
	if err == nil {
		t.Errorf("Empty slice should have returned an error")
	}
}
