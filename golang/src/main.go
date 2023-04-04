package main

import (
	"fmt"

	"github.com/cvitorin/complero/ims/src/domain"
	"github.com/cvitorin/complero/ims/src/repositories"
	"github.com/cvitorin/complero/ims/src/services"
)

func main() {
	repo := &repositories.DummyItemRepository{}
	service := services.NewItemService(repo)
	inventory, err := service.GetAllItems()
	if err != nil {
		panic(err)
	}

	inv := domain.NewInventory(inventory)

	printInventory(1, inv)
	inv.UpdateQuality()

}

func printInventory(day int, inventory *domain.Inventory) {
	fmt.Println("Day: ", day)
	for _, item := range inventory.Items {
		fmt.Println("Item: ", item.Name, "Quality: ", item.Quality, "ExpiresIn: ", item.ExpiresIn)
	}
}
