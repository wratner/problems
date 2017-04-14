package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
)

const delimeter = ";"

func main() {
	fileName := flag.String("fileName", "", "Name of file to be processed")
	flag.Parse()
	if *fileName == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	sequences, err := getSequences(*fileName)
	if err != nil {
		log.Fatalln(err.Error())
	}

	subsequences, err := getSubsequences(sequences)
	if err != nil {
		log.Fatalln(err.Error())
	}

	for _, values := range subsequences {
		fmt.Println(values)
	}
}

// getSequences takes in the file name and parses out the sequence pairs.
// It trims off any whitespace around the sequences and ignores empty lines/invalid inputs.
func getSequences(fileName string) ([]string, error) {
	sequences := []string{}
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) != "" || strings.Contains(strings.TrimSpace(scanner.Text()), delimeter) {
			sequences = append(sequences, strings.TrimSpace(scanner.Text()))
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return sequences, nil
}

// getSubsequences takes a slice of sequences and splits them by the semicolon delimeter.
// It ensure that there are only two sequences and returns the longest common subsequence.
func getSubsequences(sequences []string) ([]string, error) {
	subsequences := []string{}
	for _, val := range sequences {
		values := strings.Split(val, delimeter)
		if len(values) == 2 {
			subsequences = append(subsequences, getLongestCommonSubsequence(values[0], values[1]))
		}
	}
	return subsequences, nil
}

// getLongestCommonSubsequence returns the longest common subsequence for the sequence pair.
func getLongestCommonSubsequence(s1 string, s2 string) string {
	return readLongestCommonSubsequence(s1, s2, len(s1), len(s2), computeLengthTable(s1, s2))
}

// computeLengthTable sets the appropriate values in the length table.
func computeLengthTable(s1 string, s2 string) [][]int {
	grid := initGrid(s1, s2)
	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				grid[x][y] = grid[x-1][y-1] + 1
			} else {
				grid[x][y] = max(grid[x][y-1], grid[x-1][y])
			}
		}
	}
	return grid
}

// initGrid takes in two string values and initializes the grid that will be used to compute the LCS.
// It returns the initialized grid that will be used to compute the LCS.
func initGrid(s1 string, s2 string) [][]int {
	grid := make([][]int, 1+len(s1))
	for i := 0; i < len(grid); i++ {
		grid[i] = make([]int, 1+len(s2))
	}
	return grid
}

// readLongestCommonSubsequence calculates the longest common subsequence
func readLongestCommonSubsequence(s1 string, s2 string, x int, y int, lengthTable [][]int) string {
	if x == 0 || y == 0 {
		return ""
	} else if s1[x-1] == s2[y-1] {
		return readLongestCommonSubsequence(s1, s2, x-1, y-1, lengthTable) + string(s1[x-1])
	} else {
		if lengthTable[x][y-1] > lengthTable[x-1][y] {
			return readLongestCommonSubsequence(s1, s2, x, y-1, lengthTable)
		}
		return readLongestCommonSubsequence(s1, s2, x-1, y, lengthTable)
	}
}

// max computes the maximum of two integer values.
func max(val1 int, val2 int) int {
	if val1 > val2 {
		return val1
	}
	return val2
}
