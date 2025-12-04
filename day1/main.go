package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	position := 50
	count := 0
	readFile, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	for fileScanner.Scan() {
		dir := string(fileScanner.Text()[0])
		inputvalue, _ := strconv.Atoi(fileScanner.Text()[1:])
		value := inputvalue % 100
		count += (inputvalue - value) / 100
		fmt.Println(position, count, dir, value)
		if dir == "L" {
			if value > position {
				if position != 0 {
					count++
				}
				position = 100 - (value - position)

			} else {
				position -= value
			}
		} else {
			position += value
			if position > 99 {
				position -= 100
				if position != 0 {
					count += 1
				}
			}

		}
		if position == 0 {
			count++
		}
	}
	fmt.Println(count)
}
