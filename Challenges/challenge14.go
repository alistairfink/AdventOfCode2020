package Challenges

import (
	"strconv"
)

/*
--- Day 14: Docking Data ---
As your ferry approaches the sea port, the captain asks for your help again. The computer system that runs this port isn't compatible
with the docking program on the ferry, so the docking parameters aren't being correctly initialized in the docking program's memory.

After a brief inspection, you discover that the sea port's computer system uses a strange bitmask system in its initialization program.
Although you don't have the correct decoder chip handy, you can emulate it in software!

The initialization program (your puzzle input) can either update the bitmask or write a value to memory. Values and memory addresses are
both 36-bit unsigned integers. For example, ignoring bitmasks for a moment, a line like mem[8] = 11 would write the value 11 to memory
address 8.

The bitmask is always given as a string of 36 bits, written with the most significant bit (representing 2^35) on the left and the least
significant bit (2^0, that is, the 1s bit) on the right. The current bitmask is applied to values immediately before they are written
to memory: a 0 or 1 overwrites the corresponding bit in the value, while an X leaves the bit in the value unchanged.

For example, consider the following program:

mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
mem[8] = 11
mem[7] = 101
mem[8] = 0
This program starts by specifying a bitmask (mask = ....). The mask it specifies will overwrite two bits in every written value: the 2s
bit is overwritten with 0, and the 64s bit is overwritten with 1.

The program then attempts to write the value 11 to memory address 8. By expanding everything out to individual bits, the mask is applied
as follows:

value:  000000000000000000000000000000001011  (decimal 11)
mask:   XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
result: 000000000000000000000000000001001001  (decimal 73)
So, because of the mask, the value 73 is written to memory address 8 instead. Then, the program tries to write 101 to address 7:

value:  000000000000000000000000000001100101  (decimal 101)
mask:   XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
result: 000000000000000000000000000001100101  (decimal 101)
This time, the mask has no effect, as the bits it overwrote were already the values the mask tried to set. Finally, the program tries to
write 0 to address 8:

value:  000000000000000000000000000000000000  (decimal 0)
mask:   XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X
result: 000000000000000000000000000001000000  (decimal 64)
64 is written to address 8 instead, overwriting the value that was there previously.

To initialize your ferry's docking program, you need the sum of all values left in memory after the initialization program completes.
(The entire 36-bit address space begins initialized to the value 0 at every address.) In the above example, only two values in memory
are not zero - 101 (at address 7) and 64 (at address 8) - producing a sum of 165.

Execute the initialization program. What is the sum of all values left in memory after it completes?

--- Part Two ---
For some reason, the sea port's computer system still can't communicate with your ferry's docking program. It must be using version 2
of the decoder chip!

A version 2 decoder chip doesn't modify the values being written at all. Instead, it acts as a memory address decoder. Immediately
before a value is written to memory, each bit in the bitmask modifies the corresponding bit of the destination memory address in the
following way:

If the bitmask bit is 0, the corresponding memory address bit is unchanged.
If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
If the bitmask bit is X, the corresponding memory address bit is floating.
A floating bit is not connected to anything and instead fluctuates unpredictably. In practice, this means the floating bits will take
on all possible values, potentially causing many memory addresses to be written all at once!

For example, consider the following program:

mask = 000000000000000000000000000000X1001X
mem[42] = 100
mask = 00000000000000000000000000000000X0XX
mem[26] = 1
When this program goes to write to memory address 42, it first applies the bitmask:

address: 000000000000000000000000000000101010  (decimal 42)
mask:    000000000000000000000000000000X1001X
result:  000000000000000000000000000000X1101X
After applying the mask, four bits are overwritten, three of which are different, and two of which are floating. Floating bits take on
every possible combination of values; with two floating bits, four actual memory addresses are written:

000000000000000000000000000000011010  (decimal 26)
000000000000000000000000000000011011  (decimal 27)
000000000000000000000000000000111010  (decimal 58)
000000000000000000000000000000111011  (decimal 59)
Next, the program is about to write to memory address 26 with a different bitmask:

address: 000000000000000000000000000000011010  (decimal 26)
mask:    00000000000000000000000000000000X0XX
result:  00000000000000000000000000000001X0XX
This results in an address with three floating bits, causing writes to eight memory addresses:

000000000000000000000000000000010000  (decimal 16)
000000000000000000000000000000010001  (decimal 17)
000000000000000000000000000000010010  (decimal 18)
000000000000000000000000000000010011  (decimal 19)
000000000000000000000000000000011000  (decimal 24)
000000000000000000000000000000011001  (decimal 25)
000000000000000000000000000000011010  (decimal 26)
000000000000000000000000000000011011  (decimal 27)
The entire 36-bit address space still begins initialized to the value 0 at every address, and you still need the sum of all values left
in memory at the end of the program. In this example, the sum is 208.

Execute the initialization program using an emulator for a version 2 decoder chip. What is the sum of all values left in memory after
it completes?
*/

