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
