package qnaarray

import (
	"reflect"
	"testing"
)

func TestIntMinBasic(t *testing.T) {
	ans := IntMin(6, -4)
	if ans != -4 {
		t.Errorf("IntMin(2, -2) = %d; want -2", ans)
	}
}

func TestIntMinWithTable(t *testing.T) {
	type args struct {
		a, b int
	}
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "test 1",
			args: args{
				a: 1,
				b: 5,
			},
			want: 1,
		},
		{
			name: "test 1",
			args: args{
				a: 10,
				b: 11,
			},
			want: 10,
		},
		{
			name: "test 1",
			args: args{
				a: -1,
				b: 5,
			},
			want: -1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			min := IntMin(tt.args.a, tt.args.b)
			// if min != tt.want {
			//     t.Errorf("Expected = %d; got = %d", tt.want, min)
			// }
			if !reflect.DeepEqual(min, tt.want) {
				t.Errorf("Expected = %d; got = %d", tt.want, min)
			}
		})
	}
}

func TestCyclicRotate(t *testing.T) {
	type args struct {
		a []int
		k int
	}
	var tests = []struct {
		name string
		args args
		want []int
	}{
		{
			name: "test1",
			args: args{
				a: []int{1, 2, 3},
				k: 1,
			},
			want: []int{3, 1, 2},
		},
		{
			name: "test2",
			args: args{
				a: []int{1, 2, 3},
				k: 2,
			},
			want: []int{2, 3, 1},
		},
		{
			name: "Test a",
			args: args{
				a: []int{3, 8, 9, 7, 6},
				k: 1,
			},
			want: []int{6, 3, 8, 9, 7},
		},
		{
			name: "Test b",
			args: args{
				a: []int{3, 8, 9, 7, 6},
				k: 3,
			},
			want: []int{9, 7, 6, 3, 8},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			testing := CyclicRotate(tt.args.a, tt.args.k)
			if !reflect.DeepEqual(testing, tt.want) {
				t.Errorf("Expected = %d: got = %d\n", tt.want, testing)
			}
		})
	}
}
