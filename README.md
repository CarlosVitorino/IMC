# IMC
Inventory Management Center Exercise

> Considering the Code Exercise in the file exercise.go

We’re not very happy with the quality of the existing code base, but it does work. Feel free to make any changes to the update-quality method. You may add any new code as long as everything still works and the tests still pass. You are welcome to refactor the function if you spot places where the code quality can be improved.


## First an introduction to our system:
* All items have an expires-in value which denotes the number of days we have to sell the item
* All items have a quality value which denotes how valuable the item is
* At the end of each day our system lowers both values for every item

Pretty simple, right? Well this is where it gets interesting:
* Once the expires-in is negative, quality degrades twice as fast
* The quality of an item is never negative
* “Aged Cheese” actually increases in quality the older it gets
* The quality of an item is never more than 50

We’d like you to add some new functionality to our inventory management system.

## New Features:

### Stage 1
We want to add support for a new item: “Rock”. Being a rock, it never expires or decreases in quality. It always has a quality of 80 and an expires-in of 0.

### Stage 2
It turns out that our “Aged Cheese” is selling better than we thought! But, people do not seem to like it much when it has expired. We think that the quality should actually increase by 2 until it expires. Once it expires, its quality drops down to 0.

### Stage 3
We have recently signed a supplier of “Conjured” items. This requires an update to our system. “Conjured” items degrade in quality twice as fast as normal items. Conjured items will be any items that start with the string “Conjured”, followed by a space.

