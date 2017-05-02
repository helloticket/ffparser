package helper

import (
	"strconv"
)

func ToInteger64(str string) int64 {
	value, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		panic(err)
	}

	return value
}

func ToInteger(str string) int {
	value, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return value
}

func Float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func ToFloat64(str string) float64 {
	value, err := strconv.ParseFloat(str, 64)

	if err != nil {
		panic(err)
	}

	return value
}
