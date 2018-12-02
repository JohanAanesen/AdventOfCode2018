package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	twos := 0
	threes := 0
	var letters []string

	for r.Scan() {
		path := r.Text()

		letters = strings.Split(path, "")

		for _, letter := range letters{
			if strings.Count(path, letter) == 3{
				threes++
				break
			}
		}

		for _, letter := range letters{
			if strings.Count(path, letter) == 2 {
				twos++
				break
			}
		}


	}

	fmt.Printf("Part 1: %v\n", twos*threes)
}
