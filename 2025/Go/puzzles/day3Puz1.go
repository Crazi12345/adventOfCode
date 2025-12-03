package puzzles

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func Day3Puz1() {
	defer duration(track("Day3Puz1"))
	file, err := os.OpenFile("InputDay3.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
	reader := bufio.NewReader(file)
	sum := 0

	for i := 0; i < 200; i++ {
		data, _, err := reader.ReadLine()
		check(err, "Could not read line")
		largeNum := 0
		for j := 0; j < len(data); j++ {
			for k := j+1; k < len(data); k++ {

				num, err := strconv.Atoi(string(data[j]) + string(data[k]))
				check(err, "could not convert")
				if num > largeNum {
					// log.Print("new largest Num")
					// log.Println()
					largeNum = num
					// log.Print(largeNum)
				}

			}
		}

		sum += largeNum
	}
	log.Print(sum)
}
