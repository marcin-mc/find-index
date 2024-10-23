package service

import (
	"bufio"
	"os"
	"strconv"
)

// FindIndex finds index of target value in a slice, returning the index and value
// If there is no target value, it returns index of a value within 10% range of target value.
// If no values can be found, it returns -1 for index and value.
func FindIndex(sli []int, target int) (int, int) {
	if len(sli) == 0 {
		return -1, -1
	}
	l, r := 0, len(sli)-1
	tolerance := target / 10
	approxMaxIdx := -1
	var mid int
	for l <= r {
		mid = (r + l) / 2
		if sli[mid] == target {
			return mid, sli[mid]
		}
		if Abs(sli[mid]-target) <= tolerance {
			approxMaxIdx = mid
		}
		if target < sli[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if approxMaxIdx > -1 {
		return approxMaxIdx, sli[approxMaxIdx]
	}
	return -1, -1
}

// Abs gets absolute value of int.
func Abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}

// LoadNumbers reads a file from given path and returns a slice of numbers.
func LoadNumbers(filepath string) ([]int, error) {
	var data []int
	file, err := os.Open(filepath)
	if err != nil {
		return data, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return data, err
		}
		data = append(data, value)
	}
	return data, nil
}
