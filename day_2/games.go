package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func valid_game(input string) (bool, int, error) {

	blue := int(14)
	red := int(12)
	green := int(13)

	data := strings.Split(input, ":")
	game_int, err := strconv.Atoi(strings.Replace(data[0], "Game ", "", 1))
	if err != nil {
		log.Fatal(err)
		return false, 0, err
	}
	games := strings.Split(data[1], ";")

	for _, game := range games {
		marbels := strings.Split(game, ",")
		for _, marble := range marbels {

			count, err := strconv.Atoi(strings.Split(marble, " ")[1])
			if err != nil {
				// ... handle error
				log.Fatal(err)
				return false, 0, err
			}

			switch {
			case strings.Contains(marble, "blue"):
				if count > blue {
					return false, 0, nil
				}
			case strings.Contains(marble, "red"):
				if count > red {
					return false, 0, nil
				}
			case strings.Contains(marble, "green"):
				if count > green {
					return false, 0, nil
				}
			}
		}
	}
	return true, game_int, nil
}

func min_cubes(input string) (int, error) {

	blue := int(0)
	red := int(0)
	green := int(0)

	data := strings.Split(input, ":")
	games := strings.Split(data[1], ";")

	for _, game := range games {
		marbels := strings.Split(game, ",")
		for _, marble := range marbels {

			count, err := strconv.Atoi(strings.Split(marble, " ")[1])
			if err != nil {
				// ... handle error
				log.Fatal(err)
				return 0, err
			}

			switch {
			case strings.Contains(marble, "blue"):
				if count > blue {
					blue = count
				}
			case strings.Contains(marble, "red"):
				if count > red {
					red = count
				}
			case strings.Contains(marble, "green"):
				if count > green {
					green = count
				}
			}
		}
	}
	return (red * green * blue), nil
}

func main() {
	log.SetPrefix("Games: ")
	log.SetFlags(0)

	b, err := os.ReadFile("input_data.txt")
	// can file be opened?
	if err != nil {
		log.Fatal(err)
	}

	// convert bytes to string
	raw_str := string(b)
	split := strings.Split(raw_str, "\n")

	total := int(0)

	// for _, input := range split {
	// 	val, num, err := valid_game(input)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	} else {
	// 		if val {
	// 			total += num
	// 		}
	// 	}
	// }

	for _, input := range split {
		num, err := min_cubes(input)
		if err != nil {
			log.Fatal(err)
		} else {
			total += num
		}
	}
	log.Println(total)
}
