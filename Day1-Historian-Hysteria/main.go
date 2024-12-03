package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func intAbs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func listDistance (col1 []int, col2 []int) int {
    var res []int

	slices.Sort(col1)
	slices.Sort(col2)

	for i := range col1 {
		res = append(res, intAbs(col1[i]-col2[i]))
	}

	diff := 0

	for _, v := range res {
		diff += v
	}

    return diff
}

func similarityScore (col1 []int, col2 []int) int {
    res := 0
    for _, v := range col1 {
        count := 0 
        for _, k := range col2 {
            if v == k {
                count++
            }
        }
        res += (v * count)
    }
    return res
}

func main() {
	records := readCsvFile("./sample.csv")
	var col1 []int
	var col2 []int

    for i := range records {
		num, err := strconv.Atoi(records[i][0])
		if err == nil {
			col1 = append(col1, num)
		}
		num, err = strconv.Atoi(records[i][1])
		if err == nil {
			col2 = append(col2, num)
		}
	}

    dist := listDistance(col1, col2)
    sim := similarityScore(col1, col2)
	
    fmt.Printf("List distance:    %d\n", dist)
    fmt.Printf("Similarity score: %d\n", sim)
}
