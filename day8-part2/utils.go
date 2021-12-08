package main

import "strings"

func RemoveItemFromSlice(slice []string, removeItem string) []string {
	var newSlice []string

	for _, item := range slice {
		if item != removeItem {
			newSlice = append(newSlice, item)
		}
	}

	return newSlice
}

func StringSubstringMatches(s string, substrings []string) int {
	matches := 0
	for _, substring := range substrings {
		if strings.Contains(s, substring) {
			matches++
		}
	}
	return matches
}
