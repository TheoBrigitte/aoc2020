package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Password struct {
	LowerBound int
	UpperBound int
	Character  rune
	Password   string
}

func contestResponse() {
	scanner := bufio.NewScanner(os.Stdin)

	data := []Password{}
	for scanner.Scan() {
		line := scanner.Text()

		s := strings.Split(line, ":")

		pass := strings.TrimSpace(s[1])

		s = strings.Split(s[0], " ")

		char := rune(s[1][0])

		s = strings.Split(s[0], "-")

		lower, err := strconv.Atoi(s[0])
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}

		upper, err := strconv.Atoi(s[1])
		if err != nil {
			os.Stderr.WriteString(err.Error())
			return
		}

		p := Password{
			LowerBound: lower,
			UpperBound: upper,
			Character:  char,
			Password:   pass,
		}
		data = append(data, p)
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

func solve(data []Password) int {
	count := 0
	for _, p := range data {
		if p.isValid() {
			count++
		}
	}

	return count
}

func (p Password) isValid() bool {
	count := 0
	for _, c := range p.Password {
		if c == p.Character {
			count++
		}
	}

	return count >= p.LowerBound && count <= p.UpperBound
}

func main() {
	contestResponse()
}
