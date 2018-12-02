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
	var list []string

	for r.Scan() {
		path := r.Text()

		list = append(list, path)

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

	for i, line := range list{
		for j, line2 := range list{
			if i != j{
				checkWord(line, line2)
			}
		}
	}
}

func checkWord(word1 string, word2 string){
	w1 := strings.Split(word1, "")
	w2 := strings.Split(word2, "")

	diff := 0

	for i, letter := range w2{
		if w1[i] != letter{
			diff++
			w1[i] = ""
			w2[i] = ""
			if diff > 1{
				return
			}
		}
	}

	fmt.Printf("Part 2: %v\n", strings.Join(w1, ""))

	os.Exit(0) //quit
}
