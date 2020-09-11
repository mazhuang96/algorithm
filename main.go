/*
 * @Author: mazhuang
 * @Date: 2020-09-10 10:51:16
 * @LastEditTime: 2020-09-10 11:14:05
 * @Description:
 */
package main

import (
	"alofithm/datastruct/queue"
	"fmt"
)

func main() {
	testQueue()
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
