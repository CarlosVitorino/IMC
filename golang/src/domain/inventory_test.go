package domain

import "testing"

func TestInventory_UpdateQuality(t *testing.T) {
	tests := []struct {
		name   string
		before []*Item
		after  []*Item
	}{
		// regular item
		{
			name:   "regular item before expiry",
			before: []*Item{NewItem("foo", 1, 10)},
			after:  []*Item{NewItem("foo", 0, 9)},
		},
		{
			name:   "regular item after expiry",
			before: []*Item{NewItem("foo", 0, 10)},
			after:  []*Item{NewItem("foo", -1, 8)},
		},
		// aged cheese
		{
			name:   "quality increases by 2 before expiry",
			before: []*Item{NewItem("Aged Cheese", 1, 10)},
			after:  []*Item{NewItem("Aged Cheese", 0, 12)},
		},
		{
			name:   "quality becomes 0 when expired",
			before: []*Item{NewItem("Aged Cheese", 0, 10)},
			after:  []*Item{NewItem("Aged Cheese", -1, 0)},
		},
		{
			name:   "max quality of 50 before expiry. quality stays the same",
			before: []*Item{NewItem("Aged Cheese", 10, 50)},
			after:  []*Item{NewItem("Aged Cheese", 9, 50)},
		},
		{
			name:   "max quality of 50. Quality becomes 0 when expired",
			before: []*Item{NewItem("Aged Cheese", 0, 50)},
			after:  []*Item{NewItem("Aged Cheese", -1, 0)},
		},
		// rock item
		{
			name:   "rock item",
			before: []*Item{NewItem("Rock", 10, 12)},
			after:  []*Item{NewItem("Rock", 0, 80)},
		},
		// conjured item
		{
			name:   "conjured item before expiry",
			before: []*Item{NewItem("Conjured Sword", 10, 12)},
			after:  []*Item{NewItem("Conjured Sword", 9, 10)},
		},
		{
			name:   "conjured item after expiry",
			before: []*Item{NewItem("Conjured Sword", 0, 12)},
			after:  []*Item{NewItem("Conjured Sword", -1, 8)},
		},
		{
			name:   "conjured item after expiry with quality of 1",
			before: []*Item{NewItem("Conjured Sword", 0, 1)},
			after:  []*Item{NewItem("Conjured Sword", -1, 0)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			inventory := NewInventory(tt.before)
			inventory.UpdateQuality()

			for i := range inventory.Items {
				if inventory.Items[i].Name != tt.after[i].Name {
					t.Errorf("expected item name %q, got %q", tt.after[i].Name, inventory.Items[i].Name)
				}
				if inventory.Items[i].ExpiresIn != tt.after[i].ExpiresIn {
					t.Errorf("expected item %q to have expires-in of %d, got %d", inventory.Items[i].Name, tt.after[i].ExpiresIn, inventory.Items[i].ExpiresIn)
				}
				if inventory.Items[i].Quality != tt.after[i].Quality {
					t.Errorf("expected item %q to have quality of %d, got %d", inventory.Items[i].Name, tt.after[i].Quality, inventory.Items[i].Quality)
				}
			}
		})
	}
}
