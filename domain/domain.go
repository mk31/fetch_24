package domain

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

type Receipt struct {
	Id           string `json:"id"`
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Items        []Item `json:"items" binding:"required"`
	Total        string `json:"total" binding:"required"`
	Points       int    `json:"points"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

// if/when there are other use cases, method definition may change from being a pointer receiver type
func (receipt *Receipt) CalculatePoints() {

	/*
		√ * One point for every alphanumeric character in the retailer name.
		√ * 50 points if the total is a round dollar amount with no cents.
		√ * 25 points if the total is a multiple of `0.25`.
		√ * 5 points for every two items on the receipt.
		√ * If the trimmed length of the item description is a multiple of 3,
				multiply the price by `0.2` and
				round up to the nearest integer.
				The result is the number of points earned.
		√ * 6 points if the day in the purchase date is odd.
		√ * 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	*/

	alphaPoints := CountAlphanumeric(receipt.Retailer)

	pts50 := 0

	if DoesStringPriceEndInZero(receipt.Total) {
		pts50 = 50
	}

	mult25Pts := 0

	if DoesStringPriceEndIn25Multiple(receipt.Total) {
		mult25Pts = 25
	}

	mult2Points := CountMultiplesOf2(len(receipt.Items)) * 5

	mult3Pts := receipt.calculateItemDescriptionLengthPoints()

	points6 := 0

	if IsPurchaseDateDayOdd(receipt.PurchaseDate) {
		points6 = 6
	}

	points10 := 0

	if IsPurchaseTimeBetween14And16Exclusive(receipt.PurchaseTime) {
		points10 = 10
	}

	receipt.Points = alphaPoints + pts50 + mult25Pts + mult2Points + mult3Pts + points6 + points10

}

func (receipt *Receipt) calculateItemDescriptionLengthPoints() int {
	totalPoints := 0

	for _, item := range receipt.Items {
		if IsTrimmedLengthMultipleOf3(item.ShortDescription) {
			twoDecFlt, twoDecErr := TwoDecimalFloatFromString(item.Price)

			if twoDecErr != nil {
				continue
			}

			totalPoints += RoundUpPoint2(twoDecFlt)
		}
	}

	return totalPoints
}

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
