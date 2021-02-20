package ex1

import (
	"testing"
)

func Test_ex1(t *testing.T) {
	type args struct {
		n int
		i int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "test",
			args: args{8, 0},
			want: 6,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ex1(tt.args.n, tt.args.i); got != tt.want {
				t.Errorf("ex1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_test(t *testing.T) {
	tests := []struct {
		name string
		want int
	}{
		{"prueba", -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := test(); got != tt.want {
				t.Errorf("test() = %v, want %v", got, tt.want)
			}
		})
	}
}
