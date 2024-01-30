package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

var numbers = map[string]int64{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

func main() {
	var input []string

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	byteValue, _ := io.ReadAll(file)

	input = strings.Split(string(byteValue), "\n")

	timer := time.Now()

	val := calibrate(input)

	stop := time.Since(timer)

	fmt.Println(stop)
	fmt.Println(val)
}

func toNumber(value string, idx int) int64 {
	for i := 3; idx+i < len(value) && i <= 5; i++ {
		if num, ok := numbers[value[idx:idx+i]]; ok {
			return num
		}
	}

	return 0
}

func toNumberReversed(value string, idx int) int64 {
	for i := 2; idx-i >= 0 && i <= 5; i++ {
		if num, ok := numbers[value[idx-i:idx+1]]; ok {
			return num
		}
	}

	return 0
}

func toInt(value byte) int {
	return int(value ^ '0')
}

func getDigits(value string) int64 {
	var out int64
	left, right := 0, len(value)-1
	foundLeft, foundRight := false, false

	for !(foundLeft && foundRight) {

		if !foundLeft {
			out += search(value, &foundLeft, &left, 1, toNumber) * 10
		}

		if !foundRight {
			out += search(value, &foundRight, &right, -1, toNumberReversed)
		}
	}

	return out
}

func calibrate(input []string) int64 {
	var out int64

	for _, value := range input {
		out += getDigits(value)
	}

	return out
}

func search(value string, found *bool, idx *int, next int, searchFn func(string, int) int64) int64 {
	var num int64

	if toInt(value[*idx]) > 9 {
		num = searchFn(value, *idx)
	}

	if num == 0 && toInt(value[*idx]) > 9 {
		*idx += next
		return 0
	}

	*found = true

	if num != 0 {
		return num
	}

	return int64(value[*idx] ^ '0')
}	