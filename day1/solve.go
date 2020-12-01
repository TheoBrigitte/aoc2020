package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func contestResponse() {
	scanner := bufio.NewScanner(os.Stdin)

	data := []int{}
	for scanner.Scan() {
		line := scanner.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}
		data = append(data, i)
	}

	err := scanner.Err()
	if err != nil {
		os.Stderr.WriteString(err.Error())
		return
	}

	result := solve(data)
	if result == 0 {
		os.Stderr.WriteString("no results")
		return
	}

	fmt.Printf("%d\n", result)
}

func solve(data []int) int {
	for _, a := range data {
		for _, b := range data {
			if a+b == 2020 {
				return a * b
			}
		}
	}

	return 0
}

func main() {
	contestResponse()
}
