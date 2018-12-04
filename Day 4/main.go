package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil{
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

		//t, _ := time.Parse("1518-10-06 00:50", "2006-01-02T15:04:05Z")

		fmt.Printf("Test: %s\n", guardString)
		values = append(values, dateString+" "+timeString+" "+guardString+" "+nrString)

	}
	sort.Strings(values)

	fmt.Printf("TEST: %v\n", values)
}
