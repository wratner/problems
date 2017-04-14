package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

//checkIfFailedToRead logs and fail gracefully if unable to read input
func checkIfFailedToRead(err error) {
	if err != nil {
		log.Fatalln("Failed to read input: ", err.Error)
	}
}

// readInput takes in an io.Reader (os.Stdin as the standard, strings.NewReader for testing)
// It returns the received value with spaces trimmed off.
func readInput(reader io.Reader) (string, error) {
	scan := bufio.NewReader(reader)
	value, err := scan.ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(value), nil
}

func main() {
	s1, err := readInput(os.Stdin)
	checkIfFailedToRead(err)
	s2, err := readInput(os.Stdin)
	checkIfFailedToRead(err)

	if s1 == "" || s2 == "" {
		fmt.Println("Please make sure you provide a value for both String1 and String2")
		os.Exit(1)
	}

	fmt.Println(getLongestCommonSubstring(s1, s2))

}

// getLongestCommonSubstring takes in two strings and computes the longest common substring among them.
// It returns the computed longest common substring.
func getLongestCommonSubstring(s1 string, s2 string) string {
	grid := initGrid(s1, s2)
	longest, xLongest := 0, 0
	for x := 1; x < 1+len(s1); x++ {
		for y := 1; y < 1+len(s2); y++ {
			if s1[x-1] == s2[y-1] {
				grid[x][y] = grid[x-1][y-1] + 1
				if grid[x][y] > longest {
					longest = grid[x][y]
					xLongest = x
				}
			}
		}
	}
	return s1[xLongest-longest : xLongest]
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
