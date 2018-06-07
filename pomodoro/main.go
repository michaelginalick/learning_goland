package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"./digits"
	"./transform"
	tm "github.com/buger/goterm"
)

const finish = "Finish Time"
const current = "Current Time"

func main() {
	second := 59
	totalTime := validateInput()

	tm.Clear()
	tm.MoveCursor(100, 1)

	startTime := currentTime()
	finishTime := startTime.Add(time.Minute * totalTime)
	print(finish, finishTime)

	for {
		curTime := currentTime()
		tm.MoveCursor(1, 1)
		print(current, curTime)

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
		second = maintainSeconds(second)
	}
	fmt.Println("Your session ended at ", currentTime().Format(time.RFC1123))
}

func validateInput() time.Duration {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Please enter a number between 1 and 60: ")
	input, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println(err)
	}
	clockTime, err := strconv.Atoi(strings.TrimSpace(input))

	if err != nil || (clockTime <= 0 || clockTime >= 60) {
		fmt.Println("Oops!")
		validateInput()
	}

	return time.Duration(clockTime)
}

func shouldBreakLoop(curTime time.Time, finishTime time.Time) bool {
	if curTime.Format(time.RFC1123) > finishTime.Format(time.RFC1123) {
		return true
	}
	return false
}

func maintainSeconds(second int) int {
	second--
	if second <= 0 {
		second = 59
	}
	return second
}

func currentTime() time.Time {
	return time.Now()
}

func print(s string, curTime time.Time) {
	tm.Println(s, curTime.Format(time.RFC1123))
}
