package main

import (
	"fmt"
)

var hobbies = [3]string{"guitar", "gym", "coding"}

var goals = []string{"finish", "make api"}

type Product struct {
	id    int
	title string
	price float64
}

func main() {
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])
	fmt.Println(hobbies[0:2])
	newSlice := []string{hobbies[0], hobbies[1], hobbies[2]}
	newTwoSlice := []string{hobbies[1], hobbies[2]}
	fmt.Println(newSlice)
	fmt.Println(newTwoSlice)
	goals[1] = "reach end"
	goals = append(goals, "make cli tool")
	newSlice = newSlice[1:3]
	fmt.Println(newSlice)
	fmt.Println(goals)

	first_product := Product{
		id:    1,
		title: "apple",
		price: 699.9,
	}

	second_product := Product{
		id:    2,
		title: "pear",
		price: 9.9,
	}

	third_product := Product{
		id:    3,
		title: "banana",
		price: 9123.9,
	}

	var items = []Product{first_product, second_product}
	fmt.Println(items)
	items = append(items, third_product)
	fmt.Println(items)

}
