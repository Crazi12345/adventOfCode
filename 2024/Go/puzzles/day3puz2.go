package puzzles

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day3Puz2() {
	defer duration(track("Day2Puz2"))
	file, err := os.OpenFile("puzzleinputDay3.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
	reader := bufio.NewReader(file)
	total := 0
    lineString := ""
	for i := 0; i < 6; i++ {
		line, _, err := reader.ReadLine()
		check(err, "Could not read line")
	    lineString += string(line[:])
    }
		allText := strings.Split(lineString, "do()")
		for k := range allText {
			excludedText := strings.Split(allText[k], "don't()")
			allNumbers := strings.Split(excludedText[0], "mul(")

			for j := range allNumbers {
				parsedNum := strings.Split(allNumbers[j], ")")
				onlyNumbers := strings.Split(parsedNum[0], ",")
				if len(onlyNumbers) == 2 {
					if len(onlyNumbers[0]) > 0 && len(onlyNumbers[1]) <= 3 {
						if len(onlyNumbers[0]) <= 3 && len(onlyNumbers[1]) > 0 {
							numA, _ := strconv.ParseInt(onlyNumbers[0], 10, 64)
							numB, _ := strconv.ParseInt(onlyNumbers[1], 10, 64)
							result := numA * numB
							total += int(result)
						}
					}

				}
			}
	}
	log.Println(total)
}
