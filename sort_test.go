package sortalg

import (
	"reflect"
	"slices"
	"testing"
)

func TestSortingAlgorithms(t *testing.T) {
	tests := []struct {
		name string
		fn   func([]int)
	}{
		{"BubbleSort", BubbleSort},
		{"InsertionSort", InsertionSort},
		{"SelectionSort", SelectionSort},
		{"HeapSort", HeapSort},
		{"QuickSort", QuickSort},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := []int{42, 34, 25, 12, 22, 11, 112, 0, -5, 3}
			want := []int{-5, 0, 3, 11, 12, 22, 25, 34, 42, 112}

			if tt.name == "MergeSort" {
				got := MergeSort(slices.Clone(input))
				if !reflect.DeepEqual(got, want) {
					t.Errorf("MergeSort() = %v, want %v", got, want)
				}
				return
			}

			got := slices.Clone(input)
			tt.fn(got)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("%s() = %v, want %v", tt.name, got, want)
			}
		})
	}
}

func TestMergeSort(t *testing.T) {
	input := []int{64, 34, 25, 12, 22, 11, 90, 0, -5, 3}
	want := []int{-5, 0, 3, 11, 12, 22, 25, 34, 64, 90}
	got := MergeSort(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("MergeSort() = %v, want %v", got, want)
	}
}
