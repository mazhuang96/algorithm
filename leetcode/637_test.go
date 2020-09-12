/*
 * @Author: mazhuang
 * @Date: 2020-09-12 10:34:35
 * @LastEditors: mazhuang
 * @LastEditTime: 2020-09-12 11:08:31
 * @FilePath: /algorithm/leetcode/637_test.go
 * @Description:
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_averageOfLevels(t *testing.T) {
	type args struct {
		root *TreeNode
	}
	tests := []struct {
		name         string
		args         args
		wantAverages []float64
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotAverages := averageOfLevels(tt.args.root); !reflect.DeepEqual(gotAverages, tt.wantAverages) {
				t.Errorf("averageOfLevels() = %v, want %v", gotAverages, tt.wantAverages)
			}
		})
	}
}
