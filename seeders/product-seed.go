package seeders

import (
	"encoding/json"
	"log"
	"os"
	"techincal-test/database"
	"techincal-test/structs"
)

func SeedProducts() {
	jsonData, err := os.ReadFile("database/dummy/product.json")
	if err != nil {
		log.Printf("Error reading product.json file : %v", err)
		return
	}

	var products []structs.Product
	if err := json.Unmarshal(jsonData, &products); err != nil {
		log.Printf("Error parsing event.json: %v", err)
	}

	for _, product := range products {
		result := database.DB.Where(&structs.Product{
			ProductName: product.ProductName,
			SKU:         product.SKU,
			Quantity:    product.Quantity,
			Status:      product.Status,
		}).FirstOrCreate(&product)
		if result.Error != nil {
			log.Printf("Failed to seed event %s: %v", product.ProductName, result.Error)
		} else {
			log.Printf("Event %s seeded successfully!", product.ProductName)
		}
	}
}
