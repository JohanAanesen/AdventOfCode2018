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
	var input2 []string
	var inputString string

	for r.Scan() {
		line := r.Text()

		input = strings.Split(line, "")
		input2 = strings.Split(line, "")
		inputString = line
	}

	part1 := retract(input)
	fmt.Printf("Part 1: %v\n", part1)

	var types = make(map[string]bool)
	var results = make(map[string]int)

	for _, val := range input2{
		types[strings.ToLower(val)] = true
	}

	for val, _ := range types{
		tempString := strings.Replace(inputString, val, "", -1)
		tempString = strings.Replace(tempString, strings.ToUpper(val), "", -1)
		var tempArr []string
		tempArr = strings.Split(tempString, "")

		/*
		for i:=0; i < len(tempArr); i++{
			if strings.ToLower(tempArr[i]) == val{
				tempArr = append(tempArr[:i], tempArr[i+1:]...) //delete all instances of the type
			}
		}*/

		result := retract(tempArr)
		//fmt.Printf("Val: %s, AND STRING: %s\n", val, tempString)
		results[val] = result
	}

	smallest := 999999
	removedType := ""
	for i, val := range results{
		if val < smallest{
			smallest = val
			removedType = i
		}
	}

	fmt.Printf("Part 2: Removed %s, retracted length now %v\n", removedType, smallest)

}

func retract(slice []string)int{
	retracted := true

	for retracted{
		finished := true
		for i:=0; i < len(slice)-1; i++{
			if slice[i] == strings.ToLower(slice[i]){
				if slice[i+1] == strings.ToUpper(slice[i+1]) && strings.ToUpper(slice[i]) ==  strings.ToUpper(slice[i+1]){
					slice = append(slice[:i], slice[i+2:]...)
					finished = false
				}
			}else if slice[i] == strings.ToUpper(slice[i]){
				if slice[i+1] == strings.ToLower(slice[i+1]) && strings.ToUpper(slice[i]) ==  strings.ToUpper(slice[i+1]){
					slice = append(slice[:i], slice[i+2:]...)
					finished = false
				}
			}
		}
		if finished{
			return len(slice)
			retracted = false
		}

	}

	return 0
}