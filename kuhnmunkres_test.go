package kuhnMunkres

import (
	"reflect"
	"testing"
)

func TestNewMunkres(t *testing.T) {
	type args struct {
		matrix [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{"1", args{[][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}}}, [][]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}}},
		{"2", args{[][]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {3, 6, 9, 12}, {4, 8, 12, 16}}}, [][]int{{0, 0, 0, 1}, {0, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 0}}},
		{"3", args{[][]int{{1, 2, 3, 4}, {2, 4, 6, 8}, {3, 6, 9, 12}}}, [][]int{{0, 0, 1, 0}, {0, 1, 0, 0}, {1, 0, 0, 0}}},
		{"4", args{[][]int{{1, 2, 3}, {2, 4, 6}, {3, 6, 9}, {4, 8, 12}}}, [][]int{{0, 0, 1}, {0, 1, 0}, {1, 0, 0}, {1, 0, 0}}},
		{"5", args{[][]int{{1, 2, 3}}}, [][]int{{1, 0, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewMunkres(tt.args.matrix).RunMunkres(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMunkres().RunMunkres() = %v, want %v", got, tt.want)
			}
		})
	}
}
