package main

func index(list [] int, predicate func(int) bool) int {
	for indexOfSlice, elementOfSlice := range list {
		if predicate(elementOfSlice) {
			return indexOfSlice
		}
	}
	return -1
}

func all(list []int, predicate func(int) bool) bool {
	return index(list, func(i int) bool {
		return !predicate(i)
	}) == -1
}

func any(list []int, predicate func(int) bool) bool {
	return index(list, predicate) != -1
}

func none(list []int, predicate func(int) bool) bool {
	return index(list, predicate) == -1
}

func indexElement(list []int, predicate func(int) bool) int {
	for i, element := range list {
		if predicate(element) {
			return i
		}
	}
	return -1
}

func findElement(list []int, predicate func(int) bool) int {
	return list[index(list, predicate)]
}


func main() {

}

