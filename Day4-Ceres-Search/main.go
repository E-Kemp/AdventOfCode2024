package main

import (
	"fmt"
	"os"
	"strings"
)

type Token int

func getData() []string {
    fileData, err := os.ReadFile("./sample.txt")
	if err == nil {
		return strings.Split(string(fileData), "\n")
	}
	panic(err)
}

func partOne(data []string) int {
	total := 0
	for y := 0; y < len(data); y++ {
		for x := 0; x < len(data[y]); x++ {
			if data[y][x] == 'X' {
				if x < len(data[y])-3 {
					if data[y][x+1] == 'M' && data[y][x+2] == 'A' && data[y][x+3] == 'S' {
						total++
					}
					if y < len(data)-4 && data[y+1][x+1] == 'M' && data[y+2][x+2] == 'A' && data[y+3][x+3] == 'S' {
						total++
					}
				}
				if x >= 3 {
					if data[y][x-1] == 'M' && data[y][x-2] == 'A' && data[y][x-3] == 'S' {
						total++
					}
					if y >= 3 && data[y-1][x-1] == 'M' && data[y-2][x-2] == 'A' && data[y-3][x-3] == 'S' {
						total++
					}
				}
				if y < len(data)-4 {
					if data[y+1][x] == 'M' && data[y+2][x] == 'A' && data[y+3][x] == 'S' {
						total++
					}
					if x >= 3 && data[y+1][x-1] == 'M' && data[y+2][x-2] == 'A' && data[y+3][x-3] == 'S' {
						total++
					}
				}
				if y >= 3 {
					if data[y-1][x] == 'M' && data[y-2][x] == 'A' && data[y-3][x] == 'S' {
						total++
					}
					if x < len(data[y])-3 && data[y-1][x+1] == 'M' && data[y-2][x+2] == 'A' && data[y-3][x+3] == 'S' {
						total++
					}
				}
			}
		}
	}

	return total
}

func partTwo(data []string) int {
	total := 0
	for y := 1; y < len(data)-2; y++ {
		for x := 1; x < len(data[y])-1; x++ {
			if data[x][y] == 'A' {
				if data[x-1][y-1] == 'M' && data[x+1][y-1] == 'M' && data[x+1][y+1] == 'S' && data[x-1][y+1] == 'S' {
					total++
				}
                if data[x-1][y-1] == 'M' && data[x+1][y-1] == 'S' && data[x+1][y+1] == 'S' && data[x-1][y+1] == 'M' {
					total++
				}
                if data[x-1][y-1] == 'S' && data[x+1][y-1] == 'S' && data[x+1][y+1] == 'M' && data[x-1][y+1] == 'M' {
					total++
				}
                if data[x-1][y-1] == 'S' && data[x+1][y-1] == 'M' && data[x+1][y+1] == 'M' && data[x-1][y+1] == 'S' {
					total++
				}
			}
		}
	}

	return total
}

func main() {
	data := getData()

	ans1 := partOne(data)
	ans2 := partTwo(data)

	fmt.Printf("Total 1: %d\n", ans1)
	fmt.Printf("Total 2: %d\n", ans2)

}
