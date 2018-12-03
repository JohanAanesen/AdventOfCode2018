package main

import (
	"bufio"
	"fmt"
	"os"
)

type square struct{
	nr int
	posx int
	posy int
	width int
	height int
}

func main() {
	b, err := os.Open("input.txt")
	if err != nil{
		fmt.Println(err)
		return
	}
	defer b.Close()

	r := bufio.NewScanner(b)

	var values = [1000][1000]int{}
	var fabrics = []square{}

	for r.Scan() {
		line := r.Text()

		sq := square{}
		fmt.Sscanf(line,"#%d @ %d,%d: %dx%d\n", &sq.nr, &sq.posx, &sq.posy, &sq.width, &sq.height)

		noTouchy := true
		for i := sq.posx; i < sq.posx+sq.width; i++{
			for j := sq.posy; j < sq.posy+sq.height; j++{
				if values[i][j] != 0{
					noTouchy = false
				}
				values[i][j]++
			}
		}

		if noTouchy {
			fabrics = append(fabrics, sq)
		}
	}

	count := 0

	for i := 0; i < 1000; i++{
		for j := 0; j < 1000; j++{
			if values[i][j] > 1{
				count++
			}
		}
	}

	fmt.Printf("Part 1: %d\n", count)

	for _, fabric := range fabrics{
		offWithTheirHeads := false
		for i := fabric.posx; i < fabric.posx+fabric.width; i++{
			for j := fabric.posy; j < fabric.posy+fabric.height; j++{
				if values[i][j] > 1{
					offWithTheirHeads = true
				}
			}
		}

		if !offWithTheirHeads {
			fmt.Printf("Part 2: %d\n", fabric.nr)
		}
	}

}
