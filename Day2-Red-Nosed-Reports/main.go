package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func getData() [][]int {
	fileData, err := os.ReadFile("./sample.txt")
	if err == nil {
		var data = [][]int{}
		lines := strings.Split(string(fileData), "\n")
		for i, line := range lines {
			data = append(data, []int{})
			nums := strings.Split(line, " ")
			for _, numStr := range nums {
				if numStr == "" {
					continue
				}
				num, err := strconv.Atoi(numStr)
				if err != nil {
					panic(err)
				}
				data[i] = append(data[i], num)
			}
		}
		return data
	}
	panic(err)
}

func safetyReport() int {
	data := getData()
	safe := 0

	for _, row := range data {
		dir := "none"
		for i := 0; i < len(row)-1; i++ {
			res := row[i] - row[i+1]
			if res == 0 || res > 3 || res < -3 {
				break
			}
			if res > 0 && res <= 3 {
				if dir == "none" {
					dir = "desc"
				}
				if dir == "asc" {
					break
				}
			}
			if res < 0 && res >= -3 {
				if dir == "none" {
					dir = "asc"
				}
				if dir == "desc" {
					break
				}
			}
			if i == len(row)-2 {
				safe = safe + 1
			}

		}
	}

	return safe
}

func dampenedSafetyReport(maxStrikes int) int {
	data := getData()
	safe := 0

	for _, row := range data {
		dir := "none"
		strikes := 0
		i := 0
		for strikes < maxStrikes && dir == "done" {
			fmt.Println(row)
			for i = 0; i < len(row)-1; i++ {
				res := row[i] - row[i+1]
				if res == 0 || res > 3 || res < -3 {
					strikes = strikes + 1
					row = slices.Delete(row, i, i)
				}
				if res > 0 && res <= 3 {
					if dir == "none" {
						dir = "desc"
					}
					if dir == "asc" {
						strikes = strikes + 1
						row = slices.Delete(row, i, i)
					}
				}
				if res < 0 && res >= -3 {
					if dir == "none" {
						dir = "asc"
					}
					if dir == "desc" {
						strikes = strikes + 1
						row = slices.Delete(row, i, i)
					}
				}
				if strikes > maxStrikes {
					break
				}
				if i == len(row)-2 {
					safe = safe + 1
					dir = "done"
				}
			}
		}
	}

	return safe
}

func main() {
	fmt.Printf("Safety report: %d\n", safetyReport())
	fmt.Printf("Dampened report: %d\n", dampenedSafetyReport(1))
}
