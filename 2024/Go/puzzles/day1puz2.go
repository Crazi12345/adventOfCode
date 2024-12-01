package puzzles

import (
	"bufio"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Day1Puz2() {
	defer duration(track("Day1Puz2"))
	file, err := os.OpenFile("puzzleinput.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
	left := []string{}
	right := []string{}
	reader := bufio.NewReader(file)
	for i := 0; i < 1000; i++ {
		line, _, err := reader.ReadLine()
		var lineString = string(line[:])
		check(err, "Could not read line")
		numbers := strings.Fields(lineString)
		left = append(left, numbers[0])
		right = append(right, numbers[1])

	}
	slices.Sort(left)
	slices.Sort(right)
	total := 0
	for i := 0; i < 1000; i++ {

		leftNum, _ := strconv.ParseInt(left[i], 10, 64)
		isFound := false
		count := 0
		for j := 0; j < 1000; j++ {
			if left[i] == right[j] {
				count++
			}
			if isFound {
				break
			}

		}
		total += int(leftNum) * count
	}
	log.Println(total)
}
