package greatestCommonDivisor

import "testing"

const (
	a1, b1, c1 = 1725966, 49665, 33
	a2, b2, c2 = 84, 40, 4
	big1, big2 = 172596635665465465, 49665646546545
)

func TestGCDEuclidBinBit(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: a1,
				b: b1,
			},
			want: c1,
		},
		{
			args: args{
				a: a2,
				b: b2,
			},
			want: c2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCDEuclidBinBit(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCDEuclidBinBit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCDEuclidDiv(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: a1,
				b: b1,
			},
			want: c1,
		},
		{
			args: args{
				a: a2,
				b: b2,
			},
			want: c2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCDEuclidDiv(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCDEuclidDiv() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGCDEuclidSub(t *testing.T) {
	type args struct {
		a int
		b int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			args: args{
				a: a1,
				b: b1,
			},
			want: c1,
		},
		{
			args: args{
				a: a2,
				b: b2,
			},
			want: c2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GCDEuclidSub(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("GCDEuclidSub() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Benchmark_GCDEuclidSub(b *testing.B) {
	_ = GCDEuclidSub(a2, b2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidSub(a2, b2)
	}
}
func Benchmark_GCDEuclidDiv(b *testing.B) {
	_ = GCDEuclidDiv(a2, b2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidDiv(a2, b2)
	}
}
func Benchmark_GCDEuclidBinBit(b *testing.B) {
	_ = GCDEuclidBinBit(a2, b2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidBinBit(a2, b2)
	}
}

func Benchmark_GCDEuclidSub2(b *testing.B) {
	_ = GCDEuclidSub(a1, b1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidSub(a1, b1)
	}
}
func Benchmark_GCDEuclidDiv2(b *testing.B) {
	_ = GCDEuclidDiv(a1, b1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidDiv(a1, b1)
	}
}
func Benchmark_GCDEuclidBinBit2(b *testing.B) {
	_ = GCDEuclidBinBit(a1, b1)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidBinBit(a1, b1)
	}
}

func Benchmark_GCDEuclidSub3(b *testing.B) {
	_ = GCDEuclidSub(big1, big2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidSub(big1, big2)
	}
}
func Benchmark_GCDEuclidDiv3(b *testing.B) {
	_ = GCDEuclidDiv(big1, big2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidDiv(big1, big2)
	}
}
func Benchmark_GCDEuclidBinBit3(b *testing.B) {
	_ = GCDEuclidBinBit(big1, big2)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = GCDEuclidBinBit(big1, big2)
	}
}
