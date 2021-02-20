package ex2

import (
	"math/rand"
	"reflect"
	"sort"
	"testing"
	"time"
)

func Test_mergeSort(t *testing.T) {
	type args struct {
		s []int
	}

	rand.Seed(time.Now().UnixNano())
	manyElementSlice := rand.Perm(99)

	sorted := make([]int, 99)
	copy(sorted, manyElementSlice)
	sort.Ints(sorted)

	tests := []struct {
		name string
		args args
		want []int
	}{
		{"prueba", args{[]int{5, 4, 3, 2, 1}}, []int{1, 2, 3, 4, 5}},
		{"prueba", args{manyElementSlice}, sorted},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mergeSort(tt.args.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}
