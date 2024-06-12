package models

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Promotion struct {
	ID             string
	Price          float64
	ExpirationDate string
}

// GetPromotionFromDatabase gets promo from db
func GetPromotionFromDatabase(id string) (Promotion, error) {
	file, err := os.Open("promotions.csv")
	if err != nil {
		if os.IsNotExist(err) {
			return Promotion{}, fmt.Errorf("promotions.csv is not present, kindly upload it")
		}
		return Promotion{}, err
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		return Promotion{}, err
	}

	for _, record := range records {
		if record[0] == id {
			price, err := strconv.ParseFloat(record[1], 64)
			if err != nil {
				return Promotion{}, fmt.Errorf("error converting price: %v", err)
			}
			return Promotion{
				ID:             record[0],
				Price:          price,
				ExpirationDate: record[2],
			}, nil
		}
	}
	return Promotion{}, fmt.Errorf("promotion not found")

}

// DelFile deletes the promo file
func DelFile() error {
	err := os.Remove("promotion.csv`")
	if err != nil {
		return err
	}
	return nil
}
