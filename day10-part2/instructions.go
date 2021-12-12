package main

import "strings"

type Instruction string

func (ins Instruction) CheckSyntax() int {
	var instructionHeap StringHeap

	for _, cmd := range strings.Split(string(ins), "") {
		switch cmd {
		case "(", "[", "{", "<":
			instructionHeap.Push(cmd)
		case ")":
			if instructionHeap.Pop() != "(" {
				return 3
			}
		case "]":
			if instructionHeap.Pop() != "[" {
				return 57
			}
		case "}":
			if instructionHeap.Pop() != "{" {
				return 1197
			}
		case ">":
			if instructionHeap.Pop() != "<" {
				return 25137
			}
		}
	}

	return 0
}

func (ins Instruction) Complete() int {
	var instructionHeap StringHeap

	for _, cmd := range strings.Split(string(ins), "") {
		switch cmd {
		case "(", "[", "{", "<":
			instructionHeap.Push(cmd)
		case ")":
			if instructionHeap.Pop() != "(" {
				return -1
			}
		case "]":
			if instructionHeap.Pop() != "[" {
				return -1
			}
		case "}":
			if instructionHeap.Pop() != "{" {
				return -1
			}
		case ">":
			if instructionHeap.Pop() != "<" {
				return -1
			}
		}
	}

	completionScore := 0
	for len(instructionHeap) != 0 {
		completionScore *= 5

		switch instructionHeap.Pop() {
		case "(":
			completionScore += 1
		case "[":
			completionScore += 2
		case "{":
			completionScore += 3
		case "<":
			completionScore += 4
		}
	}

	return completionScore
}
