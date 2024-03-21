package piscine

type food struct { // food struct with meal prep times
	preptime int
}

func FoodDeliveryTime(order string) int { // takes order string as input and returns prep time int
	menu := map[string]food{ // map menu; key is food_name and value is food struct
		"burger":  {15},
		"chips":   {10},
		"nuggets": {12},
	}
	item, found := menu[order] // check if item in order is in menu,
	if found {
		return item.preptime
	}
	return 404 // Error: Item not found
}
