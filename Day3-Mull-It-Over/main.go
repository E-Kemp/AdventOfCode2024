package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Token int

const (
	MUL Token = iota
	DO
	DONT
)

type TokenStruct struct {
	token Token
	val1  int
	val2  int
}

func getData() []string {
	fileData, err := os.ReadFile("./sample.txt")
	if err == nil {
		return strings.Split(string(fileData), "\n")
	}
	panic(err)
}

func tokenise(input []string) []TokenStruct {
	var tokens = []TokenStruct{}
	rMul := regexp.MustCompile(`^mul\((\d+),(\d+)\)`)
	rDo := regexp.MustCompile(`^do\(\)`)
	rDont := regexp.MustCompile(`^don\'t\(\)`)

	for _, v := range input {
		for i := 0; i < len(v); i++ {
			if match := rMul.FindStringSubmatch(v[i:]); match != nil {
				val1, err := strconv.Atoi(match[1])
				val2, err := strconv.Atoi(match[2])

				if err != nil {
					val1 = -1
					val2 = -1
				}

				tokens = append(tokens, TokenStruct{
					token: MUL,
					val1:  val1,
					val2:  val2,
				})
				i += (len(match[0]) - 1)
			} else if match := rDo.FindStringSubmatch(v[i:]); match != nil {
				tokens = append(tokens, TokenStruct{
					token: DO,
					val1:  -1,
					val2:  -1,
				})
				i += 3
			} else if match := rDont.FindStringSubmatch(v[i:]); match != nil {
				tokens = append(tokens, TokenStruct{
					token: DONT,
					val1:  -1,
					val2:  -1,
				})
				i += 6
			}
		}
	}

	return tokens
}

func partOne(tokens []TokenStruct) int {
	sum := 0

	for _, token := range tokens {
		if token.token == MUL {
			sum += (token.val1 * token.val2)
		}
	}

	return sum
}

func partTwo(tokens []TokenStruct) int {
	sum := 0
	disabled := false

	for _, token := range tokens {
		switch token.token {
		case MUL:
			if !disabled {
				sum += (token.val1 * token.val2)
			}
		case DO:
			disabled = false
		case DONT:
			disabled = true
		}
	}

	return sum
}

func main() {
	data := getData()
	tokens := tokenise(data)
	ans1 := partOne(tokens)
	ans2 := partTwo(tokens)

	fmt.Printf("Total 1: %d\n", ans1)
	fmt.Printf("Total 2: %d\n", ans2)
}
