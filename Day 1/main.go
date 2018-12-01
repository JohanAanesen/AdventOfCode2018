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

	for r.Scan() {
		path := r.Text()


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


	fmt.Println(frequency)
}
