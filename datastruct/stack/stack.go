/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-10 16:47:53
 * @LastEditTime: 2020-09-10 18:41:52
 * @Description:
 **************************************/
package stack

type stack struct {
	data []interface{}
}

func New() *stack {
	return &stack{
		data: []interface{}{},
	}
}

func (s *stack) Push(n interface{}) {
	s.data = append(s.data, n)
}

func (s *stack) Pop() (n interface{}) {
	n = s.data[len(s.data)]
	s.data = s.data[:len(s.data)-1]
	return
}

func (s *stack) Length() int {
	return len(s.data)
}
