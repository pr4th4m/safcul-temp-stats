package main

import (
	"testing"
)

func TestTemperature(t *testing.T) {

	expected := 4.08

	fridge := FridgeComponent{
		Source: "./data/fridge_temperature.json",
	}
	output, _ := fridge.Temperature()

	for _, out := range output {
		if out.ID == "b" {
			if out.Average != expected {
				t.Error("Excepted ", expected)
			}
		}
	}
}

func TestLoadSource(t *testing.T) {

	expected := 3.53

	fridges := []Fridge{}
	fridge := FridgeComponent{}
	fridge.loadSource("./data/fridge_temperature.json", &fridges)

	if fridges[0].Temperature != expected {
		t.Error("Excepted ", expected)
	}
}

func TestQuery(t *testing.T) {

	expected := 4.13

	fridges := []Fridge{}
	fridge := FridgeComponent{}
	fridge.loadSource("./data/fridge_temperature.json", &fridges)
	indexed := fridge.query(fridges)

	if indexed["b"][0] != expected {
		t.Error("Excepted ", expected)
	}
}
