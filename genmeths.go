package main

import (
	"log"
	"math"
	"strconv"
	"strings"
	"unicode"
)

func CountAlphanumeric(str string) int {
	count := 0
	for _, char := range str {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			count++
		}
	}
	return count
}

func DoesStringPriceEndInZero(str string) bool {
	twoDecimalFloat, err := TwoDecimalFloatFromString(str)

	if err != nil {
		log.Println("Error creating Two decimal float from string:", err)
		return false
	}

	return endsWithZeroZero(twoDecimalFloat)
}

func TwoDecimalFloatFromString(floatString string) (float64, error) {
	floatPoints, floatParseErr := strconv.ParseFloat(strings.TrimSpace(floatString), 64)

	if floatParseErr != nil {
		log.Println("Error converting price to float:", floatParseErr)
		return -1, floatParseErr
	}

	return roundTo2Decimals(floatPoints), nil
}

func DoesStringPriceEndIn25Multiple(str string) bool {
	twoDecimalFloat, err := TwoDecimalFloatFromString(str)

	if err != nil {
		log.Println("Error creating Two decimal float from string:", err)
		return false
	}

	return isMultipleOf25(twoDecimalFloat)
}

func CountMultiplesOf2(count int) int {
	return count / 2
}

func isMultipleOf25(twoDecimalFloat float64) bool {
	return twoDecimalFloat*4 == float64(int(twoDecimalFloat*4))
}

func roundTo2Decimals(f float64) float64 {
	return math.Round(f*100) / 100
}

func endsWithZeroZero(f float64) bool {
	return f == float64(int(f))
}
