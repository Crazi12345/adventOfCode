package puzzles

import (
	"bufio"
	"log"
	"os"
	"strconv"
	// "strconv"
)

func Day3Puz2() {
	defer duration(track("Day3Puz2"))
	file, err := os.OpenFile("InputDay3.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
	reader := bufio.NewReader(file)
	sum := 0

	for i := 0; i < 200; i++ {
		data, _, err := reader.ReadLine()
		check(err, "Could not read line")
		largeNum := 0

		maxVal:= [12]int{0,0,0,0,0,0,0,0,0,0,0,0}

		lastfoundIndex := -1
		for j := 0; j < len(maxVal); j++ {


			for k := lastfoundIndex+1; k <= len(data)-(len(maxVal)-j); k++ {

				if int(maxVal[j]) < int(data[k]) {
					maxVal[j] = int(data[k])
					lastfoundIndex = k
				//	log.Println("Here:"+string(int(maxVal[j])))
				}
			}
		}

		result:=""
		for steps:=0; steps<len(maxVal); steps++ {
			result+=string(int(maxVal[steps]))
			//log.Print(result)
		}
		largeNum, err = strconv.Atoi(string(result))
		check(err, "could not handle the number")
		sum += largeNum
	}
	log.Print(sum)
}
