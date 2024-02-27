package anscombe

import (
	"math"
	"sort"
)

func CalcMean(nums []int) float64 {
	var sum int
	for _, val := range nums {
		sum = sum + val
	}
	return float64(sum) / float64(len(nums))
}

func CalcMedian(nums []int) float64 {
	sort.Ints(nums)
	if len(nums)%2 == 1 {
		return float64(nums[len(nums)/2])
	} else {
		return (float64(nums[len(nums)/2]) + float64(nums[len(nums)/2-1])) / 2
	}
}

func CalcMode(nums []int) int {

	dict := make(map[int]int)

	for _, num := range nums {
		if v, ok := dict[num]; ok {
			dict[num] = v + 1
		} else {
			dict[num] = 1
		}
	}

	bestNum := nums[0]
	bestTimes := dict[bestNum]
	for num, times := range dict {
		if times > bestTimes || (times == bestTimes && num < bestNum) {
			bestTimes = times
			bestNum = num
		}
	}

	return bestNum
}

func calcVariance(data []int, mean float64) float64 {
	var total float64
	for _, value := range data {
		total += math.Pow(float64(value)-mean, 2)
	}
	return total / float64(len(data))
}

func CalcDS(nums []int) float64 {

	mean := CalcMean(nums)               // Расчет среднего значения
	variance := calcVariance(nums, mean) // Расчет дисперсии
	stdDev := math.Sqrt(variance)        // Расчет стандартного отклонения

	return stdDev
}
