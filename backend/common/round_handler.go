package common

import (
	"math"
	"time"
)

const maxDecimalNumber = 3

func roundToDecimal(num float32, decimalPlaces int) float32 {
	shift := math.Pow10(decimalPlaces)
	return float32(math.Round(float64(num)*shift)) / float32(shift)
}

func RoundToInt(num float32) int {
	return int(roundToDecimal(num, 0))
}

func CustomRound(num *float32) {
	roundedNum := roundToDecimal(*num, maxDecimalNumber)
	*num = roundedNum
}

func RoundTime(date time.Time) (*time.Time, string) {
	roundedTime := date.Round(time.Hour)

	result := roundedTime.Format(time.RFC3339)

	return &roundedTime, result
}
