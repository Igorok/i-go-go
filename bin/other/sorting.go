package main

import (
	"fmt"
	"math"
)

//  := []int{54,26,1,93,100,17,15,77,57,31,44,55,11,20,94,94}

// sorting interface
type Sorting interface {
	sort() []int
}

// bubble sorting
type BubbleSorting struct {
	array []int
}

func (bs BubbleSorting) sort() []int {
	for i := len(bs.array) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if bs.array[j] > bs.array[j+1] {
				bs.array[j], bs.array[j+1] = bs.array[j+1], bs.array[j]
			}
		}
	}

	return bs.array
}

// insertion sorting
type InsertionSorting struct {
	array []int
}

func (ins InsertionSorting) sort() []int {
	for i := 1; i < len(ins.array); i++ {
		j := i
		tmp := ins.array[i]

		for j > 0 {
			if ins.array[j-1] > tmp {
				ins.array[j] = ins.array[j-1]
				j--
			} else {
				break
			}
		}

		if i != j {
			ins.array[j] = tmp
		}
	}

	return ins.array
}

// insertion sorting
type MergeSorting struct {
	array []int
}

func (ms MergeSorting) recSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	middle := int(math.Round(float64(len(arr) / 2)))
	left := ms.recSort(arr[0:middle])
	right := ms.recSort(arr[middle:])

	var mergedArr []int
	i := 0
	j := 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			mergedArr = append(mergedArr, left[i])
			i++
		} else {
			mergedArr = append(mergedArr, right[j])
			j++
		}
	}

	for i < len(left) {
		mergedArr = append(mergedArr, left[i])
		i++
	}

	for j < len(right) {
		mergedArr = append(mergedArr, right[j])
		j++
	}

	return mergedArr
}

func (ms MergeSorting) sort() []int {
	return ms.recSort(ms.array)
}

// quick sorting
type QuickSorting struct {
	array []int
	index int
}

func (qs QuickSorting) recSort(array []int) []int {
	if len(array) <= 1 {
		return array
	}

	var left []int
	var right []int

	for i := 0; i < len(array); i++ {
		if i == qs.index {
			continue
		}

		if array[i] >= array[qs.index] {
			right = append(right, array[i])
		} else {
			left = append(left, array[i])
		}
	}

	left = append(left, array[qs.index])

	left = qs.recSort(left)
	right = qs.recSort(right)

	return append(left[:], right[:]...)
}

func (qs QuickSorting) sort() []int {
	return qs.recSort(qs.array)
}

// factory for different sortings
func factory(method string, array []int) Sorting {
	if method == "insert" {
		return &InsertionSorting{array: array}
	} else if method == "merge" {
		return &MergeSorting{array: array}
	} else if method == "quick" {
		return &QuickSorting{array: array, index: 0}
	}
	return &BubbleSorting{array: array}
}

func main() {
	var sorting Sorting

	sorting = factory("bubble", []int{54, 26, 1, 93, 100, 17, 15, 77, 57, 31, 44, 55, 11, 20, 94, 94})
	fmt.Println("bubble", sorting.sort())

	sorting = factory("insert", []int{54, 26, 1, 93, 100, 17, 15, 77, 57, 31, 44, 55, 11, 20, 94, 94})
	fmt.Println("insert", sorting.sort())

	sorting = factory("merge", []int{54, 26, 1, 93, 100, 17, 15, 77, 57, 31, 44, 55, 11, 20, 94, 94})
	fmt.Println("merge", sorting.sort())

	sorting = factory("quick", []int{54, 26, 1, 93, 100, 17, 15, 77, 57, 31, 44, 55, 11, 20, 94, 94})
	fmt.Println("quick", sorting.sort())
}
