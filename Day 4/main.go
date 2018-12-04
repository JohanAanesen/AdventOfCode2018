package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	type StringSlice []string

	var values = []string{}

	for r.Scan() {
		line := r.Text()

		dateString := ""
		timeString := ""
		guardString := ""
		nrString := ""

		fmt.Sscanf(line, "[%10s %5s]%s %s\n", &dateString, &timeString, &guardString, &nrString)

		values = append(values, dateString+" "+timeString+" "+guardString+" "+nrString)

	}
	sort.Strings(values)

	timeFormat := "2006-01-02 15:04 MST"
	var guardsSleepTable = make(map[string][]int)
	var guardsSleep = make(map[string]int)

	currentGuard := ""
	var sleepTime time.Time
	var wakeUpTime time.Time
	sleepMinute := 0

	for _, val := range values {
		vals := strings.Split(val, " ")
		if vals[2] == "Guard" {
			currentGuard = vals[3]
			if guardsSleepTable[currentGuard] == nil {
				guardsSleepTable[currentGuard] = make([]int, 60)
			}
		} else if vals[2] == "wakes" {
			wakeUpTime, _ = time.Parse(timeFormat, vals[0]+" "+vals[1]+" UTC")

			diff := int(wakeUpTime.Sub(sleepTime).Minutes())

			guardsSleep[currentGuard] += diff

			for i := 0; i < diff; i++ {
				if sleepMinute == 60 {
					sleepMinute = 0
				}
				guardsSleepTable[currentGuard][sleepMinute]++

				sleepMinute++
			}
		} else if vals[2] == "falls" {
			sleepTime, _ = time.Parse(timeFormat, vals[0]+" "+vals[1]+" UTC")
			sleepMinute = sleepTime.Minute()
		}
	}

	sleepsTheMost := ""
	timeSlept := 0
	minuteSleptMost := 0
	theMinute := 0
	for i, val := range guardsSleep {
		if val > timeSlept {
			timeSlept = val
			sleepsTheMost = i
		}
	}

	for i := 0; i < 60; i++ {
		if guardsSleepTable[sleepsTheMost][i] > minuteSleptMost {
			minuteSleptMost = guardsSleepTable[sleepsTheMost][i]
			theMinute = i
		}
	}

	sleepsTheMost = strings.TrimPrefix(sleepsTheMost, "#")

	guardNr, _ := strconv.Atoi(sleepsTheMost)

	fmt.Println(guardNr)
	fmt.Println(theMinute)
	fmt.Printf("Part 1: %v\n", guardNr*theMinute)

	mostSleepAccurateGuard := ""
	minsSlept := 0
	minInQuestion := 0
	for i, val := range guardsSleepTable {
		for j, val2 := range val {
			if val2 > minsSlept {
				mostSleepAccurateGuard = i
				minsSlept = val2
				minInQuestion = j
			}
		}
	}

	mostSleepAccurateGuard = strings.TrimPrefix(mostSleepAccurateGuard, "#")
	guardNr2, _ := strconv.Atoi(mostSleepAccurateGuard)

	fmt.Println(guardNr2)
	fmt.Println(minInQuestion)
	fmt.Printf("Part 2: %v\n", guardNr2*minInQuestion)
}
