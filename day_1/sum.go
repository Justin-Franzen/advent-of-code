package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func find_num(input string) (int, error) {
	re := regexp.MustCompile("[0-9]")
	r := strings.NewReplacer("one", "o1e", "two", "t2o", "three", "t3e", "four",
		"f4r", "five", "f5e", "six", "s6x", "seven", "s7n", "eight", "e8t", "nine", "n9e")
	nums := re.FindAllString(r.Replace(r.Replace(input)), -1)
	//log.Println(nums)
	end_int, err := strconv.Atoi(nums[0] + nums[len(nums)-1])
	//log.Println(end_int)
	if err != nil {
		// ... handle error
		log.Fatal(err)
		return 0, err
	}

	return end_int, nil
}

func main() {
	log.SetPrefix("Sum: ")
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

	for _, input := range split {
		num, err := find_num(input)
		if err != nil {
			log.Fatal(err)
		} else {
			total += num
		}
	}
	log.Println(total)

}
