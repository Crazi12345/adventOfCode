package puzzles

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func Day2Puz2() {
	defer duration(track("Day2Puz2"))
	file, err := os.OpenFile("puzzleinputDay2.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
	reader := bufio.NewReader(file)
	safeCounter := 0
	for i := 0; i < 1000; i++ {
		line, _, err := reader.ReadLine()
		var lineString = string(line[:])
		check(err, "Could not read line")
		allNumbers := strings.Fields(lineString)

		hadASafe := false
		for k := 0; k < len(allNumbers); k++ {
            isBad:=false
			tempNumbers := make([]string, len(allNumbers))
			copy(tempNumbers, allNumbers)
			numbers := append(tempNumbers[:k], tempNumbers[k+1:]...)
			first, _ := strconv.ParseInt(numbers[0], 10, 64)
			second, _ := strconv.ParseInt(numbers[1], 10, 64)
			variance := first - second
			if variance == 0 {
				continue
			} else if variance < 0 && variance > -4 {
				variance = -1
			} else if variance > 0 && variance < 4 {
				variance = 1
			} else {
				continue
			}

			for j := 2; j < len(numbers); j++ {
				numA, _ := strconv.ParseInt(numbers[j-1], 10, 64)
				numB, _ := strconv.ParseInt(numbers[j], 10, 64)
				result := (numA - numB) * variance
				if result < 1 || result > 3 {
                    isBad = true
					break
				}
			}
            if !isBad {
			hadASafe = true
            break
            }
		}
		if hadASafe {
			safeCounter++
		} else {
		}
	}
	//		leftNum, err := strconv.ParseInt(left[i], 10, 64)

	log.Println(safeCounter)
}
