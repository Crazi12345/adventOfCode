package puzzles

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2Puz1() {
	defer duration(track("Day2Puz1"))
	file, err := os.OpenFile("InputDay2.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
	reader := bufio.NewReader(file)
	data, _, err := reader.ReadLine()
	dataString := string(data)
	check(err, "could not read line")
	ranges := strings.Split(dataString, ",")

	invalidSum := 0
	for i := 0; i < len(ranges); i++ {
		values := strings.Split(ranges[i], "-")
		minval, err := strconv.Atoi(values[0])
		check(err, "could not convert1")
		maxval, err := strconv.Atoi(values[1])
		check(err, "could not convert2")
		for j := minval; j < maxval; j++ {
			id := strconv.Itoa(j)
			mid := len(id) / 2
			if id[:mid] == id[mid:] {
				invalidSum = invalidSum + j
			}

		}

	}

	log.Println(invalidSum)
}
