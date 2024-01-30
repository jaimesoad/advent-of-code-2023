package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type color struct {
	red, green, blue int
}

var games = make(map[int][]color)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	byteArr, _ := io.ReadAll(file)

	input := strings.Split(string(byteArr), "\n")

	for _, i := range input {
		gm := strings.Split(i, ":")

		var newGame []color

		colors := strings.Split(strings.TrimSpace(gm[1]), ";")

		for _, j := range colors {
			individualColors := strings.Split(j, ",")

			var try color

			for _, k := range individualColors {
				ic := strings.Split(strings.TrimSpace(k), " ")
				col, _ := strconv.Atoi(ic[0])

				switch ic[1] {
				case "red":
					try.red = col

				case "green":
					try.green = col

				case "blue":
					try.blue = col
				}
			}

			newGame = append(newGame, try)
		}

		id, _ := strconv.Atoi(strings.Split(gm[0], " ")[1])

		games[id] = newGame
	}

	var idSum, powerSum int

	for id := 1; id <= len(games); id++ {
		g := games[id]
		viable := true
		var lowest color

		for _, i := range g {

			if lowest.red < i.red {
				lowest.red = i.red
			}

			if lowest.green < i.green {
				lowest.green = i.green
			}
			
			if lowest.blue < i.blue {
				lowest.blue = i.blue
			}

			viable = viable && i.red <= 12 && i.green <= 13 && i.blue <= 14
		}

		powerSum += lowest.red * lowest.green * lowest.blue

		if !viable {
			continue
		}

		idSum += id
	}

	fmt.Println("Sum if ID's", idSum)
	fmt.Println("Sum of powers", powerSum)
}
