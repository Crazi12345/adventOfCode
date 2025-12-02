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
	file, err := os.OpenFile("InputDay2.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
	log.SetOutput(os.Stdout)
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
		for j := minval; j <= maxval; j++ {
			id := strconv.Itoa(j)
			for k := 2; k <= len(id); k++ {
				step := len(id) / k
				if len(id)%k > 0 {
					continue
				}
				lastSeen := id[:step]
				isEqual := true
				for m := step; m < len(id); m = m + step {

					if id[m:m+step] != lastSeen {
						isEqual = false
						break
					}
					lastSeen = id[m : m+step]
				}
				if isEqual == true {
					// log.Println(lastSeen + " hit:" + id)
					invalidSum += j
					break
				}
			}
		}
	}
	log.Println(invalidSum)
}
