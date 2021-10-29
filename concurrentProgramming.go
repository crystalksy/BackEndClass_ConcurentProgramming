package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type LetterFrequency struct {
	char  rune
	score int
}

var data = make(map[rune]int)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		inputString := scanner.Text()
		frequency(inputString)
		time.Sleep(1 * time.Second)
		showDataMap()
	}
}

func frequency(inputString string) {
	for _, k := range inputString {
		go frequencyChar(k, inputString)
	}
}

func frequencyChar(c rune, inputString string) {
	var isFound = false
	for k, _ := range data {
		if k == c {
			isFound = true
			break
		}
	}
	if !isFound {
		countString(c, inputString)
	}
}

func countString(c rune, inputString string) {
	var count int
	for _, k := range inputString {
		if k == c {
			count = count + 1
		}
	}
	data[c] = count
}

func showDataMap() {
	for k, v := range data {
		fmt.Printf("%c : %d\n", k, v)
	}
}
