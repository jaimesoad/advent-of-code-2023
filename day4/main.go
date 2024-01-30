package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
	"strings"
)

var copies = make(map[int]int)

func loadFile(file string) ([]string, error) {
	blob, err := os.Open(file)
	if err != nil {
		return []string{}, err
	}
	defer blob.Close()

	byteArr, _ := io.ReadAll(blob)
	return strings.Split(string(byteArr), "\n"), nil
}

func main() {
	input, err := loadFile("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	points := 0

	for idx, cards := range input {
		card := strings.Split(strings.TrimSpace(strings.Split(cards, ":")[1]), "|")
		matches := 0

		copies[idx+1] += 1

		winner := convertToInt(strings.Split(strings.TrimSpace(card[0]), " "))
		held := convertToInt(strings.Split(strings.TrimSpace(card[1]), " "))

		// fmt.Println(len(winner), ",", len(held), "|", winner, "|", held,)

		for _, number := range held {
			if slices.Contains(winner, number) {
				matches++
			}
		}

		if matches > 0 {
			points += 1 << (matches - 1)

			for i := idx + 2; i <= idx+matches+1; i++ {
				copies[i] += copies[idx+1]
			}
		}
	}

	sum := 0

	for _, val := range copies {
		sum += val
	}

	fmt.Println(points)
	fmt.Println(sum)
}

func convertToInt(input []string) []int {
	var out []int

	for _, val := range input {

		num, err := strconv.Atoi(val)

		if err != nil {
			continue
		}

		out = append(out, num)
	}

	return out
}
