/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-11 16:05:37
 * @LastEditTime: 2020-09-14 14:30:55
 * @Description: 选择排序
 **************************************/
package sort

import "fmt"

func Selection(a []int) []int {
	for i := 0; i < len(a); i++ {
		//
		min := i
		for j := i + 1; j < len(a); j++ {
			// a[i] 和剩下的元素选择最小值
			if a[min] > a[j] {
				min = j
			}
		}
		// 最小值和a[i] 交换
		a[min], a[i] = a[i], a[min]
		fmt.Printf("%v times: %v\n", i, a)
	}
	return a
}
