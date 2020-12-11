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
	} else if challengeNumber == "2" {
		Challenges.Challenge2(useTestData)
	} else if challengeNumber == "3" {
		Challenges.Challenge3(useTestData)
	} else if challengeNumber == "4" {
		Challenges.Challenge4(useTestData)
	} else if challengeNumber == "5" {
		Challenges.Challenge5(useTestData)
	} else if challengeNumber == "6" {
		Challenges.Challenge6(useTestData)
	} else if challengeNumber == "7" {
		Challenges.Challenge7(useTestData)
	} else if challengeNumber == "8" {
		Challenges.Challenge8(useTestData)
	} else if challengeNumber == "9" {
		Challenges.Challenge9(useTestData)
	} else if challengeNumber == "10" {
		Challenges.Challenge10(useTestData)
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