func Challenge14(useTestData bool) {
	var data []string
	var data2 []string
	if useTestData {
		data = challenge14TestData
		data2 = challenge14TestDataPart2
	} else {
		data = challenge14InputData
		data2 = data
	}

	println("Part 1")
	challenge14Part1(data)
	println("Part 2")
	challenge14Part2(data2)
}

func challenge14Part1(data []string) {
	currMask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	regs := make(map[int]int)
	for _, row := range data {
		isMask, pos, bin, _, _ := challenge14ParseData(row)
		if isMask {
			currMask = bin
		} else {
			power := 1
			currSum := 0
			temp := ""
			for i := 0; i < 36; i++ {
				if currMask[35-i] == '1' ||
					(len(bin)-1-i >= 0 && currMask[35-i] == 'X' && bin[len(bin)-1-i] == '1') {
					currSum += power
					temp = "1" + temp
				} else {
					temp = "0" + temp
				}

				power *= 2
			}

			regs[pos] = currSum
		}
	}

	totalSum := 0
	for _, val := range regs {
		totalSum += val
	}

	println(totalSum)
}

func challenge14Part2(data []string) {
	currMask := "XXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXXX"
	regs := make(map[int]int)
	for _, row := range data {
		isMask, _, bin, posBin, numVal := challenge14ParseData(row)
		if isMask {
			currMask = bin
		} else {
			addressMask := make([]byte, 36)
			for i := 0; i < 36; i++ {
				if currMask[35-i] == '0' {
					if len(posBin)-1-i >= 0 {
						addressMask[35-i] = posBin[len(posBin)-1-i]
					} else {
						addressMask[35-i] = '0'
					}
				} else if currMask[35-i] == '1' {
					addressMask[35-i] = '1'
				} else {
					addressMask[35-i] = 'X'
				}
			}

			challenge14Part2ItterateRegs(35, 1, 0, numVal, regs, addressMask)
		}
	}

	totalSum := 0
	for _, val := range regs {
		totalSum += val
	}

	println(totalSum)
}

func challenge14Part2ItterateRegs(pos int, power int, curr int, replacementVal int, regs map[int]int, addressMask []byte) {
	if pos == -1 {
		regs[curr] = replacementVal
		return
	}

	if addressMask[pos] == 'X' || addressMask[pos] == '1' {
		challenge14Part2ItterateRegs(pos-1, power*2, curr+power, replacementVal, regs, addressMask)
	}

	if addressMask[pos] == 'X' || addressMask[pos] == '0' {
		challenge14Part2ItterateRegs(pos-1, power*2, curr, replacementVal, regs, addressMask)
	}
}

func challenge14ParseData(data string) (isMask bool, pos int, binary string, posBinary string, numVal int) {
	if data[:4] == "mask" {
		isMask = true
		binary = data[7:]
		pos = -1
	} else {
		isMask = false
		leftBracketPos, rightBracketPos, equalsPos := -1, -1, -1
		for i := 0; i < len(data); i++ {
			if data[i] == '[' {
				leftBracketPos = i
			} else if data[i] == ']' {
				rightBracketPos = i
			} else if data[i] == '=' {
				equalsPos = i
				break
			}
		}

		val, _ := strconv.Atoi(data[equalsPos+2:])
		binary = strconv.FormatInt(int64(val), 2)
		pos, _ = strconv.Atoi(data[leftBracketPos+1 : rightBracketPos])
		posBinary = strconv.FormatInt(int64(pos), 2)
		numVal = val
	}

	return isMask, pos, binary, posBinary, numVal
}

