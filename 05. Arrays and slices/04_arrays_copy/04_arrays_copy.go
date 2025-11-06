package main

import "fmt"

func main() {

	nums1 := [4]int{3, 4, 5, 6}
	nums2 := nums1 // копирование массива nums1 в nums2

	nums2[1] = 11 // меняем значение в nums2, nums1 не меняется

	fmt.Println("nums1:", nums1) // nums1: [3 4 5 6]

	fmt.Println("nums2:", nums2) // nums2: [3 11 5 6]
}
