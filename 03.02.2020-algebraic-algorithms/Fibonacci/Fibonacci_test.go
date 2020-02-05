package Fibonacci

import (
	"testing"
)

const (
	a1, a3, a5, a10 = 5, 100, 10000, 1000000000
	b, c, d         = 3, 13, 35
)

func Benchmark_CountFibonacciIteration1(b *testing.B) {
	_ = CountFibonacciIteration(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciIteration(a1)
	}
}

func Benchmark_CountFibonacciGold1(b *testing.B) {
	_ = CountFibonacciGold(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciGold(a1)
	}
}

func Benchmark_CountFibonacciMatrix1(b *testing.B) {
	_ = CountFibonacciMatrix(a1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciMatrix(a1)
	}
}

func Benchmark_CountFibonacciIteration3(b *testing.B) {
	_ = CountFibonacciIteration(a3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciIteration(a3)
	}
}

func Benchmark_CountFibonacciGold3(b *testing.B) {
	_ = CountFibonacciGold(a3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciGold(a3)
	}
}

func Benchmark_CountFibonacciMatrix3(b *testing.B) {
	_ = CountFibonacciMatrix(a3)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciMatrix(a3)
	}
}

func Benchmark_CountFibonacciIteration5(b *testing.B) {
	_ = CountFibonacciIteration(a5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciIteration(a5)
	}
}

func Benchmark_CountFibonacciGold5(b *testing.B) {
	_ = CountFibonacciGold(a5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciGold(a5)
	}
}

func Benchmark_CountFibonacciMatrix5(b *testing.B) {
	_ = CountFibonacciMatrix(a5)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciMatrix(a5)
	}
}

func Benchmark_CountFibonacciIteration10(b *testing.B) {
	_ = CountFibonacciIteration(a10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciIteration(a10)
	}
}

func Benchmark_CountFibonacciGold10(b *testing.B) {
	_ = CountFibonacciGold(a10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciGold(a10)
	}
}

func Benchmark_CountFibonacciMatrix10(b *testing.B) {
	_ = CountFibonacciMatrix(a10)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = CountFibonacciMatrix(a10)
	}
}

func TestCountFibonacciGold(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{a: a1},
			want: 5,
		},
		{
			name: "3",
			args: args{a: b},
			want: 2,
		},
		{
			name: "13",
			args: args{a: c},
			want: 233,
		},
		{
			name: "35",
			args: args{a: d},
			want: 9227465,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFibonacciGold(tt.args.a); got != tt.want {
				t.Errorf("CountFibonacciGold() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountFibonacciIteration(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{a: a1},
			want: 5,
		},
		{
			name: "3",
			args: args{a: b},
			want: 2,
		},
		{
			name: "13",
			args: args{a: c},
			want: 233,
		},
		{
			name: "35",
			args: args{a: d},
			want: 9227465,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFibonacciIteration(tt.args.a); got != tt.want {
				t.Errorf("CountFibonacciIteration() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCountFibonacciMatrix(t *testing.T) {
	type args struct {
		a int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "5",
			args: args{a: a1},
			want: 5,
		},
		{
			name: "3",
			args: args{a: b},
			want: 2,
		},
		{
			name: "13",
			args: args{a: c},
			want: 233,
		},
		{
			name: "35",
			args: args{a: d},
			want: 9227465,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountFibonacciMatrix(tt.args.a); got != tt.want {
				t.Errorf("CountFibonacciMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
