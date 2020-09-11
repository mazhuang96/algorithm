/*
 * @Author: mazhuang
 * @Date: 2020-09-10 10:51:16
 * @LastEditTime: 2020-09-10 11:14:05
 * @Description:
 */
package main

import (
	"alofithm/datastruct/queue"
	"alofithm/sort"
	"fmt"
)

func main() {
	// testQueue()
	testSelectionSort()

}

func testQueue() {
	q := queue.New()
	q.Enqueue(1)
	q.Enqueue("2")
	q.Enqueue(3.4)
	n := q.Dequeue()

	fmt.Print(n)
	q.Enqueue(&n)
	fmt.Print(q)
}

func testSlice() {
	a := [][]int{}
	b := []int{}
	b = append(b, 1, 3, 5)
	fmt.Printf("a: %v, b:%v, &a: %p, &b: %p\n", a, b, a, b)
	a = append(a, b)
	fmt.Printf("a: %v, b:%v, &a: %p, &b: %p\n", a, b, a, b)
	b = append(b, 2)
	fmt.Printf("a: %v, b:%v, &a: %p, &b: %p\n", a, b, a, b)
	b = b[:len(b)-1]
	b = b[:len(b)-1]
	fmt.Printf("a: %v, b:%v, &a: %p, &b: %p\n", a, b, a[0], b)
	b = append(b, 6)
	fmt.Printf("a: %v, b:%v", a, b)
}

func testSelectionSort() {
	a := []int{3, 5, 1, 2, 6, 4, 0, 9, 7, 8, 1, 0}
	a = sort.Selection(a)
	fmt.Printf("sort a: %v", a)
}
