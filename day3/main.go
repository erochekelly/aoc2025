package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func max(s string) int {
	j1 := 0
	j2 := 0
	i1 := 0
	for i, c := range s[:len(s)-1] {
		d, _ := strconv.Atoi(string(c))
		if d > j1 {
			j1 = d
			i1 = i
		}
	}
	for _, c := range s[i1+1 : len(s)] {
		d, _ := strconv.Atoi(string(c))
		if d > j2 {
			j2 = d
		}
	}

	fmt.Println(j1*10 + j2)
	return (j1*10 + j2)
}

func max12(s string) int {
	var j [12]int
	start := 0
	newstart := 0
	var sub string
	fmt.Println(s, len(s))
	for n := 0; n < len(j); n++ {
		maxj := 0
		if start == len(s)-1 {
			j[n], _ = strconv.Atoi(string(s[start]))
		} else {
			if n == len(j)-1 {
				sub = s[start:]
			} else {
				sub = s[start : len(s)-len(j)+n+1]
			}
			fmt.Println(start, sub)
			for i, c := range sub {
				d, _ := strconv.Atoi(string(c))
				if d > maxj {
					maxj = d
					newstart = start + i + 1

				}
			}
			start = newstart

			j[n] = maxj
		}
	}
	max := 0
	for r, i := range j {
		exp := math.Pow(10.0, float64(len(j)-r-1))
		max += i * int(exp)
	}
	fmt.Println(j, max)
	return (max)
}
func main() {

	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	total := 0
	total12 := 0

	for fileScanner.Scan() {

		total += max(fileScanner.Text())
		total12 += max12(fileScanner.Text())

	}
	fmt.Println(total)
	fmt.Println(total12)

}
