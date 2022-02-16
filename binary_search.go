/*
	Binary rearch in Go
*/
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 6, 10, 15, 15, 17, 18, 19, 30}
	find := binarySearch(slice, 10, 0, len(slice)-1)
	fmt.Printf("Found the element by recursive biany search at index %d \n", find)
	find = iterBinarySearch(slice, 10, 0, len(slice)-1)
	fmt.Printf("Found the element by iterative biany search at index %d \n", find)
}

// using recurtion
func binarySearch(array []int, target int, lowIndex int, highIndex int) int {
	if highIndex < lowIndex {
		return -1 // not possible cannot find it
	}
	mid := int((lowIndex + highIndex) / 2) // define middle index
	if array[mid] > target {
		return binarySearch(array, target, lowIndex, mid)
	} else if array[mid] < target {
		return binarySearch(array, target, mid+1, highIndex)
	} else {
		return mid
	}
}

// using iteration
func iterBinarySearch(array []int, target int, lowIndex int, highIndex int) int {
	startIndex := lowIndex
	endIndex := highIndex
	var mid int
	for startIndex <= endIndex {
		mid = int((startIndex + endIndex) / 2)
		if array[mid] > target {
			endIndex = mid
		} else if array[mid] < target {
			startIndex = mid + 1
		} else {
			return mid
		}
	}
	return -1
}