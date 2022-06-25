# maths

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/RocsSun/maths)
![GitHub](https://img.shields.io/github/license/RocsSun/maths)
[![Go](https://github.com/RocsSun/maths/actions/workflows/test.yml/badge.svg)](https://github.com/RocsSun/maths/actions/workflows/test.yml)
[![codecov](https://codecov.io/gh/RocsSun/maths/branch/main/graph/badge.svg?token=LTSXUEINY3)](https://codecov.io/gh/RocsSun/maths)
![GitHub language count](https://img.shields.io/github/languages/count/RocsSun/maths)

Some encapsulation of mathematical methods.

## install and import.

`go version >= 1.18`

`go get -u github.com/RocsSun/maths`

`import "github.com/RocsSun/maths"`

## functions

### func ArithmeticAverage[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64

Calculate the arithmetic mean of the array

### func PopulationVariance[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64

Calculate the total variance of the array

### func SampleVariance[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64

Calculate the sample variance of the array

### func OverallStandardDeviation[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64

Calculate the overall standard deviation of the array.

### func SampleStandardDeviation[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64

Calculate the sample standard deviation of the array.

### func Sum[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64

Calculate cumulative sum

### func EachReduceAvgSquareSums[args int | int8 | int16 | int32 | int64 | float64 | float32](nums []args) float64

Calculate the sum of the squares of each number minus the average.

## 参与贡献

Welcome to add functions.

