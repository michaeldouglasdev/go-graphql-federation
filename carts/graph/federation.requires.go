package graph

import (
	"context"
	"encoding/json"
	"fmt"
	"go-graphql-apollo-federation/carts/graph/model"
)

// PopulateCartRequires is the requires populator for the Cart entity.
func (ec *executionContext) PopulateCartRequires(ctx context.Context, entity *model.Cart, reps map[string]interface{}) error {
	if reps["items"] != nil {
		items, ok := reps["items"].([]interface{})
		if !ok {
			return fmt.Errorf("expected items to be a []interface{}, got %T", reps["items"])
		}

		total := 0.0

		for _, itemInterface := range items {
			itemMap, ok := itemInterface.(map[string]interface{})
			if !ok {
				return fmt.Errorf("expected itemInterface to be a map[string]interface{}, got %T", itemInterface)
			}

			// Imprimir itemMap para depuração
			fmt.Printf("itemMap: %+v\n", itemMap)

			fmt.Printf("", itemMap)
			item, ok := itemMap["item"].(map[string]interface{})
			if !ok {
				if itemMap["item"] == nil {
					fmt.Printf("item is nil in itemMap")
				} else {
					return fmt.Errorf("expected item to be a map[string]interface{}, got %T", itemMap["item"])
				}
				continue // Pular para o próximo item
			}

			// Imprimir item para depuração
			fmt.Printf("item: %+v\n", item)

			price, ok := item["price"].(map[string]interface{})
			if !ok {
				if item["price"] == nil {
					fmt.Printf("price is nil in item")
				} else {
					return fmt.Errorf("expected price to be a map[string]interface{}, got %T", item["price"])
				}
				continue // Pular para o próximo item
			}

			// Imprimir price para depuração
			fmt.Printf("price: %+v\n", price)

			value, ok := price["value"].(json.Number)
			if !ok {
				if price["value"] == nil {
					fmt.Printf("value is nil in price")
				} else {
					return fmt.Errorf("expected value to be a float64, got %T", price["value"])
				}
				continue // Pular para o próximo item
			}

			floatValue, err := value.Float64()
			if err != nil {
				return fmt.Errorf("failed to convert json.Number to float64: %v", err)
			}
			total += floatValue
		}

		entity.TotalPrice = &model.CartTotalPrice{Value: total}
	}
	return nil
}
