package main

import (
	"testing"
)

func Test_generateRandomElements(t *testing.T) {
	type args struct {
		size int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "zero",
			args: args{size: 0},
			want: nil,
		},
		{
			name: "negative",
			args: args{size: -1},
			want: nil,
		},
		{
			name: "size 1",
			args: args{size: 1},
			want: make([]int, 1),
		},
		{
			name: "size 10000",
			args: args{size: 10000},
			want: make([]int, 10000),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := generateRandomElements(tt.args.size)
			// Nill проверка
			if tt.want == nil {
				if got != nil {
					t.Errorf("got %v, want nil", got)
				}
				return
			}

			// Проверка длины
			if len(got) != len(tt.want) {
				t.Errorf("got length %d, want %d", len(got), len(tt.want))
			}
		})
	}
}
func Test_maximum(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{data: []int{}},
			want: 0,
		},
		{
			name: "one",
			args: args{data: []int{1}},
			want: 1,
		},
		{
			name: "same",
			args: args{data: []int{5, 5, 5, 5, 5}},
			want: 5,
		},
		{
			name: "normal",
			args: args{data: []int{2, 7, 5, 1, 9, 6, 4, 3, 8}},
			want: 9,
		},
		{
			name: "minus",
			args: args{data: []int{-2, -7, -5, -1, -9, -6, -4, -3, -8}},
			want: -1,
		},
		{
			name: "mixed",
			args: args{data: []int{2, -7, 5, -1, 9, -6, 4, -3, 8, -10, 10}},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maximum(tt.args.data); got != tt.want {
				t.Errorf("maximum() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxChunks(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "zero",
			args: args{data: []int{}},
			want: 0,
		},
		{
			name: "one",
			args: args{data: []int{1}},
			want: 1,
		},
		{
			name: "same",
			args: args{data: []int{5, 5, 5, 5, 5}},
			want: 5,
		},
		{
			name: "normal",
			args: args{data: []int{2, 7, 5, 1, 9, 6, 4, 3, 8}},
			want: 9,
		},
		{
			name: "minus",
			args: args{data: []int{-2, -7, -5, -1, -9, -6, -4, -3, -8}},
			want: -1,
		},
		{
			name: "mixed",
			args: args{data: []int{2, -7, 5, -1, 9, -6, 4, -3, 8, -10, 10}},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxChunks(tt.args.data); got != tt.want {
				t.Errorf("maxChunks() = %v, want %v", got, tt.want)
			}
		})
	}
}
