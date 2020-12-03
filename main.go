package main

import (
	"github.com/alistairfink/AdventOfCode2020/Challenges"
	"os"
	"path/filepath"
	"strconv"
)

var NAME string

func main() {
	args := os.Args
	NAME = formatName(args[0])
	if len(args) != 3 {
		printInvalid()
		return
	}

	challengeNumber := args[1]
	useTestData, err := strconv.ParseBool(args[2])
	if err != nil {
		panic(err)
	}

	if challengeNumber == "test" {
		Challenges.ChallengeTest()
	} else if challengeNumber == "1" {
		Challenges.Challenge1(useTestData)
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
	println("Usage:", NAME, "[Challenge Number]", "[Use Test Data (true/false)]")
}
