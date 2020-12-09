package Challenges

import (
	"strconv"
)

/*
--- Day 8: Handheld Halting ---
Your flight to the major airline hub reaches cruising altitude without incident. While you consider checking the in-flight menu
for one of those drinks that come with a little umbrella, you are interrupted by the kid sitting next to you.

Their handheld game console won't turn on! They ask if you can take a look.

You narrow the problem down to a strange infinite loop in the boot code (your puzzle input) of the device. You should be able to
fix it, but first you need to be able to run the code in isolation.

The boot code is represented as a text file with one instruction per line of text. Each instruction consists of an operation
(acc, jmp, or nop) and an argument (a signed number like +4 or -20).

acc increases or decreases a single global value called the accumulator by the value given in the argument. For example, acc +7
would increase the accumulator by 7. The accumulator starts at 0. After an acc instruction, the instruction immediately below it
is executed next.
jmp jumps to a new instruction relative to itself. The next instruction to execute is found using the argument as an offset from
the jmp instruction; for example, jmp +2 would skip the next instruction, jmp +1 would continue to the instruction immediately below
it, and jmp -20 would cause the instruction 20 lines above to be executed next.
nop stands for No OPeration - it does nothing. The instruction immediately below it is executed next.
For example, consider the following program:

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
These instructions are visited in this order:

nop +0  | 1
acc +1  | 2, 8(!)
jmp +4  | 3
acc +3  | 6
jmp -3  | 7
acc -99 |
acc +1  | 4
jmp -4  | 5
acc +6  |
First, the nop +0 does nothing. Then, the accumulator is increased from 0 to 1 (acc +1) and jmp +4 sets the next instruction to the other
acc +1 near the bottom. After it increases the accumulator from 1 to 2, jmp -4 executes, setting the next instruction to the only acc +3.
It sets the accumulator to 5, and jmp -3 causes the program to continue back at the first acc +1.

This is an infinite loop: with this sequence of jumps, the program will run forever. The moment the program tries to run any instruction
a second time, you know it will never terminate.

Immediately before the program would run an instruction a second time, the value in the accumulator is 5.

Run your copy of the boot code. Immediately before any instruction is executed a second time, what value is in the accumulator?

--- Part Two ---
After some careful analysis, you believe that exactly one instruction is corrupted.

Somewhere in the program, either a jmp is supposed to be a nop, or a nop is supposed to be a jmp. (No acc instructions were harmed in the
	corruption of this boot code.)

The program is supposed to terminate by attempting to execute an instruction immediately after the last instruction in the file. By
changing exactly one jmp or nop, you can repair the boot code and make it terminate correctly.

For example, consider the same program from above:

nop +0
acc +1
jmp +4
acc +3
jmp -3
acc -99
acc +1
jmp -4
acc +6
If you change the first instruction from nop +0 to jmp +0, it would create a single-instruction infinite loop, never leaving that
instruction.
If you change almost any of the jmp instructions, the program will still eventually find another jmp instruction and loop forever.

However, if you change the second-to-last instruction (from jmp -4 to nop -4), the program terminates! The instructions are visited in
this order:

nop +0  | 1
acc +1  | 2
jmp +4  | 3
acc +3  |
jmp -3  |
acc -99 |
acc +1  | 4
nop -4  | 5
acc +6  | 6
After the last instruction (acc +6), the program terminates by attempting to run the instruction below the last instruction in the file.
With this change, after the program terminates, the accumulator contains the value 8 (acc +1, acc +1, acc +6).

Fix the program so that it terminates normally by changing exactly one jmp (to nop) or nop (to jmp). What is the value of the accumulator
after the program terminates?
*/

func Challenge8(useTestData bool) {
	var data []string
	if useTestData {
		data = challenge8TestData
	} else {
		data = challenge8InputData
	}

	println("Part 1")
	challenge8Part1(data)
	println("Part 2")
	challenge8Part2(data)
}

func challenge8Part1(data []string) {
	instructionUsed := make([]bool, len(data))
	count := 0
	for i := 0; !instructionUsed[i]; {
		inst := data[i][:3]
		instructionUsed[i] = true
		if inst == "nop" {
			i++
		} else if inst == "acc" {
			val, _ := strconv.Atoi(data[i][4:])
			count += val
			i++
		} else if inst == "jmp" {
			val, _ := strconv.Atoi(data[i][4:])
			i += val
		}
	}

	println(count)
}

func challenge8Part2(data []string) {
	instructionUsed := make([]bool, len(data))
	resVal, res := challenge8Part2Itterate(0, 0, data, instructionUsed, false)
	if !res {
		println("Error")
	}

	println(resVal)
}

