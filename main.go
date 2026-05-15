package main

import "fmt"

const NMAX int = 100

type Menu struct {
	Name     	string
	Category 	string
	Price    	int
	composition string
	status  	bool
}
type menuList [NMAX]Menu

func main() {
	//interface menu  
	fmt.Println("Who are using this cafe's catalog?")
	var user string
	fmt.Scanln(&user)
	// only admin and customer can access the menu
	if user == "admin" {
		adminMenu()
	} else if user == "customer" {
		customerMenu()
	} else {
		fmt.Println("Invalid user")
	}
}

func adminMenu() {
	var list menuList
	var choose int
	fmt.Println("What would you like to do?")
	fmt.Println("1. Add menu")
	fmt.Println("2. Delete menu")
	fmt.Println("3. Edit menu")
	fmt.Println("4. Show menu")
	fmt.Println("5. Exit")
	fmt.Scan(&choose)
	if choose == 1 {
		addMenu(&list)
	} else if choose == 2 {
		// delete menu function
	} else if choose == 3 {
		// edit menu function
	} else if choose == 4 {
		fmt.Println("Goodbye!")
	} else {
		fmt.Println("Invalid option")
	}
}

func addMenu(list *menuList) {
	var n int
	fmt.Println("How many menu do you want to add?")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Menu name:")
		fmt.Scan(&list[i].Name)

		fmt.Print("Menu category:")
		fmt.Scan(&list[i].Category)

		fmt.Print("Menu price:")
		fmt.Scan(&list[i].Price)

		fmt.Print("Menu composition:")
		fmt.Scan(&list[i].composition)

		fmt.Print("Menu status (true/false):")
		fmt.Scan(&list[i].status)
	}
}

func customerMenu() {
	fmt.Println("Welcome to our cafe!")
	fmt.Println("Here is our menu:")
	fmt.Println("1. food")
	fmt.Println("2. drink")
	fmt.Println("3. dessert")
	fmt.Println("please select a category:")
	var category string
	fmt.Scan(&category)

	var priceOrder string
	fmt.Println("please select a price order (asc/desc):")
	fmt.Scan(&priceOrder)

	if priceOrder == "asc" {
		// sort menu by price ascending
	} else if priceOrder == "desc" {
		// sort menu by price descending
	}

	if category == "food" {
		fmt.Println("Here is our food menu:")
		// display food menu
	} else if category == "drink" {
		fmt.Println("Here is our drink menu:")
		// display drink menu
	} else if category == "dessert" {
		fmt.Println("Here is our dessert menu:")
		// display dessert menu
	} else { 
		fmt.Println("Invalid category")
	}

	// ask customer what they want to 
	fmt.Println("What would you like?")
	fmt.Println("1. Add Order")
	fmt.Println("2. Delete Order")
	fmt.Println("3. Edit Order")
	fmt.Println("4. Exit")
	var orderOption int
	fmt.Scan(&orderOption)

	if orderOption == 1 {
		addOrder(&order, list)
	} else if orderOption == 2 {
		// delete order function
	} else if orderOption == 3 {
		// edit order function
	} else if orderOption == 4 {
		fmt.Println("Thank you for visiting our cafe!")
	}

}

func addOrder(order *orderList, list menuList) {
	// Implementation for adding an order
	var n int 
	fmt.Println("How many menu do you want to order?")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Print("Menu name:")
		fmt.Scan(&order[i].Name)
		
		fmt.Print("Menu price:")
		fmt.Scan(&order[i].Price)
	}
}