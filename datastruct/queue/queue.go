/*
 * @Author: mazhuang
 * @Date: 2020-09-10 10:09:45
 * @LastEditTime: 2020-09-10 11:47:56
 * @Description:
 */
package queue

import "container/list"

type queue struct {
	data   []interface{}
	length int
}

func New() *queue {
	return &queue{
		data:   []interface{}{},
		length: 0,
	}
}

func (q *queue) Enqueue(n interface{}) {
	q.data = append(q.data, n)
	q.length += 1
}

func (q *queue) Dequeue() (n interface{}) {
	if q.length == 0 {
		panic("null queue")
	}
	n = q.data[0]
	q.data = q.data[1:]
	q.length -= 1
	return n
}

func (q *queue) Length() (len int) {
	list.New().Back()
	return q.length
}
