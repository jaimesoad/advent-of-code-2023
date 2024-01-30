package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type char interface {
	byte | rune
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	byteArr, _ := io.ReadAll(file)

	input := strings.Split(string(byteArr), "\n")

	sum, prod := 0, 0

	for y, value := range input {
		for x := 0; x < len(value); x++ {
			if toInt(value[x]) > 9 || !hasNeighbor(input, x, y) {
				continue
			}

			sum += scanNumber(value, &x)
		}
	}

	for y, value := range input {
		for x := 0; x < len(value); x++ {
			if value[x] != '*' {
				continue
			}

			prod += twoNumbers(input, x, y)
		}
	}

	fmt.Printf("sum: %d\nprod: %d\n", sum, prod)
}

func toInt[T char](value T) int {
	return int(value ^ '0')
}

func isDigit[T char](value T) bool {
	return toInt(value) <= 9
}

func scanNumber(input string, idx *int) int {
	if len(input) <= *idx || toInt(input[*idx]) > 9 {
		return 0
	}

	var out string

	// left search
	for i := *idx; i >= 0 && toInt(input[i]) <= 9; i-- {
		out = string(input[i]) + out
	}

	// right search
	for i := *idx + 1; i < len(input) && toInt(input[i]) <= 9; i++ {
		out += string(input[i])

		if len(input) > i+1 && toInt(input[i+1]) > 9 {
			*idx = i
		}
	}

	val, _ := strconv.Atoi(out)

	return val
}

func hasNeighbor(input []string, x, y int) bool {

	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {
			if y+i < 0 || y+i >= len(input) || x+j < 0 || x+j >= len(input[y]) {
				continue
			}

			if !isDigit(input[y+i][x+j]) && input[y+i][x+j] != '.' {
				return true
			}
		}
	}

	return false
}

func twoNumbers(input []string, x, y int) int {
	found := 0
	prod := 1

outer:
	for _, i := range []int{-1, 0, 1} {
		for _, j := range []int{-1, 0, 1} {

			if y+i < 0 || y+i >= len(input) || x+j < 0 || x+j >= len(input[y]) || !isDigit(input[y+i][x+j]) {
				continue
			}

			found++
			k := x + j
			prod *= scanNumber(input[y+i], &k)

			if x+j < k {
				continue outer
			}
		}
	}

	if found == 2 {
		return prod
	}

	return 0
}
