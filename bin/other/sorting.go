package main

import (
	"fmt"
)

// [17 20 26 31 44 54 55 77 93 94 94]
/*
Bubble sort
Алгоритм состоит из повторяющихся проходов по сортируемому массиву. За каждый проход элементы последовательно сравниваются попарно и, если порядок в паре неверный, выполняется обмен элементов.
*/
func sortBubble(arr []int) []int {
	for i := len(arr) - 1; i > 0; i-- {
		change := false
		for j := 0; j < i; j++ {
			if arr[j] > arr[j+1] {
				change = true
				t := arr[j+1]
				arr[j+1] = arr[j]
				arr[j] = t
			}
		}
		if !change {
			break
		}
	}

	return arr
}

/*
Selection sort
Шаги алгоритма:
 - находим номер минимального значения в текущем списке
 - производим обмен этого значения со значением первой неотсортированной позиции (обмен не нужен, если минимальный элемент уже находится на данной позиции)
 - теперь сортируем хвост списка, исключив из рассмотрения уже отсортированные элементы
*/
func sortSelection(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		min := i
		for j := i + 1; j < len(arr); j++ {
			if arr[j] < arr[min] {
				min = j
			}
		}
		if min != i {
			t := arr[i]
			arr[i] = arr[min]
			arr[min] = t
		}
	}

	return arr
}

/*
Insertion sort
Алгоритм сортировки, в котором элементы входной последовательности просматриваются по одному, и каждый новый поступивший элемент размещается в подходящее место среди ранее упорядоченных элементов
*/
func sortInsertion(arr []int) []int {
	for i := 1; i < len(arr)-1; i++ {
		t := arr[i]
		j := i
		for j > 0 {
			if arr[j-1] > t {
				arr[j] = arr[j-1]
				j = j - 1
			} else {
				break
			}
		}
		arr[j] = t
	}

	return arr
}

func main() {
	arr := []int{54, 26, 93, 17, 77, 31, 44, 55, 20, 94, 94}

	// sorted := sortBubble(arr)
	// sorted := sortSelection(arr)
	sorted := sortInsertion(arr)

	fmt.Println("sorted", sorted)
}
