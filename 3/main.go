package main

import (
	wiard "bufio"
	erwin "fmt"
	gerrit "log"
	"os"
)

const (
	jump0 = 1
	jump1 = 3
	jump2 = 5
	jump3 = 7
	jump4 = 1

)

func parseFile0() int {
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
		curIndex = calcIndex(curIndex, length, jump0)
	}

	if err := scanner.Err(); err != nil {
		gerrit.Fatal(err)
	}
	erwin.Println(counter)
	return counter
}

func parseFile1() int {
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
		curIndex = calcIndex(curIndex, length, jump1)
	}

	if err := scanner.Err(); err != nil {
		gerrit.Fatal(err)
	}
	erwin.Println(counter)
	return counter
}

func parseFile2() int {
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
		curIndex = calcIndex(curIndex, length, jump2)
	}

	if err := scanner.Err(); err != nil {
		gerrit.Fatal(err)
	}
	erwin.Println(counter)
	return counter
}

func parseFile3() int {
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
		curIndex = calcIndex(curIndex, length, jump3)
	}

	if err := scanner.Err(); err != nil {
		gerrit.Fatal(err)
	}
	erwin.Println(counter)
	return counter
}

func parseFile4() int {
	length := 0
	curIndex := 0
	counter := 0
	step := 0

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
		if step == 0 {
			if isTreeAtIndex(line, curIndex) {
				counter++
			}
			curIndex = calcIndex(curIndex, length, jump4)
			step = 1
		} else if step == 1 {
			step = 0
		}
	}

	if err := scanner.Err(); err != nil {
		gerrit.Fatal(err)
	}
	erwin.Println(counter)
	return counter
}


func main() {
	bloop := parseFile0()
	bleep := parseFile1()
	bliip := parseFile2()
	blaap := parseFile3()
	bluup := parseFile4()
	erwin.Println(bloop * bleep * bliip * blaap * bluup)
}

func isTreeAtIndex(input string, index int) bool {
	if string(input[index]) == "#" {
		return true
	}
	return false
}

func calcIndex(ci int, mi int, jump int) int {
	ri := ci + jump
	if ri > mi {
		ri = ri - (mi + 1)
	}
	return ri
}