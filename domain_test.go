package main

import (
	"encoding/json"
	"log"
	"testing"
)

func TestCalculatePoints(t *testing.T) {

	testJson1 :=
		`
	{
		"retailer": "Target",
		"purchaseDate": "2022-01-01",
		"purchaseTime": "13:01",
		"items": [
		  {
			"shortDescription": "Mountain Dew 12PK",
			"price": "6.49"
		  },{
			"shortDescription": "Emils Cheese Pizza",
			"price": "12.25"
		  },{
			"shortDescription": "Knorr Creamy Chicken",
			"price": "1.26"
		  },{
			"shortDescription": "Doritos Nacho Cheese",
			"price": "3.35"
		  },{
			"shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
			"price": "12.00"
		  }
		],
		"total": "35.35"
	  }
	`

	var receipt1 Receipt

	receipt1Err := json.Unmarshal([]byte(testJson1), &receipt1)

	if receipt1Err != nil {
		log.Println(receipt1Err)
		return
	}

	testJson2 :=
		`
		{
			"retailer": "M&M Corner Market",
			"purchaseDate": "2022-03-20",
			"purchaseTime": "14:33",
			"items": [
			  {
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  }
			],
			"total": "9.00"
		  }
	`

	var receipt2 Receipt

	receipt2Err := json.Unmarshal([]byte(testJson2), &receipt2)
	if receipt2Err != nil {
		log.Println(receipt2Err)
		return
	}


	testJson3 :=
		`
		{
			"retailer": "M&M Corner Market",
			"purchaseDate": "2022-03-20",
			"purchaseTime": "14:33",
			"items": [
			  {
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  },{
				"shortDescription": "Gatorade",
				"price": "2.25"
			  }
			],
			"total": "9"
		  }
	`

	var receipt3 Receipt

	receipt3Err := json.Unmarshal([]byte(testJson3), &receipt3)
	if receipt3Err != nil {
		log.Println(receipt3Err)
		return
	}



	tests := []struct {
		Receipt Receipt
		ExpectedPoints int
	}{
		{
			Receipt: receipt1,
			ExpectedPoints: 28,			
		},
		{
			Receipt: receipt2,
			ExpectedPoints: 109,			
		},
		{
			Receipt: receipt3,
			ExpectedPoints: 109,			
		},
		
	}
	
	// TODO: Test receipt with more than 2 decimal places
	// TODO: Test receipt with .00 and with no decimals
	for _, test := range tests {

		test.Receipt.calculatePoints()

		if test.Receipt.Points != test.ExpectedPoints {
			t.Errorf("When receipt is: %+v, Expected %v points, but got %v points", test.Receipt, test.ExpectedPoints, test.Receipt.Points)
		}
	}
}
