package domain

import "strings"

type Item struct {
	Name               string
	ExpiresIn, Quality int
}

func (i *Item) IsExpired() bool {
	return i.ExpiresIn < 0
}

func (i *Item) IsConjured() bool {
	return strings.HasPrefix(i.Name, "Conjured")
}

func (i *Item) DecreaseQuality() {
	if i.Quality > 0 {
		if i.IsConjured() {
			i.Quality -= 2
			return
		}

		i.Quality--
	}

	if i.Quality < 0 {
		i.Quality = 0
	}
}

func (i *Item) IncreaseQuality() {
	if i.Quality < 50 {
		switch i.Name {
		case "Aged Cheese":
			if !i.IsExpired() {
				i.Quality += 2
				return
			}

			i.Quality = 0
		default:
			i.Quality++
		}
	}
}

func (i *Item) DecreaseExpiresIn() {
	i.ExpiresIn--
	if i.IsExpired() {
		i.DecreaseQuality()
	}
}

func NewItem(name string, expiresIn, quality int) *Item {
	return &Item{
		Name:      name,
		ExpiresIn: expiresIn,
		Quality:   quality,
	}
}
