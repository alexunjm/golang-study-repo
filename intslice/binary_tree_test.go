package intslice

import (
	"reflect"
	"testing"
)

func TestSort2(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "example",
			args: args{
				slice: []int{4, 5, 2, 1, 2, 3},
			},
			want: []int{1, 2, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sort2(tt.args.slice); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Sort2() = %v, want %v", got, tt.want)
			}
		})
	}
}
