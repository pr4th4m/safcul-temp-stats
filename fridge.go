package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// Fridge details
type Fridge struct {
	ID          string
	TimeStamp   int64
	Temperature float64
}

// Output to be displayed
type Output struct {
	ID      string    `json:"id"`
	Average float64   `json:"average"`
	Median  float64   `json:"median"`
	Mode    []float64 `json:"mode"`
}

// FridgeComponent - provides fridge stats
type FridgeComponent struct {
	Source string
}

// Temperature stats
func (frg *FridgeComponent) Temperature() (
	result []Output, err error) {

	// Load source file data into struct
	fridges := []Fridge{}
	if err := frg.loadSource(frg.Source, &fridges); err != nil {
		return nil, err
	}

	// Index data to get stats
	indexedFridges := frg.query(fridges)
	for key, slice := range indexedFridges {

		output := Output{}
		output.ID = key

		// NOTE: Slice is sorted to get
		// median and mode values
		sortedSlice := Sort(slice)

		// Get average
		avg, err := Average(sortedSlice)
		if err != nil {
			return nil, err
		}
		output.Average = avg

		// Get median
		med, err := Median(sortedSlice)
		if err != nil {
			return nil, err
		}
		output.Median = med

		// Get mode
		mod, err := Mode(sortedSlice)
		if err != nil {
			return nil, err
		}
		output.Mode = mod

		result = append(result, output)
	}

	return result, nil
}

// loadSource from file data
func (frg *FridgeComponent) loadSource(file string, fridges *[]Fridge) error {

	fridgeData, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	log.Println("Loaded source file ", file)
	return json.Unmarshal(fridgeData, fridges)
}

// query to get arranged data
func (frg *FridgeComponent) query(fridges []Fridge) map[string][]float64 {

	indexed := map[string][]float64{}
	for _, fridge := range fridges {

		if val, ok := indexed[fridge.ID]; ok {
			indexed[fridge.ID] = append(val, fridge.Temperature)
		} else {
			indexed[fridge.ID] = []float64{fridge.Temperature}
		}
	}

	log.Println("Indexed data for further quering")
	return indexed
}
