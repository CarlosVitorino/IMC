package domain

type Inventory struct {
	Items []*Item
}

func (i *Inventory) UpdateQuality() {
	for _, item := range i.Items {
		switch item.Name {
		case "Rock":
			item.Quality = 80
			item.ExpiresIn = 0
		case "Aged Cheese":
			item.DecreaseExpiresIn()
			item.IncreaseQuality()
		default:
			item.DecreaseExpiresIn()
			item.DecreaseQuality()
		}
	}
}

func NewInventory(items []*Item) *Inventory {
	return &Inventory{Items: items}
}
