package main

import (
	"fmt"
	"sort"
	"strings"
)

type Display struct {
	UniqueSignals  []string
	OutputSignals  []string
	DecodedSignals map[string]int
}

func NewDisplay(uniqueSignals []string, outputSignals []string) Display {
	var sortedUniqueSignals []string
	var sortedOutputSignals []string

	for _, uniqueSignal := range uniqueSignals {
		s := strings.Split(uniqueSignal, "")
		sort.Strings(s)
		sortedUniqueSignals = append(sortedUniqueSignals, strings.Join(s, ""))
	}

	for _, outputSignal := range outputSignals {
		s := strings.Split(outputSignal, "")
		sort.Strings(s)
		sortedOutputSignals = append(sortedOutputSignals, strings.Join(s, ""))
	}

	return Display{
		UniqueSignals:  sortedUniqueSignals,
		OutputSignals:  sortedOutputSignals,
		DecodedSignals: make(map[string]int),
	}
}

func (d *Display) DecodeSignals() {
	valueToSignalMap := make(map[int]string)

	fmt.Printf("%#v\n", d.UniqueSignals)

	for _, uniqueSignal := range d.UniqueSignals {
		decoded := false

		switch len(uniqueSignal) {
		case 2: // when 2 wires are active -> display shows value: 1
			valueToSignalMap[1] = uniqueSignal
			fmt.Println("1", valueToSignalMap[1])
			decoded = true
		case 4: // when 4 wires are active -> display shows value: 4
			valueToSignalMap[4] = uniqueSignal
			fmt.Println("4", valueToSignalMap[4])
			decoded = true
		case 3: // when 3 wires are active -> display shows value: 7
			valueToSignalMap[7] = uniqueSignal
			fmt.Println("7", valueToSignalMap[7])
			decoded = true
		case 7: // when 7 wires are active -> display shows value: 8
			valueToSignalMap[8] = uniqueSignal
			fmt.Println("8", valueToSignalMap[8])
			decoded = true
		}

		if decoded {
			d.UniqueSignals = RemoveItemFromSlice(d.UniqueSignals, uniqueSignal)
		}
	}

	// 9 is the only value that uses 6 wires, including signals used by 4 and signals used by 7
	for _, uniqueSignal := range d.UniqueSignals {
		if len(uniqueSignal) == 6 && StringSubstringMatches(uniqueSignal, strings.Split(valueToSignalMap[4], "")) == len(valueToSignalMap[4]) &&
			StringSubstringMatches(uniqueSignal, strings.Split(valueToSignalMap[7], "")) == len(valueToSignalMap[7]) {

			valueToSignalMap[9] = uniqueSignal
			fmt.Println("9", valueToSignalMap[9])
			d.UniqueSignals = RemoveItemFromSlice(d.UniqueSignals, uniqueSignal)
			break
		}
	}

	// initially, there were three signals with a length of six, we found one of them (9) the last two might belong to either 0 and 6
	// 0 uses both wires of 1
	for _, uniqueSignal := range d.UniqueSignals {
		if len(uniqueSignal) == 6 && StringSubstringMatches(uniqueSignal, strings.Split(valueToSignalMap[1], "")) == len(valueToSignalMap[1]) {

			valueToSignalMap[0] = uniqueSignal
			fmt.Println("0", valueToSignalMap[0])
			d.UniqueSignals = RemoveItemFromSlice(d.UniqueSignals, uniqueSignal)
			break
		}
	}

	// initially, there were three signals with a length of six, we found two of them (0, 9) the last one belongs to 6
	for _, uniqueSignal := range d.UniqueSignals {
		if len(uniqueSignal) == 6 {

			valueToSignalMap[6] = uniqueSignal
			fmt.Println("6", valueToSignalMap[6])
			d.UniqueSignals = RemoveItemFromSlice(d.UniqueSignals, uniqueSignal)
			break
		}
	}

	// 5 uses similar wires like 9 with one difference while using one of the wires of 1
	for _, uniqueSignal := range d.UniqueSignals {
		if len(uniqueSignal) == 5 && StringSubstringMatches(uniqueSignal, strings.Split(valueToSignalMap[9], "")) == len(valueToSignalMap[9])-1 &&
			StringSubstringMatches(uniqueSignal, strings.Split(valueToSignalMap[1], "")) == 1 {

			valueToSignalMap[5] = uniqueSignal
			fmt.Println("5", valueToSignalMap[5])
			d.UniqueSignals = RemoveItemFromSlice(d.UniqueSignals, uniqueSignal)
			break
		}
	}

	// intially, there were three signals with a length of fice, we found one of them (5) the last two might belong to either 2 or 3
	// 3 uses both wires of 1
	for _, uniqueSignal := range d.UniqueSignals {
		if len(uniqueSignal) == 5 && StringSubstringMatches(uniqueSignal, strings.Split(valueToSignalMap[1], "")) == len(valueToSignalMap[1]) {

			valueToSignalMap[3] = uniqueSignal
			fmt.Println("3", valueToSignalMap[3])
			d.UniqueSignals = RemoveItemFromSlice(d.UniqueSignals, uniqueSignal)
			break
		}
	}

	// at this point there is only one signal left, which belongs to 2
	valueToSignalMap[2] = d.UniqueSignals[0]
	fmt.Println("2", valueToSignalMap[2])
	d.UniqueSignals = nil

	for value, signal := range valueToSignalMap {
		d.DecodedSignals[signal] = value
	}
}

func (d Display) DecodeOutput() ([]int, error) {
	var output []int
	var err error

	for _, outputSignal := range d.OutputSignals {
		if value, ok := d.DecodedSignals[outputSignal]; ok {
			output = append(output, value)
		} else {
			err = fmt.Errorf("failed to decode full output (did not find value for %s [debug: %#v])", outputSignal, d.DecodedSignals)
		}
	}

	return output, err
}
