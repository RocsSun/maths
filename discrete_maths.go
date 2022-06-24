package maths

import "math"

// Sum 求和。
func Sum[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64 {
	var sum float64

	n := len(nums)

	if n == 0 {
		return 0
	}

	sum = 0
	for _, num := range nums {
		sum += float64(num)
	}
	return sum
}

// EachReduceAvgSquareSums 求每个数减去平均数的平方和。
func EachReduceAvgSquareSums[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64 {
	var sum float64

	n := len(nums)

	if n == 0 {
		return 0
	}

	avg := ArithmeticAverage(nums)

	for _, num := range nums {
		t := float64(num) - avg
		sum += t * t
	}
	return sum
}

// ArithmeticAverage 求取数组的算术平均值
func ArithmeticAverage[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64 {
	if n := len(nums); n == 0 {
		return 0
	} else {
		return Sum(nums) / float64(len(nums))
	}
}

// PopulationVariance 总体方差
func PopulationVariance[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64 {

	if n := len(nums); n == 0 {
		return 0
	} else {
		return EachReduceAvgSquareSums(nums) / float64(len(nums))
	}
}

// SampleVariance 样本方差
func SampleVariance[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64 {
	if n := len(nums); n == 0 {
		return 0
	} else {
		return EachReduceAvgSquareSums(nums) / float64(len(nums)-1)
	}
}

// OverallStandardDeviation 总体标准差。
func OverallStandardDeviation[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64 {
	if n := len(nums); n == 0 {
		return 0
	} else {
		return math.Sqrt(EachReduceAvgSquareSums(nums) / float64(len(nums)))
	}
}

// SampleStandardDeviation 样本标准差。
func SampleStandardDeviation[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64 {
	if n := len(nums); n == 0 {
		return 0
	} else {
		return math.Sqrt(EachReduceAvgSquareSums(nums) / float64(len(nums)-1))
	}
}
