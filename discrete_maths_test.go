package maths_test

import (
	"fmt"
	"testing"

	"github.com/RocsSun/maths"
)

func TestSum(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	sum := maths.Sum(a)
	if sum != 15.0 {
		t.Error("测试失败。")
	}
}

func TestSum1(t *testing.T) {
	var a = []int{}
	sum := maths.Sum(a)
	if sum != 0.0 {
		t.Error("测试失败。")
	}
}

func TestEachReduceAvgSquareSums(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	sum := maths.EachReduceAvgSquareSums(a)
	if sum != 10.0 {
		t.Error("测试失败。")
	}
}

func TestEachReduceAvgSquareSums1(t *testing.T) {
	var a = []int{}
	sum := maths.EachReduceAvgSquareSums(a)
	if sum != 0.0 {
		t.Error("测试失败。")
	}
}

func TestArithmeticAverage(t *testing.T) {
	var a = []int{1, 2, 3, 4, 5}
	avg := maths.ArithmeticAverage(a)
	if avg != 3.0 {
		t.Error("测试失败。")
	}
}

func TestArithmeticAverage1(t *testing.T) {
	var a = []int{1, 2, 3, 4, 4}
	avg := maths.ArithmeticAverage(a)
	if avg != 2.8 {
		t.Log("测试失败。")
	}
}

func TestArithmeticAverage2(t *testing.T) {
	var a = []int{}
	avg := maths.ArithmeticAverage(a)
	if avg != 0.0 {
		t.Log("测试失败。")
	}
}

func TestOverallStandardDeviation(t *testing.T) {
	var a = []int{1, 2, 3, 4, 4}
	avg := maths.OverallStandardDeviation(a)

	if fmt.Sprintf("%.3f", avg) != "1.166" {
		t.Log("测试失败。")
	}
}

func TestOverallStandardDeviation1(t *testing.T) {
	var a []int = []int{}
	avg := maths.OverallStandardDeviation(a)

	if fmt.Sprintf("%.3f", avg) != "0.000" {
		t.Log("测试失败。")
	}
}

func TestSampleStandardDeviation(t *testing.T) {
	var a = []int{1, 2, 3, 4, 4}
	avg := maths.SampleStandardDeviation(a)
	if fmt.Sprintf("%.3f", avg) != "1.304" {
		t.Error("测试失败。")
	}
}

func TestSampleStandardDeviation1(t *testing.T) {
	var a = []int{}
	avg := maths.SampleStandardDeviation(a)
	if avg != 0.0 {
		t.Error("测试失败。")
	}
}

func TestPopulationVariance(t *testing.T) {
	var a = []int{1, 2, 3, 4, 4}
	avg := maths.PopulationVariance(a)

	if fmt.Sprintf("%.3f", avg) != "1.360" {
		t.Error("测试失败。")
	}
}

func TestPopulationVariance1(t *testing.T) {
	var a = []int{}
	avg := maths.PopulationVariance(a)

	if avg != 0.0 {
		t.Error("测试失败。")
	}
}

func TestSampleVariance(t *testing.T) {
	var a = []int{1, 2, 3, 4, 4}
	avg := maths.SampleVariance(a)

	if fmt.Sprintf("%.3f", avg) != "1.700" {
		t.Error("测试失败。")
	}
}

func TestSampleVariance1(t *testing.T) {
	var a = []int{}
	avg := maths.SampleVariance(a)

	if avg != 0.0 {
		t.Error("测试失败。")
	}
}
