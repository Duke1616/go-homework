package main

type RankBubbleSort struct {
}

type RankQuickSort struct {
}

func (RankBubbleSort) GetRank(fatRate float64) (rank int) {
	rankArr := make([]float64, 0, len(items))
	for _, v := range items {
		rankArr = append(rankArr, v.FatRate)
	}

	orderedSize := 0
	for i := 0; i < len(rankArr)-1; i++ {
		for j := orderedSize; j < len(rankArr)-i-1; j++ {
			if rankArr[j] > rankArr[j+1] {
				rankArr[j], rankArr[j+1] = rankArr[j+1], rankArr[j]
				if orderedSize == j && orderedSize > 0 {
					orderedSize--
				}
			} else if orderedSize == j {
				orderedSize++
			}
		}
	}

	for i, v := range rankArr {
		if v >= fatRate {
			return i + 1
		}
	}
	return 0
}

func (r *RankQuickSort) GetRank(fatRate float64) (rank int) {
	rankArr := make([]float64, 0, len(items))
	for _, v := range items {
		rankArr = append(rankArr, v.FatRate)
	}

	r.QuickSort(rankArr, 0, len(rankArr)-1)

	for i, v := range rankArr {
		if v >= fatRate {
			return i + 1
		}
	}
	return 0
}

func (r *RankQuickSort) QuickSort(data []float64, start, end int) {
	if start < end {
		base := data[start]
		left := start
		right := end
		// 进入循环
		for left < right {
			// 由于左边的(第0个)被取走当成基准值，所以在左边就留下一个洞，用于存储比基准小的值
			// 所以先从右边找，找到第一个比基准值小的
			for left < right && data[right] >= base {
				right--
			}
			// 找到了比基准值小的，那就把这个值填入左边的洞，做索引要++
			if left < right {
				data[left] = data[right]
				left++
			}
			// 因为上面的操作让右边留了一个洞，开始从左边找比基准值大的
			for left < right && data[left] <= base {
				left++
			}
			// 找到比基准值大的再次把右边洞填上，并且在左边又留了一个洞
			if left < right {
				data[right] = data[left]
				right--
			}
		}

		// 把基准值填入到洞中，这就是本应该他所在的位置
		data[left] = base
		// 继续分两组递归排序，注意此时基准值已经不用参与排序了
		r.QuickSort(data, start, left-1)
		r.QuickSort(data, left+1, end)
	}
}
