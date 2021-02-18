package main

import (
	"reflect"
	"testing"
)

/*
func Test_sortr(t *testing.T) {
	type args struct {
		s []int
	}
	type expected struct {
		value []int
	}
	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			"sort empty slice",
			args{s: []int{}},
			expected{value: []int{}},
		},
		{
			"sort one element slice",
			args{s: []int{1}},
			expected{value: []int{1}},
		},
		{
			"sort two sorted elements slice",
			args{s: []int{1, 2}},
			expected{value: []int{1, 2}},
		},
		{
			"sort two elements slice",
			args{s: []int{3, 2}},
			expected{value: []int{2, 3, 4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := sort(tt.args.s)
			if !reflect.DeepEqual(actual, tt.expected.value) {
				t.Log("****************************************")
				t.Errorf("\n\nsort(%d):\n expected %d,\n actual %d\n",
					tt.args.s,
					tt.expected.value,
					actual)
				t.Log("****************************************")
			}
		})
	}
} */

type builder []int

func (b builder) make() []interface{} {
	interfaceSlice := make([]interface{}, len(b))
	for i, d := range b {
		interfaceSlice[i] = d
	}
	return interfaceSlice
}

func Test_reverse(t *testing.T) {
	type args struct {
		s []interface{}
	}
	tests := []struct {
		name string
		args args
		want []interface{}
	}{
		{
			name: "sort empty slice",
			args: args{make([]interface{}, 0)},
			want: make([]interface{}, 0),
		},
		{
			name: "sort one element slice",
			args: args{s: builder{1}.make()},
			want: builder{1}.make(),
		},
		{
			name: "sort two sorted elements slice",
			args: args{s: builder{1, 2}.make()},
			want: builder{1, 2}.make(),
		},
		{
			name: "sort two elements slice",
			args: args{s: builder{3, 2}.make()},
			want: builder{3, 2}.make(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reverse(%d) = %v, want %v",
					tt.args.s, got, tt.want)
				// t.Log("****************************************")
				// t.Errorf("\n\nsort(%d):\n expected %d,\n actual %d\n",
				// 	tt.args.s,
				// 	tt.expected.value,
				// 	actual)
				// t.Log("****************************************")
			}
		})
	}
}

/*
func Test_sort(t *testing.T) {
	type args struct {
		sliceArg []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sort(tt.args.sliceArg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sort() = %v, want %v", got, tt.want)
				// t.Log("****************************************")
				// t.Errorf("\n\nsort(%d):\n expected %d,\n actual %d\n",
				// 	tt.args.s,
				// 	tt.expected.value,
				// 	actual)
				// t.Log("****************************************")
			}
		})
	}
}
*/
