package sort

import "testing"

func BenchmarkSelectionSort(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	nums := generateNums()
	for i := 0; i < b.N; i++ {
		SelectionSort(nums)
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	nums := generateNums()
	for i := 0; i < b.N; i++ {
		InsertionSort(nums)
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	nums := generateNums()
	for i := 0; i < b.N; i++ {
		BubbleSort(nums)
	}
}

func BenchmarkMergeSort(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()
	nums := generateNums()
	for i := 0; i < b.N; i++ {
		MergeSort(nums)
	}
}
