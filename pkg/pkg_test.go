package pkg

import "testing"

func TestFun1(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "Fun1 test", want: 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fun1(); got != tt.want {
				t.Errorf("Fun1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFun2(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "Fun2 test", want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fun2(); got != tt.want {
				t.Errorf("Fun2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFun3(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "Fun3 test", want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fun3(); got != tt.want {
				t.Errorf("Fun3() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFun4(t *testing.T) {
	tests := []struct {
		name string
		want int64
	}{
		{name: "Fun4 test", want: 6},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Fun4(); got != tt.want {
				t.Errorf("Fun4() = %v, want %v", got, tt.want)
			}
		})
	}
}
