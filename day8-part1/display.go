package main

import (
	"errors"
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

func (d *Display) FindTrivialSignals() {
	for _, uniqueSignal := range d.UniqueSignals {
		switch len(uniqueSignal) {
		case 2: // when 2 wires are active -> display shows value: 1
			d.DecodedSignals[uniqueSignal] = 1
		case 4: // when 4 wires are active -> display shows value: 4
			d.DecodedSignals[uniqueSignal] = 4
		case 3: // when 3 wires are active -> display shows value: 7
			d.DecodedSignals[uniqueSignal] = 7
		case 7: // when 7 wires are active -> display shows value: 8
			d.DecodedSignals[uniqueSignal] = 8
		}
	}
}

func (d Display) DecodeOutput() ([]int, error) {
	var output []int
	var err error

	for _, outputSignal := range d.OutputSignals {
		if value, ok := d.DecodedSignals[outputSignal]; ok {
			output = append(output, value)
		} else {
			err = errors.New("failed to decode full output")
		}
	}

	return output, err
}
