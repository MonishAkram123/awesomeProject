package main

import "testing"

func TestAdd(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	var tests = []struct {
		name string
		args args
		want int64
	}{
		{"Add 2 and 3", args{2, 3}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubtract(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	var tests = []struct {
		name string
		args args
		want int64
	}{
		{"Subtract 2 from 3", args{3, 2}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMultiply(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	var tests = []struct {
		name string
		args args
		want int64
	}{
		{"Multiply 2 and 3", args{2, 3}, 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Multiply(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDivide(t *testing.T) {
	type args struct {
		a int64
		b int64
	}
	var tests = []struct {
		name string
		args args
		want int64
	}{
		{"Dividing 6 by 3", args{6, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Divide(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
