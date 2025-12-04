package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func invalid(x int) bool {
	s := strconv.Itoa(x)
	length := len(s)
	if length%2 != 0 {
		return false
	}
	midpoint := length / 2
	firstHalf := s[:midpoint]
	secondHalf := s[midpoint:]
	return firstHalf == secondHalf
}

func invalid2(x int) bool {
	s := strconv.Itoa(x)
	length := len(s)
	if length <= 1 {
		return false // need at least two matches
	}
	for i := 1; i < length/2+1; i++ {
		if length%i == 0 {
		}
		pattern := s[:i]
		if strings.Repeat(pattern, length/i) == s {
			return true
		}
	}
	return false
}

func main() {

	//sum := 0
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	invalidSum := 0
	invalidSum2 := 0

	for fileScanner.Scan() {
		ranges := strings.Split(fileScanner.Text(), ",")
		fmt.Println(ranges)
		for _, r := range ranges {
			ids := strings.Split(r, "-")
			first, _ := strconv.Atoi(ids[0])
			last, _ := strconv.Atoi(ids[1])
			invalidRange := 0
			for i := first; i <= last; i++ {
				if invalid(i) {
					invalidRange += i

				}
				if invalid2(i) {
					invalidSum2 += i
				}
			}
			invalidSum += invalidRange
		}
	}
	fmt.Println(invalidSum)
	fmt.Println(invalidSum2)

}
