package main

import "fmt"

const NMAX int = 100

type Menu struct {
	Name            string
	Category        string
	Price           int
	Ingredients     [10]string
	IngredientCount int
	Status          bool
}
type menuList [NMAX]Menu

type Order struct {
	Name  string
	Price int
}
type orderList [NMAX]Order

func main() {
	var catalog menuList
	var menuCount int = 0
	// var currOrders orderList
	// var orderCount int = 0
	var exit bool = false

	for !exit {
		fmt.Println("Welcome to our cafe!")
		fmt.Println("login as (admin/customer/exit):")
		var user string
		fmt.Scan(&user)
		// only admin and customer can access the menu
		if user == "admin" {
			adminMenu(&catalog, &menuCount)
		} else if user == "customer" {
			customerMenu()
		} else if user == "exit" {
			exit = true
			fmt.Println("Goodbye!")
		} else {
			fmt.Println("Invalid user. Please type exactly 'admin', 'customer', or 'exit'.")
		}
	}
}

func adminMenu(catalog *menuList, menuCount *int) {
	var exit bool = false
	for !exit {
		var choose int
		fmt.Println("\n--- Admin Menu ---")
		fmt.Println("1. Add menu")
		fmt.Println("2. Delete menu")
		fmt.Println("3. Edit menu")
		fmt.Println("4. Show menu")
		fmt.Println("5. Exit")
		fmt.Scan(&choose)
		if choose == 1 {
			addMenu(&*catalog, &*menuCount)
		} else if choose == 2 {
			// delete menu function
		} else if choose == 3 {
			// edit menu function
		} else if choose == 4 {
			showMenu(*catalog, *menuCount)
		} else if choose == 5 {
			exit = true
			fmt.Println("Goodbye!")
		} else {
			fmt.Println("Invalid option")
		}
	}
}

func addMenu(catalog *menuList, menuCount *int) {
	var n int
	fmt.Println("How many menu do you want to add?")
	fmt.Scan(&n)
	for i := 0; i < n; i++ {

		if *menuCount >= NMAX {
			fmt.Println("Menu list is full. Cannot add more menu.")
			return
		}

		var newMenu Menu

		fmt.Print("Menu name:")
		fmt.Scan(&newMenu.Name)

		fmt.Print("Menu category:")
		fmt.Scan(&newMenu.Category)

		fmt.Print("Menu price:")
		fmt.Scan(&newMenu.Price)

		fmt.Print("How many ingredients are in the composition? (Max 10): ")
		fmt.Scan(&newMenu.IngredientCount)
		if newMenu.IngredientCount > 10 {
			fmt.Println("Limited to 10 ingredients.")
			newMenu.IngredientCount = 10
		}
		if newMenu.IngredientCount > 0 {
			fmt.Println("Enter each ingredient: ")
			var j int
			for j = 0; j < newMenu.IngredientCount; j++ {
				fmt.Printf("Ingredient %d: ", j+1)
				fmt.Scan(&newMenu.Ingredients[j])
			}
		}

		fmt.Print("Menu status (true/false):")
		fmt.Scan(&newMenu.Status)
		(*catalog)[*menuCount] = newMenu
		*menuCount++
	}
	fmt.Println("Menus added successfully!")
}

func showMenu(catalog menuList, menuCount int) {
	if menuCount == 0 {
		fmt.Println("The menu is currently empty.")
		return
	}
	fmt.Println("\n--- Current Menu Catalog ---")
	var i int
	for i = 0; i < menuCount; i++ {
		var menu Menu = catalog[i]
		fmt.Printf("%d. [%s] %s - Rp%d (Available: %t)\n", i+1, menu.Category, menu.Name, menu.Price, menu.Status)

		fmt.Print("    Composition:")
		if menu.IngredientCount == 0 {
			fmt.Print(" None")
		} else {
			var j int
			for j = 0; j < menu.IngredientCount; j++ {
				fmt.Print(menu.Ingredients[j])
				if j < menu.IngredientCount-1 {
					fmt.Print(", ")
				}
			}
		}
		fmt.Println()
	}
}

func customerMenu() {
	fmt.Println("Welcome to our cafe!")
	fmt.Println("Here is our menu:")
	fmt.Println("1. drink")
	fmt.Println("2. dessert")
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
		// addOrder(&order, list)
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

// func selectionSortAsc(list menuList) {
// 	// Implementation for selection sort ascending
// 	var i, j int
// 	var minIndex int
// 	for i = 0; i < n-1; i++ {
// 		minIndex = i
// 		for j = i + 1; j < n; j++ {
// 			if list[j].Price < list[minIndex].Price {
// 				minIndex = j
// 			}
// 		}
