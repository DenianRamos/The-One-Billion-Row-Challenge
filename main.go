package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Measurements struct {
	Min   float64
	Max   float64
	Sum   float64
	Count int
}

func main() {
	measurements, err := os.Open("measurements.txt")
	if err != nil {
		panic(err)
	}
	defer measurements.Close()

	data := make(map[string]Measurements)

	scanner := bufio.NewScanner(measurements)
	for scanner.Scan() {
		rawData := scanner.Text()
		semicolon := strings.Index(rawData, ";")
		location := rawData[:semicolon]
		rawTemp := rawData[semicolon+1:]

		temp, _ := strconv.ParseFloat(rawTemp, 64)

		measerument, ok := data[location]
		if !ok {
			measerument = Measurements{
				Min:   temp,
				Max:   temp,
				Sum:   temp,
				Count: 1,
			}
		} else {
			measerument.Min = min(measerument.Min, temp)
			measerument.Max = max(measerument.Max, temp)
			measerument.Sum += temp
			measerument.Count++

		}
		data[location] = measerument
	}
	locations := make([]string, 0, len(data))
	for name := range data {
		locations = append(locations, name)

		sort.Strings(locations)
		fmt.Println("Measurements:")
		for _, name := range locations {
			measurement := data[name]
			fmt.Printf("%s: Min: %.1f, Max: %.1f, Avg: %.1f\n", name, measurement.Min, measurement.Max, measurement.Sum/float64(measurement.Count))
		}
	}
	fmt.Println("\n")
}
