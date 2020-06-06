package ffparser

import (
	"strconv"
	"strings"
)

func toInteger64(str string) int64 {
	if strings.TrimSpace(str) == "" {
		return int64(0)
	}

	value, err := strconv.ParseInt(str, 10, 64)

	if err != nil {
		panic(err)
	}

	return value
}

func toInteger(str string) int {
	if strings.TrimSpace(str) == "" {
		return 0
	}

	value, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}

	return value
}

func float64ToString(v float64) string {
	return strconv.FormatFloat(v, 'f', -1, 64)
}

func toFloat64(str string) float64 {
	if strings.TrimSpace(str) == "" {
		return float64(0)
	}

	value, err := strconv.ParseFloat(str, 64)

	if err != nil {
		panic(err)
	}

	return value
}
