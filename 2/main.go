package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type password struct {
	lowLimit int
	highLimit int
	searchLetter string
	password string
}

func parseFile() (int, int) {

	counter := 0
	counter2 := 0

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		pw := parseLine(scanner.Text())
		if pw.validatePw() {
			counter++
		}
		if pw.validatePwMore() {
			counter2++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return counter, counter2
}


func parseLine(input string) password {

	inputSlice := strings.Split(input, " ")
	loHiSlice := strings.Split(inputSlice[0], "-")
	lowLimit, err := strconv.Atoi(loHiSlice[0])
	hiLimit, err := strconv.Atoi(loHiSlice[1])
	checkerr(err)
	pw := password{
		lowLimit: lowLimit,
		highLimit: hiLimit,
		searchLetter: strings.Trim(inputSlice[1], ":"),
		password: inputSlice[2],
	}
	return pw
}

func (p *password) validatePw() bool {
	counted := strings.Count(p.password, p.searchLetter)
	if counted > p.lowLimit-1 && counted < p.highLimit+1 {
		return true
	}
	return false
}

func (p *password) validatePwMore() bool {
	index1 := p.lowLimit-1
	index2 := p.highLimit-1


	if (string(p.password[index1]) == p.searchLetter || string(p.password[index2]) == p.searchLetter) && !(string(p.password[index1]) == string(p.password[index2])) {
		return true
	}
	return false
}

func main() {
	yep, rly := parseFile()
	fmt.Println(yep)
	fmt.Println(rly)
}

func checkerr(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)

	}
}