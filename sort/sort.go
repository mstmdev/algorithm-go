package sort

// SelectionSort 选择排序
//
// 算法描述：遍历数组，筛选出最小值放到数组的第一个位置，然后再次遍历剩下的数组，筛选出最小值放到第二个位置，依次类推，直到剩下最后一个元素
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
//
// 算法描述：将第一个元素当作一个有序数列，其他的元素为无序数列，依次从无序数列中取出元素，
// 按照正确的排序插入到有序数列中（从有序数列的末尾开始比对，如果小于有序数列的元素，则交换两者的位置，
// 从后往前进行对比，直到大于前面的元素或者来到第一个位置），直到无序数列中没有元素为止
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
//
// 算法描述：依次比较相邻两个元素的值，如果前者大于后者，则交换两者的位置，一次比较结束后最大值会在最后一个元素，
// 即大数往后冒泡，排除已经排序好的最大数，对剩下的数列再次进行比较冒泡，直至全部元素排序完毕
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
//
// 算法描述：将数组通过不断的对半切分，使其变为多个长度小于等于2的小数组，再对每个小型数组分别进行排序，最后将排好序的多个小数组合并为一个有序的大数组
func MergeSort(nums []int) []int {
	length := len(nums)
	if length <= 1 {
		return nums
	}
	numsL, numsR := divide(nums)
	// 如果切片元素大于4个，那么对切分后的切片进行递归处理
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

// min 返回两个数中的最小值，并告知最小值是否为第二个传入参数y
func min(x, y int) (minValue int, right bool) {
	if x > y {
		return y, true
	}
	return x, false
}

// generateNums 返回无序的测试数列以供测试
func generateNums() []int {
	return []int{10, 2, 54, 44, 61, -23, 123, 45, -94, 43, 61, 0, 34, 247, 113}
}
