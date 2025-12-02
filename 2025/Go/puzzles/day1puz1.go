package puzzles

import (
	"bufio"
	"log"
	"os"
	"strconv"
)


func Day1Puz1() {
	defer duration(track("Day1Puz1"))
	file, err := os.OpenFile("InputDay1Puz1.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
	reader := bufio.NewReader(file)
	dial := 50
	password := 0
	dir := 1
	for i := 0; i<4831; i++ {
		code,_,err := reader.ReadLine()
		if err != nil { break;}

		if code[0] == 'R' {
			dir = 1 }
		if code[0] == 'L' {
			dir = -1 }
		codeVal, err := strconv.Atoi(string(code[1:]))
		check(err, "could not convert to int")
		dial = (dial +(codeVal*dir)) % 100
		if (dial == 0) {password++}

	}
	log.Print(password)


}


