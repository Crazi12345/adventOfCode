package puzzles
import (
"time"
"log"
)
func check(e error, message string) {
	if e != nil {
		log.Fatal(message)
	}
}
func track(msg string) (string, time.Time) {
	return msg, time.Now()
}

func duration(msg string, start time.Time) {
	log.Printf("%v: %v\n", msg, time.Since(start))

}
