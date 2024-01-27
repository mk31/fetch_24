package main

import (
	"log"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)


type Receipt struct {
    Retailer    	string  `json:"retailer" binding:"required"`
    PurchaseDate 	string  `json:"purchaseDate" binding:"required"`
    PurchaseTime 	string  `json:"purchaseTime" binding:"required"`
    Items        	[]Item  `json:"items" binding:"required"`
    Total        	string  `json:"total" binding:"required"`
	Points 			int		`json:"points"`
}

type Item struct {
    ShortDescription string `json:"shortDescription" binding:"required"`
    Price            string `json:"price" binding:"required"`
}

// if/when there are other use cases, method definition may change from being a pointer receiver type
func (receipt *Receipt) calculatePoints() {
	floatPoints, floatParseErr := strconv.ParseFloat(receipt.Total, 64)

    if floatParseErr != nil {
        log.Println("Error converting price to float:", floatParseErr)
    }

}