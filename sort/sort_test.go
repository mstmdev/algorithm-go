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
