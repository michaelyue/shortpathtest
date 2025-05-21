package main

import (
	"reflect"
	"testing"
)

func Test_shortestPath(t *testing.T) {
	type args struct {
		grid [][]int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Invalid input",
			args: args{
				grid: [][]int{},
			},
			want: [][]int{},
		},
		{
			name: "Case 1",
			args: args{
				grid: [][]int{{0}},
			},
			want: [][]int{{0}},
		},
		{
			name: "Case 2",
			args: args{
				grid: [][]int{{0, -1}},
			},
			want: [][]int{{0, 1}},
		},
		{
			name: "Case 3",
			args: args{
				grid: [][]int{{0, -1, -1}},
			},
			want: [][]int{{0, 1, 2}},
		},
		{
			name: "Case 4",
			args: args{
				grid: [][]int{
					{0, -1, -1},
					{0, -1, -1}},
			},
			want: [][]int{
				{0, 1, 2},
				{0, 1, 2}},
		},
		{
			name: "Case 5",
			args: args{
				grid: [][]int{
					{0, -1, -1},
					{0, -1, 0}},
			},
			want: [][]int{
				{0, 1, 1},
				{0, 1, 0}},
		},
		{
			name: "Case 6",
			args: args{
				grid: [][]int{
					{0, -1, -1},
					{0, -1, -1},
					{0, -1, 0}},
			},
			want: [][]int{
				{0, 1, 2},
				{0, 1, 1},
				{0, 1, 0}},
		},
		{
			name: "Case 7",
			args: args{
				grid: [][]int{
					{-1, -1, -1},
					{-1, -1, -1}},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0}},
		},
		{
			name: "Case 8",
			args: args{
				grid: [][]int{
					{0, 0, 0},
					{0, 0, 0}},
			},
			want: [][]int{
				{0, 0, 0},
				{0, 0, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := shortestPath(tt.args.grid); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestPath() = %v, want: %v", got, tt.want)
			}
		})
	}
}
