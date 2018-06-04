package main

import (
	"time"
	"strconv"
	"./digits"
	"./transform"
	tm "github.com/buger/goterm"
)

func main() {
	second := 59
	tm.Clear()
	tm.MoveCursor(110, 1)
	startTime := time.Now()
	
	totalTime := time.Duration(25)
	finishTime := startTime.Add(time.Minute * totalTime)
	tm.Println("Finish Time:", finishTime.Format(time.RFC1123))


	for {
		curTime := time.Now()
		tm.MoveCursor(1, 1)
		tm.Println("Current Time:", curTime.Format(time.RFC1123))

		if shouldBreakLoop(curTime, finishTime) {
			break
		}

		minuteTens, minuteOnes := transform.ExtractNumbers(int(time.Until(finishTime).Minutes()))
		secondTens, secondOnes := transform.ExtractNumbers(second)
	
		transform.FormatNumber(digits.GetDigits(strconv.Itoa(minuteTens)), 
													digits.GetDigits(strconv.Itoa(minuteOnes)), 
													digits.GetDigits(":"), 
													digits.GetDigits(strconv.Itoa(secondTens)), 
													digits.GetDigits(strconv.Itoa(secondOnes)))

		tm.Flush()
		
		time.Sleep(time.Second)
	}

}


func shouldBreakLoop(curTime time.Time, finishTime time.Time) bool {
	if curTime.Format(time.RFC1123) > finishTime.Format(time.RFC1123) {
		return true
	}
	return false
}

func maintainSeconds(s int) int{
	s--
		if s <= 0 {
			s = 59
		}
		return s
}
