/*
A zero-indexed array A consisting of N integers is given. Rotation of the array means that each element is
shifted right by one index, and the last element of the array is also moved to the first place.
For example, the rotation of array A = [3, 8, 9, 7, 6] is [6, 3, 8, 9, 7].
The goal is to rotate array A K times; that is, each element of A will be shifted to the right by K indexes.
Write a function:
func Solution(A []int, K int) []int
that, given a zero-indexed array A consisting of N integers and an integer K, returns the array A rotated K times.
For example, given array A = [3, 8, 9, 7, 6] and K = 3, the function should return [9, 7, 6, 3, 8].
Assume that:
N and K are integers within the range [0..100];
each element of array A is an integer within the range [−1,000..1,000].
In your solution, focus on correctness. The performance of your solution will not be the focus of the assessment.
*/

/* return min number from the argument passed in the function */

package qnaarray

import "fmt"

func IntMin(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func CyclicRotate(A []int, K int) []int {
	if K == 0 || len(A) == 0 || len(A) == 1 {
		return A
	}
	length := len(A)
	if K > length {
		K = K % length
	}
	// A[lower:upper] //lower included/ upper not included
	// lhs = A[5-3:5]

	lhs := A[length-K : length]
	fmt.Println("Value of lhs = ", lhs)
	return append(lhs, A[0:length-K]...)

	// length := len(A)
	// if length > 0 {
	// 	if K > length {
	// 		K = K % length
	// 	}
	// 	A = append(A[length-K:length], A[0:length-K]...)
	// }
	// return A
}
