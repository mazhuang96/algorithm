/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-11 17:10:14
 * @LastEditTime: 2020-09-14 14:39:45
 * @Description: 插入排序
 **************************************/
package sort

import "fmt"

func insertion(a []int) []int {
	for i := 1; i < len(a); i++ {
		for j := i; j > 0 && (a[j] < a[j-1]); j-- {
			// fmt.Println(i, j, a[i], a[j-1])
			// 从i位置向前，一一比较相邻元素，并交换位置，i 之前的元素都是有序的
			a[j], a[j-1] = a[j-1], a[j]
			fmt.Println(i, a)
		}
	}
	return a
}
