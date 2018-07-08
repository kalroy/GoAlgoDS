package main

import ("fmt")

func main() {
	numberToSearch := 5
	numbers := []int { 1,4,5,6,8}
	if Search(numbers, numberToSearch) {
		fmt.Println("Available")
	}
	fmt.Println("Missing")
}

// Search searches an integer from an array of integers
func Search (numbers []int, numberToSearch int) bool {
	lastIndex := len(numbers)

	isAvailable := BinarySearch(numberToSearch, numbers, 0, lastIndex)
	return isAvailable
}

// BinarySearch implements binary search algorithm using Recursion
func BinarySearch(numberToSearch int, numbers []int, startIndex int, lastIndex int) bool {
	if lastIndex < startIndex {
		return false
	}
	middleIndex := (startIndex + lastIndex) / 2
	if numbers[middleIndex] == numberToSearch {
		return true
	} else if numbers[middleIndex] > numberToSearch {
		return BinarySearch(numberToSearch, numbers, startIndex, middleIndex - 1)
	}
	return BinarySearch(numberToSearch, numbers, middleIndex + 1, lastIndex)
}