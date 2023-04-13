package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)
var (
	g, p, b int
	B       int
	A       int
	S       int
)

func getGValueAndPValue() {
	input := extractInput()

	gValueIndex := strings.Index(input, "g is ")
	pValueIndex := strings.Index(input, "p is ")
	separatorValueIndex := strings.Index(input, " and ")
	if gValueIndex != -1 && pValueIndex != -1 && separatorValueIndex != -1 {
		g, _ = strconv.Atoi(input[gValueIndex+5 : separatorValueIndex])
		p, _ = strconv.Atoi(input[pValueIndex+5:])
	}
}

func extractInput() string {
	var input string
	input, _ = reader.ReadString('\n')
	input = strings.TrimSpace(input)
	return input
}

func getAValue() {
	input := extractInput()
	AIndex := strings.Index(input, "A is ")
	if AIndex != -1 {
		A, _ = strconv.Atoi(input[AIndex+5:])
	}
}

func computeBValue() {
	// b = rand.Intn(100)
	b = 7
	B = 1
	for i := 0; i < b; i++ {
		B = (B * g) % p
	}
}

func computeSValue() {
	S = 1
	for i := 0; i < b; i++ {
		S = (S * A) % p
	}
}

func caesarCipher(s string, shift int) string {
	shift = shift % 26
	if shift < 0 {
		shift += 26
	}

	runes := []rune(s)
	for i, r := range runes {
		if r >= 'A' && r <= 'Z' {
			runes[i] = (((r + rune(shift)) - 'A') % 26) + 'A'
		} else if r >= 'a' && r <= 'z' {
			runes[i] = (((r + rune(shift)) - 'a') % 26) + 'a'
		}
	}
	return string(runes)
}

func getAndProcessResponse() {
	input := extractInput()
	decodedResponse := caesarCipher(input, -S)
	if decodedResponse == "Yeah, okay!" {
		fmt.Println(caesarCipher("Great!", S))
	} else if decodedResponse == "Let's be friends." {
		fmt.Println(caesarCipher("What a pity!", S))
	}
}

func main() {
	getGValueAndPValue()
	fmt.Println("OK")
	computeBValue()
	getAValue()
	computeSValue()
	fmt.Printf("B is %d\n", B)
	fmt.Println(caesarCipher("Will you marry me?", S))
	getAndProcessResponse()
}
