package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	b, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var input []string

	for r.Scan() {
		line := r.Text()

		input = strings.Split(line, "")
	}

	for len(input) > 10{
		for i:=0; i < len(input)-1; i++{
			if input[i] == strings.ToLower(input[i]){
				if input[i+1] == strings.ToUpper(input[i+1]) && strings.ToUpper(input[i]) ==  strings.ToUpper(input[i+1]){
					input = append(input[:i], input[i+2:]...)

				}
			}else if input[i] == strings.ToUpper(input[i]){
				if input[i+1] == strings.ToLower(input[i+1]) && strings.ToUpper(input[i]) ==  strings.ToUpper(input[i+1]){
					input = append(input[:i], input[i+2:]...)

				}
			}
		}
		fmt.Println(len(input))
	}



	fmt.Println(input)
}
