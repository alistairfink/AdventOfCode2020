package main

import (
	"github.com/alistairfink/AdventOfCode2020/Challenges"
	"os"
	"path/filepath"
)

var NAME string

func main() {
	args := os.Args
	NAME = formatName(args[0])
	if len(args) != 2 {
		printInvalid()
		return
	}

	challengeNumber := args[1]

	if challengeNumber == "test" {
		Challenges.ChallengeTest()
	} else {
		println("Invalid Challenge Number:", challengeNumber)
	}
}

func formatName(input string) string {
	input = filepath.ToSlash(input)
	nameI := len(input) - 1
	for ; nameI > 0 && input[nameI] != '/'; nameI-- {
	}

	if input[nameI] == '/' {
		nameI++
	}

	return input[nameI:]
}

func printInvalid() {
	println("Invalid Command.")
	println("Usage:", NAME, "[Challenge Number]")
}
