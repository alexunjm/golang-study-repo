package main

import (
	"testing"
)

func Test_sortr(t *testing.T) {
	type args struct {
		slice intSlice
	}
	type expected struct {
		value intSlice
	}
	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "sort empty slice",
			args: args{
				slice: intSlice{},
			},
			expected: expected{
				value: intSlice{},
			},
		},
		{
			name: "sort one element slice",
			args: args{
				slice: intSlice{1},
			},
			expected: expected{
				value: intSlice{1},
			},
		},
		{
			name: "sort two sorted elements slice",
			args: args{
				slice: intSlice{1, 2},
			},
			expected: expected{
				value: intSlice{1, 2},
			},
		},
		{
			name: "sort two elements slice",
			args: args{
				slice: intSlice{4, 3, 2},
			},
			expected: expected{
				value: intSlice{2, 3, 4},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sort(tt.args.slice)
			if !got.elementsEqual(tt.expected.value) {
				t.Logf("\n\n============================\n UNEXPECTED %v", tt.name)
				t.Errorf("\n\nsort(%d):\n expected %d,\n got %d\n",
					tt.args.slice,
					tt.expected.value,
					got)
				t.Log("\n============================\n\n")

			}
		})
	}
}

func Test_reverse(t *testing.T) {
	type args struct {
		slice intSlice
	}
	type expected struct {
		value intSlice
	}
	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "reverse empty slice",
			args: args{
				slice: intSlice{},
			},
			expected: expected{
				value: intSlice{},
			},
		},
		{
			name: "reverse one element slice",
			args: args{
				slice: intSlice{1},
			},
			expected: expected{
				value: intSlice{1},
			},
		},
		{
			name: "reverse two elements slice",
			args: args{
				slice: intSlice{3, 2},
			},
			expected: expected{
				value: intSlice{2, 3},
			},
		},
		{
			name: "reverse (n > 2) elements slice",
			args: args{
				slice: intSlice{3, 2, 5, 7, 9},
			},
			expected: expected{
				value: intSlice{9, 7, 5, 2, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverse(tt.args.slice); !got.elementsEqual(tt.expected.value) {
				t.Logf("\n\n============================\n UNEXPECTED %v", tt.name)
				t.Errorf("\n\nreverse(%d) -> %d,\n expected %d\n",
					tt.args.slice, got, tt.expected.value)
				t.Log("\n============================\n\n")
			}
		})
	}
}

func Test_movement(t *testing.T) {
	type args struct {
		slice    intSlice
		found    int
		position int
	}
	type expected struct {
		value intSlice
	}
	tests := []struct {
		name     string
		args     args
		expected expected
	}{
		{
			name: "movement found == position do nothing",
			args: args{
				slice:    intSlice{1, 2},
				found:    0,
				position: 0,
			},
			expected: expected{
				value: intSlice{1, 2},
			},
		},
		{
			name: "movement [1, 4, 3, 2], found = 3, position = 1 -> [1, 2, 4, 3]",
			args: args{
				slice:    intSlice{1, 4, 3, 2},
				found:    3,
				position: 1,
			},
			expected: expected{
				value: intSlice{1, 2, 4, 3},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sliceArg := make(intSlice, len(tt.args.slice))
			copy(sliceArg, tt.args.slice)
			movement(sliceArg, tt.args.found, tt.args.position)
			if !sliceArg.elementsEqual(tt.expected.value) {
				t.Logf("\n\n============================\n UNEXPECTED %v", tt.name)
				t.Errorf("\n\nmovement(%d, %v, %v) -> %d,\n expected %d\n",
					tt.args.slice, tt.args.found, tt.args.position,
					sliceArg, tt.expected.value)
				t.Log("\n============================\n\n")
			}
		})
	}
}