var challenge14InputData []string = []string{
	"mask = 11100010111110X010100001X00000011XXX",
	"mem[6540] = 1053547115",
	"mem[7184] = 10509",
	"mem[13014] = 7128",
	"mask = 110110X110100X10001X110110001000100X",
	"mem[23173] = 12045",
	"mem[414] = 4313357",
	"mem[18580] = 213079949",
	"mem[22303] = 1714",
	"mask = 1X01101X0110011XX1001X1XX001X00X1100",
	"mem[38327] = 97337978",
	"mem[24290] = 12593688",
	"mem[63524] = 3392",
	"mem[4214] = 243223",
	"mem[5420] = 935740301",
	"mem[18031] = 632714698",
	"mask = X10110101X1X011001100X11010010X00111",
	"mem[61349] = 1336335",
	"mem[50853] = 1269",
	"mem[58635] = 59668009",
	"mem[41045] = 18639",
	"mem[11918] = 57288",
	"mem[45645] = 14996953",
	"mem[26134] = 100867777",
	"mask = 11011010011XX110X10010X010011XX10000",
	"mem[10191] = 44785522",
	"mem[58901] = 12733",
	"mem[10035] = 11490",
	"mem[47614] = 20548428",
	"mem[31150] = 3738",
	"mem[4130] = 50709924",
	"mask = 1X0X101X01100100X101X11XX00X00X10000",
	"mem[49047] = 151086476",
	"mem[19156] = 271697",
	"mem[2453] = 498179",
	"mem[35232] = 401613",
	"mask = 110X0XX010XXX1001010011101011X00010X",
	"mem[46908] = 324",
	"mem[58098] = 7082",
	"mem[33784] = 87754344",
	"mem[55455] = 13375",
	"mem[7412] = 407",
	"mem[40954] = 287953968",
	"mem[8750] = 4872",
	"mask = X1011010X1X00X1001001110000110111010",
	"mem[16588] = 240485130",
	"mem[26942] = 4653",
	"mem[49588] = 2929",
	"mask = 010010X0X10X0X00X000X11010100011000X",
	"mem[20819] = 16307212",
	"mem[43449] = 110081329",
	"mem[53732] = 218919035",
	"mem[33252] = 1410886",
	"mem[9342] = 367201704",
	"mask = 1X01X010011XX11011000X10011XX1X0X100",
	"mem[31007] = 4767916",
	"mem[28355] = 6523",
	"mem[43831] = 522189",
	"mem[50230] = 25123775",
	"mem[9436] = 2916444",
	"mem[34444] = 29463",
	"mask = 010X101010XX011XX0110X001XX0X0001001",
	"mem[5352] = 766573",
	"mem[2337] = 1421026",
	"mem[16462] = 1440669",
	"mem[57643] = 243578224",
	"mem[37865] = 9557133",
	"mask = 11001010XXX0X11100100X011X0100010000",
	"mem[10674] = 92911",
	"mem[52874] = 314",
	"mem[4175] = 7882105",
	"mem[4647] = 437962",
	"mem[54577] = 6350",
	"mask = X1011110X1X1X1011000XXX1X1100X000001",
	"mem[8939] = 152882660",
	"mem[32991] = 1900820",
	"mem[26566] = 6277291",
	"mem[3398] = 338",
	"mem[198] = 1092255",
	"mem[42779] = 706535",
	"mask = 110110X100100110110X00XX000110001100",
	"mem[41624] = 107699",
	"mem[136] = 10907",
	"mem[59881] = 1639653",
	"mem[10152] = 13541233",
	"mem[29192] = 569789",
	"mask = 110011X10001XX0X00X11011111110111010",
	"mem[29497] = 61196354",
	"mem[33642] = 1690",
	"mem[3924] = 10240",
	"mem[30991] = 321393",
	"mem[6007] = 8252",
	"mem[10529] = 383134",
	"mem[22419] = 439560",
	"mask = 11011011001X00X00110X00111011X001011",
	"mem[9085] = 1844781",
	"mem[21531] = 14606246",
	"mem[6151] = 346328157",
	"mask = 1XX110X00011111011X100X000X00011X111",
	"mem[20254] = 4775875",
	"mem[18751] = 34439",
	"mem[51931] = 10866",
	"mem[13884] = 1396272",
	"mem[38876] = 110002",
	"mem[10554] = 3811164",
	"mask = 110X10101X1101110010X01X1001X000000X",
	"mem[24882] = 5606803",
	"mem[60456] = 235252",
	"mem[29117] = 1172",
	"mask = 11111010X10X011010100001X10010XXXX10",
	"mem[22085] = 3767",
	"mem[35246] = 637281",
	"mem[57643] = 1015666",
	"mask = 010X1010X00X011110X1X1X100000110X0X1",
	"mem[2766] = 1017781",
	"mem[18472] = 65544229",
	"mem[60177] = 44130",
	"mem[49328] = 252860",
	"mem[63588] = 2718",
	"mem[25385] = 1040",
	"mask = 1X1101X01X10XX0010101X101101X00X1101",
	"mem[1085] = 1816",
	"mem[19718] = 2409",
	"mem[61449] = 466156",
	"mask = 11011000100X01X1000000X0010X01001001",
	"mem[27929] = 169017783",
	"mem[3960] = 138790",
	"mem[49378] = 3765",
	"mem[56654] = 395338",
	"mem[64888] = 11900103",
	"mask = 11X01010XXX001X0101010X0010110101X11",
	"mem[44604] = 28291214",
	"mem[22847] = 316570",
	"mask = X111X11X11X111001X1011X101000X100000",
	"mem[45942] = 2947119",
	"mem[14974] = 2025710",
	"mem[24828] = 2567",
	"mem[10597] = 765995",
	"mem[45790] = 461",
	"mem[4860] = 621896",
	"mask = X10X1X110XX1000X001X11100111001100X1",
	"mem[34726] = 1325661",
	"mem[15011] = 13271770",
	"mem[49977] = 455",
	"mem[13385] = 359763",
	"mem[64427] = 442514062",
	"mem[37436] = 240727423",
	"mask = X10X10X00X1001000X0X10X000011010X0X1",
	"mem[5887] = 567410",
	"mem[15160] = 51659",
	"mask = 110X1X10100101110000X10X110X0X011XX1",
	"mem[48320] = 1175498",
	"mem[8272] = 829060",
	"mem[62848] = 373913",
	"mem[51709] = 1454",
	"mem[7351] = 28999323",
	"mask = 110X1010X11X011XX010X101X0001001XX10",
	"mem[5420] = 194080",
	"mem[64003] = 114629",
	"mem[47520] = 30982831",
	"mem[9857] = 9284958",
	"mem[54577] = 1423",
	"mem[48414] = 5845567",
	"mem[40311] = 19305",
	"mask = 11X110X10011000100X01110XX1100X11011",
	"mem[6785] = 55956876",
	"mem[27761] = 2895313",
	"mem[48178] = 815",
	"mem[42069] = 264595943",
	"mem[18031] = 42",
	"mem[40874] = 974080914",
	"mask = 0X0110100X10010X0X01100X0111XX1XX010",
	"mem[27948] = 238815667",
	"mem[4959] = 142720651",
	"mem[62738] = 28494",
	"mask = 1101X01001100X10X0X01X11001000XX00XX",
	"mem[54937] = 1139391",
	"mem[25222] = 2512343",
	"mem[9470] = 120874",
	"mem[26552] = 3096",
	"mem[24626] = 51146359",
	"mem[60877] = 1405",
	"mask = 1X00101001X1011110XXX1010101101X0X10",
	"mem[29192] = 1147",
	"mem[5280] = 64971",
	"mem[59069] = 18787283",
	"mem[8736] = 252",
	"mem[27948] = 7736001",
	"mask = 110110X00110X1X00101XX10010XX0111001",
	"mem[39600] = 150148438",
	"mem[19510] = 2028",
	"mem[10561] = 760736",
	"mem[33637] = 3786",
	"mem[27840] = 19776",
	"mem[43173] = 113509213",
	"mask = 0101X010100X0111001100001X00010X1101",
	"mem[19019] = 109550",
	"mem[61970] = 9527643",
	"mem[54718] = 146897567",
	"mem[34440] = 23504976",
	"mask = X101101010X10111001X0XX01X000XXX1001",
	"mem[27559] = 32347",
	"mem[31834] = 5553",
	"mem[25888] = 816",
	"mem[50958] = 143680145",
	"mem[40800] = 1579721",
	"mem[3365] = 7205751",
	"mask = 11011X101110X1111X1011XX101X00000000",
	"mem[39744] = 1628496",
	"mem[26998] = 174246124",
	"mem[55100] = 78394066",
	"mem[41107] = 515946",
	"mem[11859] = 6199516",
	"mem[6466] = 49035",
	"mem[9382] = 132658816",
	"mask = 1001X010011001X1010011X01X0X0XX11110",
	"mem[43573] = 261285686",
	"mem[62900] = 80070",
	"mem[39087] = 711323353",
	"mem[3697] = 2909066",
	"mem[30760] = 42870790",
	"mem[19499] = 380182",
	"mask = 1101101XXX100110X100X011X00000X01000",
	"mem[48681] = 42494",
	"mem[58173] = 6444223",
	"mem[42666] = 167047779",
	"mem[55462] = 4061450",
	"mem[4524] = 2982705",
	"mem[25898] = 122811236",
	"mask = X10010X00X00011X10111111100001001100",
	"mem[55341] = 907685",
	"mem[25896] = 5234228",
	"mem[46456] = 205011032",
	"mask = 110X10X001X1111011000000011X0101110X",
	"mem[14875] = 1333",
	"mem[56659] = 15180",
	"mem[45783] = 104761",
	"mem[4214] = 26033",
	"mask = 0X001X0001010100X000010010X100101000",
	"mem[2356] = 182",
	"mem[12067] = 524650703",
	"mem[14111] = 5231",
	"mem[46413] = 60005",
	"mem[14331] = 221755997",
	"mem[10416] = 589456",
	"mask = 1101101010010X11001110101X110X01X0X0",
	"mem[3517] = 125300",
	"mem[48255] = 4322",
	"mem[25758] = 724",
	"mem[3620] = 780",
	"mem[35335] = 3264540",
	"mask = 011X0X10111X1100101X0010X101XXX1X110",
	"mem[50922] = 1340",
	"mem[49956] = 7971",
	"mem[16765] = 117074",
	"mem[37341] = 101880",
	"mem[6382] = 270794",
	"mem[62129] = 3524",
	"mask = X1110X10111X110010101X111100X0XX01X1",
	"mem[29143] = 51870362",
	"mem[44539] = 5471456",
	"mem[17812] = 555343768",
	"mem[25896] = 76993",
	"mem[61860] = 5583277",
	"mem[33249] = 1649382",
	"mask = 10X110X00X111XX01X110X001001X0X00101",
	"mem[60418] = 40765",
	"mem[9791] = 276",
	"mem[20819] = 1939",
	"mem[16099] = 50776",
	"mem[44908] = 13705675",
	"mem[3842] = 12939",
	"mask = 1100101X1X000110101000011101X0XX1001",
	"mem[64795] = 7473",
	"mem[20801] = 10933303",
	"mem[45749] = 15523455",
	"mem[4130] = 14179780",
	"mem[24197] = 22821",
	"mem[5171] = 1847",
	"mem[56522] = 80",
	"mask = X0X1X011011001001101011X1XX00111X000",
	"mem[23931] = 391795",
	"mem[62129] = 46071660",
	"mem[10529] = 340467299",
	"mem[22948] = 590",
	"mem[18580] = 986058",
	"mem[23943] = 6228242",
	"mask = 111X10101X1101X01010X1010X0100000101",
	"mem[22045] = 1236872",
	"mem[43412] = 11981",
	"mem[51116] = 228048",
	"mem[4214] = 34658",
	"mem[14961] = 28301063",
	"mem[19718] = 125855455",
	"mask = 110110110XX101X0011001X10100001110X1",
	"mem[29192] = 375215772",
	"mem[43812] = 106881120",
	"mem[22916] = 11512817",
	"mem[35299] = 161518",
	"mem[39569] = 390",
	"mask = 1X0110100110X100X101011X1101X0010X11",
	"mem[5014] = 15731815",
	"mem[42765] = 90079",
	"mem[44656] = 360472556",
	"mem[4214] = 2023626",
	"mem[5777] = 15044384",
	"mem[35463] = 1668",
	"mask = 1X10X010X01X011010X0X00X00X100101011",
	"mem[29192] = 975691",
	"mem[3256] = 478673",
	"mem[3960] = 4314",
	"mem[56647] = 827",
	"mem[4162] = 228123",
	"mask = X101001001100110XX101010011010001X00",
	"mem[3719] = 312891161",
	"mem[44656] = 19029928",
	"mem[45713] = 23393",
	"mem[45594] = 1879",
	"mem[20653] = 392",
	"mem[55312] = 115889500",
	"mask = 01X01010XX0X0100000X0100X010XX0110X1",
	"mem[1085] = 466780",
	"mem[27568] = 13308",
	"mem[6034] = 362313775",
	"mask = 10011X1001X0X100010110100001100X0101",
	"mem[49413] = 422627",
	"mem[29467] = 25028765",
	"mem[11378] = 11086",
	"mask = 0X0110X0100101110011X1X0X10X01100010",
	"mem[62515] = 10720832",
	"mem[55462] = 1012292",
	"mem[21513] = 762797",
	"mem[2053] = 21634153",
	"mem[11742] = 168511869",
	"mem[38283] = 90760561",
	"mem[36076] = 23295980",
	"mask = 11110010110111X010X0X00111111010X011",
	"mem[1828] = 3667",
	"mem[11358] = 522202",
	"mem[41729] = 28977762",
	"mem[37384] = 25384",
	"mem[33843] = 450",
	"mem[3188] = 71694",
	"mem[43265] = 107204037",
	"mask = 1XX110100X111110X10XX000X00000X1111X",
	"mem[51454] = 15544",
	"mem[30926] = 3747088",
	"mem[57622] = 3670672",
	"mem[37268] = 1355",
	"mem[31994] = 492591",
	"mask = 1X01X0100X1X0X10X000011110X11100X1X0",
	"mem[3719] = 889",
	"mem[60942] = 4205",
	"mem[5426] = 1425",
	"mem[39400] = 1357",
	"mem[48231] = 110916",
	"mask = X1XX0X10111011X0101X01X00X01000001X0",
	"mem[25222] = 565",
	"mem[25657] = 4325",
	"mem[6319] = 358",
	"mask = 1X1X101010000110X0X01111110X0010001X",
	"mem[10685] = 354498",
	"mem[40671] = 1078",
	"mem[20831] = 694693",
	"mem[56347] = 2893",
	"mem[41145] = 38570",
	"mem[877] = 24807",
	"mask = 100110X001111X10X1111010X101010010X1",
	"mem[59369] = 1536028",
	"mem[12650] = 31302991",
	"mem[62737] = 1063",
	"mem[44908] = 1121896",
	"mem[50981] = 1793",
	"mem[41549] = 467",
	"mask = 1101X01X1110011000X001X1000011X11000",
	"mem[7888] = 23809",
	"mem[51709] = 94998",
	"mem[57643] = 5739",
	"mem[4860] = 58006462",
	"mem[9555] = 172891",
	"mem[44267] = 17211850",
	"mem[40597] = 28094",
	"mask = 10011011X11001X0X100010X000110001100",
	"mem[61860] = 4412269",
	"mem[7888] = 113388380",
	"mem[57911] = 100615718",
	"mask = 1101101X0X110XX00X100101X10110X01X10",
	"mem[58901] = 337510",
	"mem[20507] = 14011947",
	"mem[1297] = 7190797",
	"mask = 1100101X0X100100010X10XX110X101XX011",
	"mem[58901] = 338413",
	"mem[29192] = 18405358",
	"mem[14356] = 1740",
	"mask = X10X1010X1100111011X101000X101110100",
	"mem[63588] = 62",
	"mem[37958] = 10875",
	"mem[57170] = 11532",
	"mask = 11001011011X0100110101X00000100100XX",
	"mem[23685] = 4073432",
	"mem[52476] = 48853",
	"mem[8464] = 1362",
	"mask = 110X10X011X111X0101010X1000110100101",
	"mem[58112] = 202002",
	"mem[31524] = 348910462",
	"mem[7463] = 379275",
	"mem[46783] = 516",
	"mask = 1XXX00XX1XX11X001010000X010000000100",
	"mem[59845] = 1705",
	"mem[29610] = 908",
	"mem[38876] = 7419",
	"mem[50665] = 27526",
	"mem[55828] = 239766",
	"mem[7198] = 2332510",
	"mask = 1101101X111001100X10X101010X1XX10010",
	"mem[58901] = 4157116",
	"mem[30155] = 2181",
	"mem[27948] = 167505",
	"mem[5615] = 6191",
	"mem[10685] = 6718001",
	"mem[10907] = 8060",
	"mask = 11001010X111X11X10100000X10011010111",
	"mem[7185] = 199563702",
	"mem[21941] = 381",
	"mem[58901] = 5536",
	"mem[57911] = 16158",
	"mem[61179] = 43997",
	"mask = 110110X0011X01XX01X10001101100110011",
	"mem[37515] = 733",
	"mem[21609] = 4275",
	"mem[60877] = 24488",
	"mask = 11X1101110100110XX10000X0X100111X100",
	"mem[62234] = 5991",
	"mem[18480] = 60587",
	"mask = X0X11X1000100110000010010011111001X0",
	"mem[54105] = 6396",
	"mem[26916] = 65750915",
	"mask = 1001101001X0011100000X111111000011XX",
	"mem[10709] = 1636726",
	"mem[64283] = 8450605",
	"mem[36506] = 2982695",
	"mem[16588] = 84253",
	"mem[20114] = 2702",
	"mask = 010X1X1001XXX100X00001000X1000X10000",
	"mem[41718] = 849",
	"mem[55828] = 198795786",
	"mem[44985] = 238093315",
	"mask = 1X0110101110X111011001011X0X0100X100",
	"mem[61680] = 1953",
	"mem[1986] = 28915429",
	"mem[50342] = 453382485",
	"mem[30946] = 901529185",
	"mem[61748] = 349278",
	"mem[48297] = 78511621",
	"mask = 1101111101100111110000001XXX10101X00",
	"mem[34726] = 424077",
	"mem[662] = 6076351",
	"mem[36217] = 6631187",
	"mem[40090] = 55913370",
	"mask = 1X01101001X00X110000100010X000100000",
	"mem[48146] = 7118",
	"mem[47902] = 793",
	"mem[42824] = 612043",
	"mem[55984] = 49827115",
	"mem[2356] = 2180174",
	"mem[5352] = 156",
	"mask = X001101101X00X10110X1X1X11X100101100",
	"mem[7853] = 16332563",
	"mem[14356] = 792",
	"mem[17700] = 9478",
	"mask = 010110100010010000XX111X1X1110X00000",
	"mem[47094] = 295",
	"mem[43272] = 321261",
	"mem[48398] = 791",
	"mask = 10011010011001000X011X1011011000X0X1",
	"mem[59153] = 359805",
	"mem[59453] = 2389171",
	"mem[34721] = 1050",
	"mem[25615] = 205647",
	"mask = 11011010011001XX0XXX111X10010011XXX0",
	"mem[49588] = 24258309",
	"mem[9857] = 125489601",
	"mem[26998] = 57828611",
	"mem[30155] = 1226221",
	"mask = X10X10X001100100000X0001X010X01100X0",
	"mem[61349] = 7180",
	"mem[21835] = 22741",
	"mem[40954] = 191961",
	"mem[57643] = 482698",
	"mem[41411] = 17811",
	"mem[55462] = 21319776",
	"mem[47515] = 2615187",
	"mask = X1111010XX1101XX10101X001X0X00000100",
	"mem[6151] = 1463126",
	"mem[52644] = 7783778",
	"mask = 11X110111X100110XX1011010XX00X01111X",
	"mem[30760] = 64",
	"mem[53149] = 225556464",
	"mem[28495] = 124093",
	"mem[14116] = 18796146",
	"mem[16522] = 1185502",
	"mask = 01X0XX1X1110110X1011X011X10110000100",
	"mem[2511] = 1415",
	"mem[6645] = 21276",
	"mem[15382] = 14111452",
	"mask = X1XX1010011101110X1010X11X10X0000X11",
	"mem[52197] = 1486281",
	"mem[64632] = 12145559",
	"mem[61183] = 74388",
	"mem[51657] = 6708652",
	"mem[57489] = 62339532",
	"mask = 110010100XX0X1X0010X000010001000XXX1",
	"mem[57761] = 2008492",
	"mem[27948] = 247057",
	"mem[33843] = 92966018",
	"mem[44458] = 278",
	"mem[52801] = 437",
	"mask = 11011X1101100X1X11X01001XXXX00011100",
	"mem[24290] = 1456055",
	"mem[23630] = 20246",
	"mem[33978] = 352201",
	"mem[12584] = 1718047",
	"mask = 1X0110X0100X011100XX0110100001X010X1",
	"mem[2513] = 4373",
	"mem[16175] = 26467036",
	"mem[19792] = 19895",
	"mem[40664] = 28643",
	"mem[57643] = 264",
	"mem[45471] = 6119525",
	"mem[35127] = 5644595",
	"mask = 1X01X010011001X101000101000000111X10",
	"mem[52859] = 1753381",
	"mem[42765] = 62301101",
	"mem[30314] = 13618",
	"mem[19897] = 2528961",
	"mem[44693] = 14301",
	"mem[51386] = 3407",
	"mem[29467] = 244447",
	"mask = 1X01101001X0011X0100X000010100101001",
	"mem[43014] = 7202",
	"mem[55776] = 255530898",
	"mem[58195] = 10425",
	"mem[43482] = 57072363",
	"mem[17185] = 237800057",
	"mask = 1XX11XXX1101110X10X000011X100000X111",
	"mem[7170] = 16343221",
	"mem[7498] = 52010",
	"mem[14356] = 11739",
	"mask = 1X0X0000X111X1X010X0011XX1010100X000",
	"mem[64110] = 318",
	"mem[36357] = 10409392",
	"mem[27061] = 17517",
	"mem[42437] = 15677",
	"mem[49032] = 9620714",
	"mask = X10110101011011100X01X01101101000X01",
	"mem[55980] = 134072224",
	"mem[4807] = 6332",
	"mem[23989] = 1457360",
	"mem[16588] = 1148127",
	"mem[65214] = 634126",
	"mem[56601] = 198043",
	"mem[1770] = 860",
	"mask = 1X01101001X00111010010010X10000X1101",
	"mem[10101] = 495754815",
	"mem[48259] = 1648",
	"mem[45883] = 84939765",
	"mem[10110] = 10261",
	"mask = 1X0110100110011100X01XX010X10XX1X000",
	"mem[34342] = 4318",
	"mem[49853] = 19927",
	"mem[49142] = 105722641",
	"mem[2453] = 110009314",
	"mem[26988] = 55762",
	"mem[26552] = 10874",
	"mask = 1001X0XX1101100X101XX00111000000X100",
	"mem[29409] = 6718",
	"mem[42621] = 26872185",
	"mem[2074] = 509",
	"mem[30155] = 161238",
	"mem[11121] = 115558625",
	"mem[18317] = 62070",
	"mask = 11X1X01011X1X1X01010000110X0000XX11X",
	"mem[54718] = 3211036",
	"mem[34679] = 107269681",
	"mem[37648] = 32371325",
	"mem[34726] = 1541",
	"mem[63265] = 1069841040",
	"mem[7185] = 1112080",
}

var challenge14TestData []string = []string{
	"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
	"mem[8] = 11",
	"mem[7] = 101",
	"mem[8] = 0",
}

var challenge14TestDataPart2 []string = []string{
	"mask = 000000000000000000000000000000X1001X",
	"mem[42] = 100",
	"mask = 00000000000000000000000000000000X0XX",
	"mem[26] = 1",
}
