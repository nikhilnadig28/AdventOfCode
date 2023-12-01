package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func main() {

	file, err := os.Open("1.txt")
	if err != nil {
		fmt.Println("Something is wrong with the input")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	re := regexp.MustCompile(`\d`)
	var total int
	for scanner.Scan() {
		line := scanner.Text()
		numbers := re.FindAllString(line, -1)
		fmt.Println(line)

		if len(numbers) > 0 {
			firstNumber := numbers[0]
			lastNumber := numbers[len(numbers)-1]

			rowNumber, err := strconv.Atoi(firstNumber + lastNumber)
			total += rowNumber

			if err != nil {
				fmt.Println("Error: ", err)
			}
		}

		fmt.Println("total: ", total)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error:", err)
	}
}
