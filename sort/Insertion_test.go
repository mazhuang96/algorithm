/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-11 17:10:14
 * @LastEditTime: 2020-09-14 10:15:10
 * @Description: 插入排序
 **************************************/
package sort

import (
	"reflect"
	"testing"
)

func TestInsertion(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		{
			name: "1",
			args: args{
				a: []int{3, 6, 1, 5, 8, 2, 4, 9, 0, 7},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertion(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertion() = %v, want %v", got, tt.want)
			}
		})
	}
}
