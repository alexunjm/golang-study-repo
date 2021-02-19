package main

import (
	"testing"
)

func Test_intSlice_elementsEqual(t *testing.T) {
	type args struct {
		to intSlice
	}
	type expected struct {
		value bool
	}
	tests := []struct {
		name     string
		s        intSlice
		args     args
		expected expected
	}{
		{
			name: "intSliceEqual returns true when compare empty slices",
			s:    intSlice{},
			args: args{
				to: intSlice{},
			},
			expected: expected{
				value: true,
			},
		},
		{
			name: "intSliceEqual returs false when two arrays are different len",
			s:    intSlice{1, 2, 4, 8},
			args: args{
				to: intSlice{1, 2, 4, 8, 16},
			},
			expected: expected{
				value: false,
			},
		},
		{
			name: "intSliceEqual returns false when elements are differents",
			s:    intSlice{1, 2, 6, 8, 16},
			args: args{
				to: intSlice{1, 2, 4, 8, 16},
			},
			expected: expected{
				value: false,
			},
		},
		{
			name: "intSliceEqual 1",
			s:    intSlice{1, 2, 4, 8, 16},
			args: args{
				to: intSlice{1, 2, 4, 8, 16},
			},
			expected: expected{
				value: true,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.s.elementsEqual(tt.args.to); got != tt.expected.value {
				t.Logf("\n\n============================\n UNEXPECTED %v", tt.name)
				t.Errorf("\n\n%d.intSliceEqual(%d) = %v,\n expected %v",
					tt.s,
					tt.args.to,
					got,
					tt.expected.value,
				)
				t.Log("\n============================\n\n")
			}
		})
	}
}