func challenge8Part2Itterate(curr int, acc int, data []string, instructionUsed []bool, used bool) (int, bool) {
	if curr == len(data) {
		return acc, true
	} else if instructionUsed[curr] {
		return -1, false
	}

	inst := data[curr][:3]
	val, _ := strconv.Atoi(data[curr][4:])
	instructionUsed[curr] = true
	resVal, res := 0, false
	if inst == "acc" {
		resVal, res = challenge8Part2Itterate(curr+1, acc+val, data, instructionUsed, used)
	} else if inst == "nop" {
		resVal, res = challenge8Part2Itterate(curr+1, acc, data, instructionUsed, used)
		if !res && !used && val != 0 {
			resVal, res = challenge8Part2Itterate(curr+val, acc, data, instructionUsed, true)
		}
	} else {
		resVal, res = challenge8Part2Itterate(curr+val, acc, data, instructionUsed, used)
		if !res && !used {
			resVal, res = challenge8Part2Itterate(curr+1, acc, data, instructionUsed, true)
		}
	}

	instructionUsed[curr] = false
	return resVal, res
}

var challenge8InputData []string = []string{
	"acc +33",
	"acc -7",
	"acc +39",
	"jmp +214",
	"jmp +250",
	"jmp +51",
	"acc +29",
	"acc +6",
	"acc +20",
	"jmp +489",
	"nop +181",
	"acc +4",
	"jmp +187",
	"nop +454",
	"acc -10",
	"acc +44",
	"jmp +343",
	"acc +14",
	"acc +24",
	"acc +37",
	"acc -12",
	"jmp +596",
	"acc +21",
	"acc +39",
	"jmp +601",
	"acc -15",
	"jmp +304",
	"acc -7",
	"jmp +302",
	"acc +38",
	"jmp +148",
	"acc -6",
	"jmp +235",
	"acc +6",
	"nop +429",
	"acc +49",
	"acc +3",
	"jmp +255",
	"acc +2",
	"jmp +10",
	"acc +27",
	"acc +0",
	"acc -3",
	"acc +28",
	"jmp +565",
	"acc -16",
	"acc +39",
	"acc -5",
	"jmp +513",
	"acc +43",
	"acc +24",
	"jmp +26",
	"nop +19",
	"nop +71",
	"nop +182",
	"jmp +477",
	"acc +42",
	"jmp +535",
	"acc +38",
	"acc +29",
	"acc +1",
	"jmp +1",
	"jmp +72",
	"acc +25",
	"acc +43",
	"acc +6",
	"jmp +1",
	"jmp +111",
	"acc +43",
	"acc +13",
	"jmp +30",
	"acc +4",
	"acc +24",
	"acc +20",
	"acc -14",
	"jmp +161",
	"jmp +73",
	"nop +108",
	"jmp +547",
	"nop +273",
	"acc -8",
	"nop +358",
	"nop +284",
	"jmp +526",
	"acc +50",
	"jmp +274",
	"jmp +486",
	"nop +167",
	"acc -13",
	"jmp +11",
	"acc +10",
	"jmp +508",
	"acc -11",
	"acc +46",
	"acc +44",
	"jmp +335",
	"jmp +1",
	"acc -16",
	"acc +30",
	"jmp +289",
	"acc +15",
	"nop +265",
	"jmp +1",
	"nop +68",
	"jmp +107",
	"acc -15",
	"jmp -101",
	"acc +28",
	"acc -13",
	"jmp +17",
	"acc +21",
	"acc +46",
	"acc +19",
	"acc -8",
	"jmp +274",
	"nop +237",
	"jmp -111",
	"nop +419",
	"acc +28",
	"acc +26",
	"jmp +275",
	"acc -4",
	"jmp +483",
	"jmp +1",
	"jmp +201",
	"jmp +234",
	"acc +26",
	"acc +21",
	"acc +18",
	"jmp +149",
	"acc +0",
	"acc +29",
	"acc +11",
	"jmp -41",
	"nop +111",
	"nop +212",
	"jmp +172",
	"acc +31",
	"acc +17",
	"acc +6",
	"jmp -40",
	"acc +7",
	"acc +44",
	"acc +41",
	"acc +4",
	"jmp -74",
	"acc -16",
	"acc +37",
	"jmp +119",
	"acc -13",
	"acc +44",
	"acc +21",
	"acc +38",
	"jmp +92",
	"acc +30",
	"jmp +444",
	"jmp +35",
	"acc +3",
	"acc +11",
	"acc +31",
	"jmp -104",
	"acc -10",
	"acc +5",
	"acc +8",
	"acc +31",
	"jmp +127",
	"nop +168",
	"acc +16",
	"acc +6",
	"acc +0",
	"jmp +455",
	"acc +15",
	"acc +0",
	"acc +22",
	"acc -1",
	"jmp +191",
	"acc +16",
	"jmp +56",
	"acc -12",
	"acc +40",
	"nop -140",
	"acc +44",
	"jmp +138",
	"acc +44",
	"jmp +237",
	"acc +15",
	"acc +40",
	"jmp +360",
	"acc +14",
	"acc +14",
	"jmp +185",
	"nop +211",
	"acc +27",
	"acc -8",
	"acc +17",
	"jmp +247",
	"acc +50",
	"acc -2",
	"jmp -49",
	"acc +37",
	"jmp +330",
	"acc +14",
	"acc +44",
	"acc +15",
	"nop -43",
	"jmp +382",
	"jmp -45",
	"acc +46",
	"acc -11",
	"acc +47",
	"jmp +61",
	"nop +252",
	"acc +44",
	"acc -13",
	"jmp +292",
	"acc -6",
	"jmp +199",
	"acc +44",
	"acc +28",
	"acc +17",
	"acc +31",
	"jmp -158",
	"acc -8",
	"jmp +338",
	"acc +0",
	"acc -2",
	"nop +306",
	"jmp -78",
	"acc +11",
	"acc +33",
	"acc +40",
	"acc +33",
	"jmp -169",
	"jmp +273",
	"acc +8",
	"jmp -135",
	"acc +20",
	"acc -14",
	"acc -15",
	"nop +370",
	"jmp +20",
	"nop +51",
	"acc -4",
	"acc -10",
	"jmp -215",
	"acc +22",
	"acc +22",
	"jmp +209",
	"acc +40",
	"acc -18",
	"jmp -158",
	"jmp -130",
	"acc +13",
	"jmp -169",
	"nop +225",
	"acc +7",
	"jmp -23",
	"acc +21",
	"acc +0",
	"jmp +273",
	"jmp +293",
	"acc +39",
	"jmp -71",
	"acc +20",
	"jmp +49",
	"acc +6",
	"jmp -60",
	"acc +35",
	"jmp +84",
	"acc +14",
	"jmp +266",
	"acc +47",
	"jmp -247",
	"acc -3",
	"acc +47",
	"acc +23",
	"acc +30",
	"jmp +105",
	"acc +18",
	"jmp +109",
	"jmp -188",
	"nop -70",
	"acc -2",
	"acc +0",
	"jmp +195",
	"acc +15",
	"jmp +246",
	"acc +49",
	"acc +28",
	"jmp -18",
	"nop +120",
	"jmp +91",
	"acc -15",
	"acc +15",
	"acc +30",
	"jmp +39",
	"acc +46",
	"nop +250",
	"acc +49",
	"jmp -250",
	"acc -10",
	"acc +0",
	"acc +39",
	"jmp -254",
	"nop +55",
	"acc -4",
	"acc -3",
	"jmp +88",
	"jmp +35",
	"acc +47",
	"nop -154",
	"acc -16",
	"jmp +271",
	"nop +253",
	"jmp -199",
	"acc +5",
	"acc +35",
	"jmp +1",
	"acc +49",
	"jmp +234",
	"acc +27",
	"acc +33",
	"acc -3",
	"jmp -138",
	"jmp -107",
	"acc -11",
	"acc +47",
	"acc +14",
	"jmp -288",
	"jmp -205",
	"acc +0",
	"jmp +191",
	"acc -15",
	"jmp -116",
	"acc +35",
	"nop +121",
	"acc +2",
	"acc -14",
	"jmp +223",
	"acc +33",
	"acc -10",
	"acc +24",
	"jmp +73",
	"acc +39",
	"jmp +255",
	"acc +19",
	"jmp -16",
	"nop +1",
	"jmp -177",
	"nop +107",
	"nop -194",
	"jmp +260",
	"acc -16",
	"acc -12",
	"jmp -148",
	"acc +11",
	"acc +18",
	"acc +33",
	"jmp +84",
	"acc +27",
	"acc -13",
	"acc +36",
	"acc +26",
	"jmp +100",
	"nop -110",
	"jmp -98",
	"acc -2",
	"acc +29",
	"acc +25",
	"acc -8",
	"jmp +128",
	"acc +16",
	"acc +1",
	"acc +7",
	"jmp -290",
	"acc +18",
	"nop -235",
	"acc +0",
	"jmp -127",
	"acc -18",
	"acc +38",
	"jmp -297",
	"acc +19",
	"acc -8",
	"acc +20",
	"acc +3",
	"jmp -230",
	"jmp -67",
	"jmp +124",
	"acc -15",
	"acc +26",
	"acc -19",
	"jmp +120",
	"jmp +173",
	"jmp -338",
	"acc -15",
	"jmp -309",
	"acc +19",
	"acc +26",
	"acc +18",
	"acc +8",
	"jmp -6",
	"acc -7",
	"acc +10",
	"jmp -375",
	"acc +5",
	"acc -16",
	"acc +18",
	"acc +46",
	"jmp -309",
	"acc +48",
	"acc +40",
	"nop -227",
	"jmp -380",
	"jmp -290",
	"acc +46",
	"acc +5",
	"jmp -154",
	"acc -9",
	"acc +15",
	"jmp -187",
	"acc -10",
	"acc +0",
	"acc +28",
	"acc +30",
	"jmp -284",
	"acc +43",
	"acc +25",
	"acc +14",
	"jmp -205",
	"acc -13",
	"acc +1",
	"nop -340",
	"jmp -326",
	"jmp +1",
	"acc +9",
	"acc +17",
	"acc +1",
	"jmp -346",
	"jmp -158",
	"acc +23",
	"jmp -26",
	"nop -257",
	"jmp +140",
	"acc +11",
	"acc +10",
	"acc +29",
	"acc +48",
	"jmp +177",
	"acc +28",
	"acc -12",
	"acc -19",
	"acc +37",
	"jmp +79",
	"acc -14",
	"jmp -184",
	"nop +153",
	"jmp -170",
	"acc -17",
	"acc +10",
	"acc -6",
	"nop -174",
	"jmp -391",
	"jmp +148",
	"acc +50",
	"acc -8",
	"jmp -426",
	"jmp +1",
	"acc +16",
	"jmp +20",
	"jmp +1",
	"jmp -217",
	"nop +84",
	"jmp +71",
	"acc +16",
	"acc -7",
	"acc +23",
	"acc +24",
	"jmp -329",
	"acc +9",
	"acc -7",
	"acc -4",
	"nop +117",
	"jmp -16",
	"acc +30",
	"nop -222",
	"acc +32",
	"acc +9",
	"jmp -175",
	"acc +18",
	"acc +15",
	"acc +41",
	"jmp -192",
	"acc -3",
	"acc +8",
	"acc -13",
	"acc +24",
	"jmp -210",
	"acc +17",
	"acc -7",
	"acc -19",
	"jmp +76",
	"acc +26",
	"acc +2",
	"acc +4",
	"jmp +27",
	"jmp -104",
	"acc +38",
	"acc +46",
	"nop -67",
	"nop +37",
	"jmp -186",
	"jmp +5",
	"acc +37",
	"acc +8",
	"acc +30",
	"jmp -409",
	"acc +44",
	"acc +4",
	"jmp +109",
	"nop -8",
	"jmp -395",
	"acc +20",
	"acc +12",
	"acc +16",
	"acc +9",
	"jmp -87",
	"nop -406",
	"acc -8",
	"jmp -209",
	"jmp -137",
	"jmp -179",
	"acc +44",
	"jmp -399",
	"nop -141",
	"jmp +18",
	"jmp +1",
	"nop +55",
	"jmp +39",
	"acc +20",
	"acc +40",
	"acc +44",
	"acc +45",
	"jmp +74",
	"acc -16",
	"jmp -170",
	"jmp -48",
	"jmp -537",
	"acc -9",
	"acc +6",
	"nop -101",
	"acc +2",
	"jmp -418",
	"jmp -81",
	"jmp +1",
	"jmp -338",
	"nop +43",
	"acc +20",
	"jmp -109",
	"acc -1",
	"jmp -343",
	"acc +29",
	"acc +11",
	"nop -439",
	"jmp -310",
	"jmp -374",
	"acc +33",
	"nop +25",
	"acc -16",
	"nop -333",
	"jmp -14",
	"jmp -5",
	"jmp -162",
	"nop -432",
	"acc +16",
	"acc +17",
	"jmp -87",
	"acc -16",
	"nop -265",
	"acc +20",
	"jmp -356",
	"acc +0",
	"jmp +5",
	"acc +39",
	"acc -15",
	"jmp -325",
	"jmp -39",
	"nop -376",
	"nop -116",
	"acc +38",
	"jmp -175",
	"jmp -450",
	"jmp +1",
	"acc +19",
	"jmp -58",
	"nop -39",
	"acc +40",
	"acc +42",
	"jmp -232",
	"acc -14",
	"jmp -17",
	"acc +4",
	"acc -9",
	"acc +45",
	"jmp -229",
	"jmp -18",
	"acc +13",
	"acc +17",
	"jmp -591",
	"jmp -604",
	"jmp -356",
	"acc +1",
	"acc +18",
	"nop -52",
	"acc +39",
	"jmp -361",
	"jmp -303",
	"acc +8",
	"nop -477",
	"acc +3",
	"acc -8",
	"jmp -404",
	"acc +24",
	"acc +5",
	"jmp -88",
	"acc +27",
	"jmp -54",
	"jmp -18",
	"acc +31",
	"acc +40",
	"acc +18",
	"acc -16",
	"jmp +1",
}

var challenge8TestData []string = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}
