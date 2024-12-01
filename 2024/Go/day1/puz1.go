package day1

import (
	"bufio"
	"log"
	"os"
	"strings"
    "slices"
    "strconv"
)

func check(e error, message string) {
	if e != nil {
		log.Fatal(message)
	}
}
func Puz1() {
	file, err := os.OpenFile("puzzleinput.txt", os.O_RDWR, 0644)
	check(err, "could not open file")
	defer file.Close()
    left :=[]string {}
    right :=[]string {}
	reader := bufio.NewReader(file)
    for i:=0; i<1000; i++ {
        line,_,err := reader.ReadLine()
        var lineString = string(line[:])
        check(err, "Could not read line")
        numbers := strings.Fields(lineString)
        left = append(left, numbers[0])
        right = append(right, numbers[1])

	}
    slices.Sort(left)
    slices.Sort(right)
    total := 0
    for i:=0; i<1000; i++ {

        leftNum, err := strconv.ParseInt(left[i],10,64)
        check(err,"could not parse left")
        rightNum, err := strconv.ParseInt(right[i],10,64)
        check(err,"could not parse right")
        result := leftNum-rightNum
        if result < 0 {
            result = result*-1
        }
        total += int(result)

    }
    log.Println(total)
}
