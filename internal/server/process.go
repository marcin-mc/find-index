package server

import (
	"bufio"
	"math"
	"os"
	"strconv"
)

// FindIndex finds index of target value in a slice.
// If there is no target value, it returns index of a value within 10% range of target value.
func FindIndex(sli []int, target int) (int, int) {
	if len(sli) == 0 {
		return -1, -1
	}
	var mid int
	l, r := 0, len(sli)-1
	for l <= r {
		mid = (r + l) / 2
		if sli[mid] == target {
			return mid, sli[mid]
		} else if target < sli[mid] {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	if math.Abs(float64(sli[mid])-float64(target)) <= float64(target)/10 {
		return mid, sli[mid]
	}
	return -1, -1
}

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
