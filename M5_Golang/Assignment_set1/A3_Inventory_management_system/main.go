package main

import (
	"fmt"
	"sort"
	"strconv"
)

type Item struct {
	Code   int
	Title  string
	Cost   float64
	Amount int
}

func addItem(inventory []Item, code int, title string, costInput string, quantity int) ([]Item, error) {
	cost, err := strconv.ParseFloat(costInput, 64)
	if err != nil {
		return inventory, fmt.Errorf("invalid cost input")
	}

	item := Item{
		Code:   code,
		Title:  title,
		Cost:   cost,
		Amount: quantity,
	}
	inventory = append(inventory, item)
	return inventory, nil
}

func modifyQuantity(inventory []Item, code int, updatedQuantity int) ([]Item, error) {
	for i := 0; i < len(inventory); i++ {
		if inventory[i].Code == code {
			if updatedQuantity < 0 {
				return inventory, fmt.Errorf("quantity cannot be negative")
			}
			inventory[i].Amount = updatedQuantity
			return inventory, nil
		}
	}
	return inventory, fmt.Errorf("item not found")
}

func findItem(inventory []Item, keyword string) (*Item, error) {
	for _, item := range inventory {
		if strconv.Itoa(item.Code) == keyword || item.Title == keyword {
			return &item, nil
		}
	}
	return nil, fmt.Errorf("item not found")
}

func showInventory(inventory []Item) {
	fmt.Println("Code\tTitle\t\tCost\tQuantity")
	fmt.Println("----------------------------------------")
	for _, item := range inventory {
		fmt.Printf("%d\t%s\t%.2f\t%d\n", item.Code, item.Title, item.Cost, item.Amount)
	}
}

func arrangeByCost(inventory []Item) {
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i].Cost < inventory[j].Cost
	})
}

func arrangeByQuantity(inventory []Item) {
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i].Amount < inventory[j].Amount
	})
}

func main() {
	var inventory []Item

	inventory, _ = addItem(inventory, 101, "Tablet", "650.40", 12)
	inventory, _ = addItem(inventory, 102, "Headphones", "120.99", 50)

	inventory, err := modifyQuantity(inventory, 101, 20)
	if err != nil {
		fmt.Println(err)
	}

	item, err := findItem(inventory, "Tablet")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Item found: %+v\n", *item)
	}

	showInventory(inventory)

	arrangeByCost(inventory)
	fmt.Println("\nArranged by Cost:")
	showInventory(inventory)

	arrangeByQuantity(inventory)
	fmt.Println("\nArranged by Quantity:")
	showInventory(inventory)
}
