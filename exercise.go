package main

import (
	"fmt"
	"os"
	"strings"
)

/**********************************
* Please don't alter this struct *
**********************************/
type Item struct {
	name               string
	expiresIn, quality int
}

var items = []Item{}

func New() []Item {
	return []Item{Item{"Rock1", 0, 80}}
}

/**************************************************
* UpdateQuality is run at the end of each day to *
* update our inventory                           *
**************************************************/
func UpdateQuality() {
	for i := 0; i < len(items); i++ {

		if items[i].name == "Rock" {
			items[i].quality = 80
			items[i].expiresIn = 0
			continue
		}

		if strings.HasPrefix(items[i].name, "Conjured") {
			items[i].quality = items[i].quality - 2
		}

		if items[i].name != "Aged Cheese" {
			if items[i].quality > 0 {
				items[i].quality = items[i].quality - 1
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 2
			}
		}

		items[i].expiresIn = items[i].expiresIn - 1

		if items[i].expiresIn < 0 {
			if items[i].name != "Aged Cheese" {
				if items[i].quality > 0 {
					fmt.Println("test: %s", items[i].quality)
					items[i].quality = items[i].quality - 1
				}
			} else {
				items[i].quality = 0
			}
		}

	}
}

/***********************************************
* Tests                                       *
***********************************************/
type Test func() error

func Test_RegularItem() error {
	items := []Item{Item{"foo", 1, 10}}
	UpdateQuality()

	if items[0].expiresIn != 0 {
		return fmt.Errorf("Expected item to have expires-in of 0, got %v", items[0].expiresIn)
	}

	if items[0].quality != 9 {
		return fmt.Errorf("Expected item to have quality of 9, got %v", items[0].quality)
	}

	return nil
}

func Test_ExpiredRegularItem() error {
	items = []Item{Item{"foo", 0, 10}}
	//x := New()
	//x.UpdateQuality()
	UpdateQuality()

	if items[0].expiresIn != -1 {
		return fmt.Errorf("Expected item to have expires-in of -1, got %v", items[0].expiresIn)
	}

	if items[0].quality != 8 {
		return fmt.Errorf("Expected item to have quality of 8, got %v", items[0].quality)
	}

	return nil
}

func Test_AgedCheese() error {
	testCases := map[string]struct {
		itemsBefore []Item
		itemsAfter  []Item
	}{
		"quality increases by 2 before expiry": {
			itemsBefore: []Item{Item{"Aged Cheese", 1, 10}},
			itemsAfter:  []Item{{"Aged Cheese", 0, 12}},
		},
		"quality becomes 0 when expired": {
			itemsBefore: []Item{Item{"Aged Cheese", 0, 10}},
			itemsAfter:  []Item{{"Aged Cheese", -1, 0}},
		},
		"max quality of 50 after expiry. quality stays the same": {
			itemsBefore: []Item{Item{"Aged Cheese", 0, 50}},
			itemsAfter:  []Item{{"Aged Cheese", -1, 0}},
		},
	}

	for k, testCase := range testCases {
		fmt.Println("case: " + k)

		items = testCase.itemsBefore
		UpdateQuality()
		if items[0] != testCase.itemsAfter[0] {
			return fmt.Errorf("Failure at case: " + k)
		}
	}
	return nil
}

func Test_ExpiredAgedCheeseMaxQuality() error {
	items = []Item{Item{"Aged Cheese", 0, 50}}
	UpdateQuality()

	if items[0].expiresIn != -1 {
		return fmt.Errorf("Expected item to have expires-in of -1, got %v", items[0].expiresIn)
	}

	if items[0].quality != 50 {
		return fmt.Errorf("Expected item to have quality of 50, got %v", items[0].quality)
	}

	return nil
}

func Test_Rock() error {
	testCases := map[string]struct {
		itemsBefore []Item
		itemsAfter  []Item
	}{
		"different quality": {
			itemsBefore: []Item{{"Rock", 10, 12}},
			itemsAfter:  []Item{{"Rock", 0, 80}},
		},
		"same quality and expires in": {
			itemsBefore: []Item{{"Rock", 0, 80}},
			itemsAfter:  []Item{{"Rock", 0, 80}},
		},
	}

	for k, testCase := range testCases {
		fmt.Println("case: " + k)

		items = testCase.itemsBefore
		UpdateQuality()
		if items[0] != testCase.itemsAfter[0] {
			return fmt.Errorf("Failure at case: " + k)
		}
	}
	return nil
}

func main() {
	tests := []Test{
		Test_RegularItem,
		Test_ExpiredRegularItem,
		Test_AgedCheese,
		Test_Rock,
	}

	allPassed := true
	for _, test := range tests {
		err := test()
		if err != nil {
			allPassed = false
			fmt.Println("Test Failure:", err.Error())
		}
	}

	if allPassed {
		fmt.Println("All tests passed")
		os.Exit(0)
	}

	os.Exit(1)
}
