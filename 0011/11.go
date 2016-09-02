package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type matrix [20][20]int

const matrixSeed string = `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48`

// Largest product in a grid
// https://projecteuler.net/problem=11
func main() {
	points, err := createPoints(matrixSeed)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(checkProduct(points, 4))
}

// createPoints takes a string that has to be seperated by spaces for columns
// and \n for rows and returns a [20][20]int{} of type matrix
func createPoints(str string) (matrix, error) {
	m := [20][20]int{}

	// Split string by rows \n
	rows := strings.Split(str, "\n")
	for i, r := range rows {
		slice := strings.Split(r, " ")
		for j, s := range slice {
			p, err := strconv.Atoi(s)
			if err != nil {
				return [20][20]int{}, errors.New("str contains a value that's not an int")
			}
			if i >= 20 || j >= 20 {
				return [20][20]int{}, errors.New("str doesn't confirm to matrix 2d array")
			}

			m[i][j] = p
		}
	}

	return m, nil
}

// checkProduct loops through each row and col to get the product n-set of ints horizontally,
// vertically and diagonally (ascending, descending)
func checkProduct(m matrix, n int) int {
	maxRow := len(m)
	maxCol := len(m[0]) // all rows have the same number of columns
	var result int

	// Largest product of each row & column (horizontal & vertical)
	result = 1
	for i, r := range m {
		value := 1
		cValue := 1 // vertical product
		rValue := 1 // horizontal product
		dValue := 1 // decending diagonal product
		aValue := 1 // ascending diagonal product

		// Loop through each column value in row
		for j := range r {
			// We have to check point value then offset by one and check the next set
			for k := 0; k < n; k++ {
				// Product n adjacent row values (horizontal)
				cOffset := j + k
				if j < maxCol-(n-1) {
					p := m[i][cOffset]
					rValue *= p
				}

				rOffset := i + k
				if i < maxRow-(n-1) {
					p := m[rOffset][j] // vertical check
					cValue *= p
				}

				// Check descending diagonal
				if i < maxRow-(n-1) && j < maxCol-(n-1) {
					p := m[rOffset][cOffset]
					dValue *= p
				}

				// Check ascending diagonal
				rOffset = i + (n - 1) - k
				if i < maxRow-(n-1) && j < maxCol-(n-1) {
					p := m[rOffset][cOffset]
					aValue *= p
				}
			}

			// Check to see which value is the highest then reset them
			if rValue > value {
				value = rValue
			}
			rValue = 1

			if cValue > value {
				value = cValue
			}
			cValue = 1

			if dValue > value {
				value = dValue
			}
			dValue = 1

			if aValue > value {
				value = aValue
			}
			aValue = 1

		}

		// Set the largest product
		if value > result {
			result = value
		}
	}
	return result
}

// checkSum loops through each row and col to sum n-set of ints horizontally,
// vertically and diagonally (ascending, descending)
func checkSum(m matrix, n int) int {
	maxRow := len(m)
	maxCol := len(m[0]) // all rows have the same number of columns
	var result int

	// Largest sum of each row & column (horizontal & vertical)
	result = 0
	for i, r := range m {
		sum := 0
		cSum := 0 // vertical sum
		rSum := 0 // horizontal sum
		dSum := 0 // decending diagonal sum
		aSum := 0 // ascending diagonal sum

		// Loop through each column value in row
		for j := range r {
			// We have to check point value then offset by one and check the next set
			for k := 0; k < n; k++ {
				// Sum n adjacent row values (horizontal)
				cOffset := j + k
				if j < maxCol-(n-1) {
					p := m[i][cOffset]
					rSum += p
				}

				rOffset := i + k
				if i < maxRow-(n-1) {
					p := m[rOffset][j] // vertical check
					cSum += p
				}

				// Check descending diagonal
				if i < maxRow-(n-1) && j < maxCol-(n-1) {
					p := m[rOffset][cOffset]
					dSum += p
				}

				// Check ascending diagonal
				rOffset = i + (n - 1) - k
				if i < maxRow-(n-1) && j < maxCol-(n-1) {
					p := m[rOffset][cOffset]
					aSum += p
				}
			}

			// Check to see which sum is the highest then reset them
			if rSum > sum {
				sum = rSum
			}
			rSum = 0

			if cSum > sum {
				sum = cSum
			}
			cSum = 0

			if dSum > sum {
				sum = dSum
			}
			dSum = 0

			if aSum > sum {
				sum = aSum
			}
			aSum = 0

		}

		// Set the largest sum
		if sum > result {
			result = sum
		}
	}
	return result
}
