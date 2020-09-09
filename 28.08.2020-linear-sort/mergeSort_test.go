package main

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"testing"
)

func readNumbersFromFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return file, err
	}
	return file, nil
}
func TestMergeSort(t *testing.T) {
	type args struct {
		numbers []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "mergeSort",
			args: args{numbers: []int{1, 5, 3, 2, 8, 7, 3, 0, 13, 6}},
			want: []int{0, 1, 2, 3, 3, 5, 6, 7, 8, 13},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MergeSort(tt.args.numbers); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MergeSort() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkMergeSort(b *testing.B) {
	file, err := readNumbersFromFile("./generated_1000000.bin")
	if err != nil {
		b.Errorf("MergeSort() = %v", err)
	}
	for i := 0; i < b.N; i++ {
		b.StopTimer()

		b.StartTimer()
		fmt.Println(file)
	}
}
