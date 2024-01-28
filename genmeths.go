package main

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/google/uuid"
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

func IsTrimmedLengthMultipleOf3(str string) bool {
	trimmed := strings.TrimSpace(str)

	strLen := len(trimmed)

	return strLen > 0 && strLen%3 == 0
}

func RoundUpPoint2(price float64) int {
	point2Flt := price * .2
	return int(math.Ceil(point2Flt))
}

// just handling one date format... really would need to handle more date formats.
func IsPurchaseDateDayOdd(dateString string) bool {
	layout := "2006-01-02"

	date, err := time.Parse(layout, dateString)

	if err != nil {
		fmt.Println("Error parsing date:", err)
		return false
	}

	return date.Day()%2 == 1
}

// just handling one time format... really would need to handle more
func IsPurchaseTimeBetween14And16Exclusive(timeString string) bool {
	layout := "15:04"

	time, err := time.Parse(layout, timeString)

	if err != nil {
		fmt.Println("Error parsing time:", err)
		return false
	}

	hour := time.Hour()

	isHour14or15 := hour == 14 || hour == 15

	if !isHour14or15 {
		return false
	}

	minute := time.Minute()

	// we know hour is 2 or 3
	// if 3, all times work
	// if 2, minute must be > 0
	return hour == 3 || minute > 0
}

func GenerateHash() (string, error) {
	u, err := uuid.NewRandom()

	if err != nil {
		fmt.Printf("Something went wrong: %v", err)
		return "", err
	}

	return u.String(), nil
}
