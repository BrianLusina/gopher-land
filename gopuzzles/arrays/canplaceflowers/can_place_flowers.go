package canplaceflowers

func canPlaceFlowers(flowerbed []int, n int) bool {
	count := 0

	for idx, flowerPlot := range flowerbed {
		if flowerPlot == 0 {
			emptyLeftPlot := idx == 0 || flowerbed[idx-1] == 0
			emptyRightPlot := idx == len(flowerbed)-1 || flowerbed[idx+1] == 0

			if emptyLeftPlot && emptyRightPlot {
				flowerbed[idx] = 1
				count++
			}
		}

	}

	return count >= n
}

func canPlaceFlowers2(flowerbed []int, n int) bool {
	count := 0

	for idx, flowerPlot := range flowerbed {
		if flowerPlot == 0 {
			emptyLeftPlot := idx == 0 || flowerbed[idx-1] == 0
			emptyRightPlot := idx == len(flowerbed)-1 || flowerbed[idx+1] == 0

			if emptyLeftPlot && emptyRightPlot {
				flowerbed[idx] = 1
				count++

				if count >= n {
					return true
				}
			}
		}

	}

	return count >= n
}
