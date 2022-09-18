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

// MergeSort 归并排序
func MergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	numsL, numsR := divide(nums)
	// 如何切片元素大于4个，那么对切分后的切片进行递归处理
	// 即保证通过divide切分之后的2个子切片的元素个数分别都不会超过2个
	if len(nums) > 4 {
		return combine(MergeSort(numsL), MergeSort(numsR))
	}
	conquer(numsL, numsR)
	return combine(numsL, numsR)
}

// divide 将传入的切片切分为2份
func divide(nums []int) (numsL []int, numsR []int) {
	length := len(nums)
	subLen := length / 2
	numsL = nums[:subLen]
	numsR = nums[subLen:]
	return numsL, numsR
}

// conquer 对两个切片分别进行顺序排序，要求每个切片元素不能超过2个
func conquer(numsL []int, numsR []int) {
	if len(numsL) == 2 && numsL[0] > numsL[1] {
		numsL[0], numsL[1] = numsL[1], numsL[0]
	}

	if len(numsR) == 2 && numsR[0] > numsR[1] {
		numsR[0], numsR[1] = numsR[1], numsR[0]
	}
}

// combine 将两个切片按照顺序排序组合成一个新切片返回
func combine(numsL []int, numsR []int) []int {
	x, y, z := 0, 0, 0
	length := len(numsL) + len(numsR)
	sortNums := make([]int, length)
	for {
		if z == length {
			break
		}
		// 左侧切片是否已经被处理完毕
		if x >= len(numsL) || numsL[x] > numsR[y] {
			sortNums[z] = numsR[y]
			y++
		} else {
			sortNums[z] = numsL[x]
			x++
		}
		z++
	}
	return sortNums
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
