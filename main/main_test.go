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
	}{{"Success", args{1, 2}, 3}}
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
	}{{"Success", args{1, 2}, -1}}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Subtract(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}
