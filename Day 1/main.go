package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

	var frequency int
	var numbers []string

	for r.Scan() {
		path := r.Text()

		numbers = append(numbers, path)

		if strings.HasPrefix(path, "+"){
			path = strings.TrimPrefix(path, "+")

			nr, _ := strconv.Atoi(path)

			frequency += nr

		}else if strings.HasPrefix(path, "-"){
			path = strings.TrimPrefix(path, "-")

			nr, _ := strconv.Atoi(path)

			frequency -= nr

		}

	}

	//PART 1:
	fmt.Printf("Part 1: %v\n", frequency)

	found := 0
	var freqs = make(map[int]bool)
	frequency1 := 0

	for found < 1{
		for i := 0; i < len(numbers); i++{
			if strings.HasPrefix(numbers[i], "+"){
				temp := strings.TrimPrefix(numbers[i], "+")

				nr, _ := strconv.Atoi(temp)

				frequency1 += nr

			}else if strings.HasPrefix(numbers[i], "-"){
				temp := strings.TrimPrefix(numbers[i], "-")

				nr, _ := strconv.Atoi(temp)

				frequency1 -= nr


			}

			if freqs[frequency1] == false{
				freqs[frequency1] = true
			}else{
				fmt.Printf("Part 2: %v\n",frequency1)
				found++
				os.Exit(0)
			}
		}
	}
}
