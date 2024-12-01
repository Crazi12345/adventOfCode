package main

import (
	"aoc/2024/day1"
	"log"
	"time"
)

func main() {

    start := time.Now()
	day1.Puz2()
    duration := time.Since(start)
    log.Println(duration.Microseconds())
	day1.Puz1()
}
