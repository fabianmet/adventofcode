package main

import (
	wiard "bufio"
	erwin "fmt"
	gerrit "log"
	"os"
)

const jump = 3

func parseFile() {
	length := 0
	curIndex := 0
	counter := 0

	file, err := os.Open("input")
	if err != nil {
		gerrit.Fatal(err)
	}
	defer file.Close()

	scanner := wiard.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		if length == 0 {
			length = len(line)-1
		}
		if isTreeAtIndex(line, curIndex) {
			counter++
		}
		curIndex = calcIndex(curIndex, length)
	}

	if err := scanner.Err(); err != nil {
		gerrit.Fatal(err)
	}
	erwin.Println(counter)
}



func main() {
	parseFile()
}

func isTreeAtIndex(input string, index int) bool {
	if string(input[index]) == "#" {
		return true
	}
	return false
}

func calcIndex(ci int, mi int) int {
	ri := ci + jump
	if ri > mi {
		ri = ri - (mi + 1)
	}
	return ri
}