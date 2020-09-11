/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-11 10:57:56
 * @LastEditTime: 2020-09-11 17:05:00
 * @Description:
 **************************************/
package leetcode

import (
	"reflect"
	"testing"
)

func Test_combinationSum3(t *testing.T) {
	type args struct {
		k int
		n int
	}

	type test struct {
		name string
		args args
		want [][]int
	}

	tests := []test{
		// TODO: Add test cases.
		test{
			name: "first",
			args: args{
				k: 3,
				n: 9,
			},
			want: [][]int{[]int{1, 2, 6}, []int{1, 3, 5}, []int{2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := combinationSum3(tt.args.k, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("combinationSum3() = %v, want %v", got, tt.want)
			}
		})
	}
}
