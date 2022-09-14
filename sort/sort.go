package sort

// SelectionSort 选择排序
func SelectionSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	// 最后只剩一个元素 无需比较
	for i := 0; i < length-1; i++ {
		minValue := nums[i]
		minIndex := i
		var right bool
		for j := i; j < length; j++ {
			minValue, right = min(minValue, nums[j])
			if right {
				minIndex = j
			}
		}
		nums[minIndex] = nums[i]
		nums[i] = minValue
	}
	return nums
}

// InsertionSort 插入排序
func InsertionSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	for i := 1; i < length; i++ {
		num := nums[i] // 未排序的下一个元素
		numIndex := i
		// 与前面的元素依次比较，如果小于前面的元素，则交换位置，直到大于前面的元素或者来到第一个位置
		for j := i - 1; j >= 0; j-- {
			prev := nums[j]
			if num >= prev {
				break
			} else {
				nums[j] = num
				nums[numIndex] = prev
				numIndex = j
			}
		}
	}
	return nums
}

// BubbleSort 冒泡排序
func BubbleSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	loopTimes := length - 1
	for i := 0; i < loopTimes; i++ {
		loopElements := loopTimes - i
		for j := 0; j < loopElements; j++ {
			if nums[j] > nums[j+1] {
				nums[j], nums[j+1] = nums[j+1], nums[j]
			}
		}
	}
	return nums
}

func min(x, y int) (minValue int, right bool) {
	if x > y {
		return y, true
	}
	return x, false
}

func generateNums() []int {
	return []int{10, 2, 54, 44, 61, -23, 123, 45, -94, 43, 61, 0, 34, 247, 113}
}
