package main

import "fmt"

func reverse(arr []int) {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
}

// func rotate(arr int[], n int) []int {
// 	arr = arr[-n:] + arr[:-n]
// }

// func rotate(arr []int, n int) []int {
// 	return append(arr[n:], arr[:n]...)
// }

// use in-place rotation
func rotate(arr []int, n int) []int {
	reverse(arr[:n])
	reverse(arr[n:])
	reverse(arr)
	return arr
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	n := 3
	rotate(arr, n)
	arr = rotate(arr, n)
	fmt.Println(arr)
	fmt.Println(rotate(arr, n))
}
