/**************************************
 * @Author: mazhuang
 * @Date: 2020-09-11 16:05:37
 * @LastEditTime: 2020-09-11 18:16:27
 * @Description:
 **************************************/
package sort

import (
	"reflect"
	"testing"
)

func TestSelection(t *testing.T) {
	type args struct {
		a []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
		struct {
			name string
			args args
			want []int
		}{
			name: "1",
			args: args{
				a: []int{3, 6, 1, 5, 8, 2, 4, 9, 0, 7},
			},
			want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Selection(tt.args.a); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Selection() = %v, want %v", got, tt.want)
			}
		})
	}
}
