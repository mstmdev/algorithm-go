package sort

import (
	"testing"
)

func TestSelectionSort(t *testing.T) {
	t.Logf("SelectionSort result:%v", SelectionSort(generateNums()))
}

func TestInsertionSort(t *testing.T) {
	t.Logf("InsertionSort result:%v", InsertionSort(generateNums()))
}

func TestBubbleSort(t *testing.T) {
	t.Logf("BubbleSort result:%v", BubbleSort(generateNums()))
}

func TestMergeSort(t *testing.T) {
	t.Logf("MergeSort result:%v", MergeSort(generateNums()))
}

func TestQuickSort(t *testing.T) {
	t.Logf("QuickSort result:%v", QuickSort(generateNums()))
}

func TestShellSort(t *testing.T) {
	t.Logf("ShellSort result:%v", ShellSort(generateNums()))
}

// generateNums 返回无序的测试数列以供测试
func generateNums() []int {
	return []int{10, 2, 54, 44, 61, -23, 123, 45, -94, 43, 61, 0, 34, 247, 113}
}
