package banners

import (
	"fmt"
	"os"
	"strings"
)

var encodeDictionary = map[string]int{
	" ":  1,
	"!":  10,
	"\"": 19,
	"#":  28,
	"$":  37,
	"%":  46,
	"&":  55,
	"'":  64,
	"(":  73,
	")":  82,
	"*":  91,
	"+":  100,
	",":  109,
	"-":  118,
	".":  127,
	"/":  136,
	"0":  145,
	"1":  154,
	"2":  163,
	"3":  172,
	"4":  181,
	"5":  190,
	"6":  199,
	"7":  208,
	"8":  217,
	"9":  226,
	":":  235,
	";":  244,
	"<":  253,
	"=":  262,
	">":  271,
	"?":  280,
	"@":  289,
	"A":  298,
	"B":  307,
	"C":  316,
	"D":  325,
	"E":  334,
	"F":  343,
	"G":  352,
	"H":  361,
	"I":  370,
	"J":  379,
	"K":  388,
	"L":  397,
	"M":  406,
	"N":  415,
	"O":  424,
	"P":  433,
	"Q":  442,
	"R":  451,
	"S":  460,
	"T":  469,
	"U":  478,
	"V":  487,
	"W":  496,
	"X":  505,
	"Y":  514,
	"Z":  523,
	"[":  532,
	"\\": 541,
	"]":  550,
	"^":  559,
	"_":  568,
	"`":  577,
	"a":  586,
	"b":  595,
	"c":  604,
	"d":  613,
	"e":  622,
	"f":  631,
	"g":  640,
	"h":  649,
	"i":  658,
	"j":  667,
	"k":  676,
	"l":  685,
	"m":  694,
	"n":  703,
	"o":  712,
	"p":  721,
	"q":  730,
	"r":  739,
	"s":  748,
	"t":  757,
	"u":  766,
	"v":  775,
	"w":  784,
	"x":  793,
	"y":  802,
	"z":  811,
	"{":  820,
	"|":  829,
	"}":  838,
	"~":  847,
}


// returns encoded text as an string array
func EncodeText(input string, trigger string) []string {
	fileStream, err := os.ReadFile("banners/" + trigger + ".txt")
	if err != nil {
		fmt.Println(err)
	}
	standardLetters := strings.Split(string(fileStream), "\n")

	outputRow1 := ""
	outputRow2 := ""
	outputRow3 := ""
	outputRow4 := ""
	outputRow5 := ""
	outputRow6 := ""
	outputRow7 := ""
	outputRow8 := ""

	for i := 0; i < len(input); i++ {
		outputRow1 += standardLetters[encodeDictionary[string(input[i])]+0]
		outputRow2 += standardLetters[encodeDictionary[string(input[i])]+1]
		outputRow3 += standardLetters[encodeDictionary[string(input[i])]+2]
		outputRow4 += standardLetters[encodeDictionary[string(input[i])]+3]
		outputRow5 += standardLetters[encodeDictionary[string(input[i])]+4]
		outputRow6 += standardLetters[encodeDictionary[string(input[i])]+5]
		outputRow7 += standardLetters[encodeDictionary[string(input[i])]+6]
		outputRow8 += standardLetters[encodeDictionary[string(input[i])]+7]
	}
	output := []string{outputRow1, outputRow2, outputRow3, outputRow4, outputRow5, outputRow6, outputRow7, outputRow8}
	return output
}
