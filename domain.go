package main

type Receipt struct {
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
func (receipt *Receipt) calculatePoints() {

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

	mult2Points := CountMultiplesOf2(len(receipt.Items))*5

	mult3Pts := receipt.calculateItemDescriptionLengthPoints()

	points6 := 0

	if IsPuchaseDateDayOdd(receipt.PurchaseDate) {
		points6 = 6
	}

	points10 := 0

	if IsPuchaseTimeBetween14And16Exclusive(receipt.PurchaseTime) {
		points10 = 10
	}

	receipt.Points = alphaPoints + pts50 + mult25Pts + mult2Points + mult3Pts + points6 + points10

}

func (receipt Receipt) calculateItemDescriptionLengthPoints() int {
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
