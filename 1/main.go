package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func parseFile() []int {

	ret := []int{}

	file, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		i, err := strconv.Atoi(scanner.Text())
		checkerr(err)
		ret = append(ret, i)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ret
}

func checkerr(e error) {
	if e != nil {
		fmt.Println(e)
		os.Exit(1)

	}
}

func calculateSliceForInt(i int, moop []int) int {
	for _, num := range moop {
		if i + num == 2020 {
			return i * num
		}
	}
	return 0
}

func main() {
	intList := parseFile()
	for _, mep := range intList {
		calc := calculateSliceForInt(mep, intList)
		if calc != 0 {
			fmt.Println(calc)
		}
	}

	for _, first := range intList {
		for _, second := range intList {
			for _, third := range intList {

				if first + second + third == 2020 {
					fmt.Println(first*second*third)
				}
			}
		}
	}
}

