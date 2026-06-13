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
	Name     string
	Price    int
	Quantity int
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
			// customerMenu()
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
		fmt.Println("5. Show statistics")
		fmt.Println("6. Exit")
		fmt.Scan(&choose)
		if choose == 1 {
			addMenu(&*catalog, &*menuCount)
		} else if choose == 2 {
			// delete menu function
			deleteMenu(&*catalog, &*menuCount)
		} else if choose == 3 {
			// edit menu function
			showMenu(*catalog, *menuCount)
			editMenu(&*catalog, &*menuCount)
		} else if choose == 4 {
			// show menu function
			showMenu(*catalog, *menuCount)
		} else if choose == 5{
			statisticPerCatergory(*catalog, *menuCount)
		} else if choose == 6 {
			exit = true
			fmt.Println("Goodbye!")
		} else	 {
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

		fmt.Print("Menu name: ")
		fmt.Scan(&newMenu.Name)

		fmt.Print("Menu category: ")
		fmt.Scan(&newMenu.Category)

		fmt.Print("Menu price: ")
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

		fmt.Print("Menu status (true/false): ")
		fmt.Scan(&newMenu.Status)
		(*catalog)[*menuCount] = newMenu
		*menuCount++
	}
	fmt.Println("Menus added successfully!")
}

func deleteMenu(catalog *menuList, menuCount *int) {
	var idx, i int
	if *menuCount == 0 {
		fmt.Println("The menu was empty, nothing to delete")
		return
	}
	showMenu(*catalog, *menuCount)
	fmt.Print("Enter the name of menu to delete: ")
	var name string
	fmt.Scan(&name)
	idx = SearchMenu(*catalog, *menuCount, name)
	if idx == -1 {
		fmt.Println("Menu not found")
		return
	}
	for i = idx; i < *menuCount-1; i++ {
		(*catalog)[i] = (*catalog)[i+1]
	}
	*menuCount = *menuCount - 1
	fmt.Printf("Menu '%s' deleted succesfully\n", name)
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

		fmt.Print("    Composition: ")
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

func editMenu(catalog *menuList, menuCount *int) {
	var price int
	var enter bool
	var MenuName string
	enter = false
	for !enter {
		fmt.Println("Choose Menu to edit:")
		fmt.Scan(&MenuName)
		for i := 0; i < *menuCount; i++ {
			if MenuName == (*catalog)[i].Name {
				enter = true
				fmt.Println("Enter new price:")
				fmt.Scan(&price)
				(*catalog)[i].Price = price
			} else {
				fmt.Println("Menu not found. Please enter a valid menu name.")
			}
		}
	}
}

// customer menu
func customerMenu(catalog *menuList, menuCount int) {
	fmt.Println("welcome to our cafe!")
	if menuCount == 0 {
		fmt.Println("Sorry, the menu is currently empty, please come back later.")
		return
	}
	fmt.Println("Here is our menu:")
	fmt.Println("categories: Drinks / Food ")
	fmt.Print("Please choose a category: ")
	var category string
	fmt.Scan(&category)
	fmt.Print("Please select a price order (1: Low to High, 2: High to Low): ")
	var priceOrder int
	fmt.Scan(&priceOrder)

	var sorted menuList
	sorted = *catalog
	if priceOrder == 1 {
		sorted = sortbyAsc(sorted, menuCount)
	} else {
		sorted = sortbyDesc(sorted, menuCount)
	}
	fmt.Printf("Here is our %s menu:\n", category)
	showMenuByCategory(sorted, menuCount, category)

	//var currentOrders orderList
	//var orderCount int = 0
	var exit bool
	exit = false
	for !exit {
		fmt.Println("what would you like to do?")
		fmt.Println("1. Add order")
		fmt.Println("2. delete order")
		fmt.Println("3. edit order")
		fmt.Println("4. show bill")
		fmt.Println("5. exit")
		var choose int
		fmt.Scan(&choose)
		if choose == 1 {
			//add order func
		} else if choose == 2 {
			//delete order func deleteMenu(currentOrders *orderList, orderCount *int)
		} else if choose == 3 {
			//edit order

		} else if choose == 4 {
			//func show bill
		}
	}
}

func showMenuByCategory(catalog menuList, menuCount int, category string) {
	var found bool
	found = false
	var count, i int
	count = 0
	for i = 0; i < menuCount; i++ {
		if catalog[i].Category == category {
			count = count + 1
			var menu Menu = catalog[i]
			fmt.Printf("%d. %s - Rp%d (Available: %t)\n", count, menu.Name, menu.Price, menu.Status)
			found = true
		}
	}
	if !found {
		fmt.Printf("Sorry, we don't have any %s menu at the moment.\n", category)
	}
}

func showBill(cart orderList, orderCount int) {
	var i, total_amount, total_item int
	total_item = 0
	total_amount = 0
	fmt.Println("|-------Total Bill-------|")
	for i = 0; i < orderCount; i++ {
		fmt.Printf("| %-15s | %-10d, | %-10d |\n", cart[i].Name, cart[i].Price, cart[i].Quantity)
		total_amount = total_amount + cart[i].Price
		total_item = total_item + cart[i].Quantity
	}
	fmt.Printf("Amount %-15d | total item %d", total_amount, total_item)
}

//helper function
func SearchMenu(catalog menuList, menuCount int, name string) int {
	var i, idx int
	idx = -1
	for i = 0; i < menuCount; i++ {
		if catalog[i].Name == name {
			idx = i
		}
	}
	return idx
}

func sortbyAsc(list menuList, n int) menuList {
	var i, j, minIdx int
	for i = 0; i < n-1; i++ {
		minIdx = i
		for j = i + 1; j < n; j++ {
			if list[j].Price < list[minIdx].Price {
				minIdx = j
			}
		}
		var temp Menu = list[minIdx]
		list[minIdx] = list[i]
		list[i] = temp
	}
	return list
}

func sortbyDesc(list menuList, n int) menuList {
	var i, j int
	var keyMenu Menu
	for i = 1; i < n; i++ {
		keyMenu = list[i]
		j = i - 1
		for j >= 0 && list[j].Price < keyMenu.Price {
			list[j+1] = list[j]
			j = j - 1
		}
		list[j+1] = keyMenu
	}
	return list
}

func statisticPerCatergory(catalog menuList, menuCount int) {
	var drinkCount, foodCount, totalOrders, totalPrice int
	var averagePrice float64
	drinkCount = 0
	foodCount = 0
	totalOrders = 0
	totalPrice = 0
	for i := 0; i < menuCount; i++ {
		if catalog[i].Category == "Drink" {
			drinkCount = drinkCount + 1
		} else if catalog[i].Category == "Food" {
			foodCount = foodCount + 1
		}
		totalOrders = totalOrders + 1
		totalPrice = totalPrice + catalog[i].Price
	}
	if menuCount > 0 {
		averagePrice = float64(totalPrice) / float64(menuCount)
	}
	fmt.Printf("Statistics by Category:\n")
	fmt.Printf("Drinks: %d\n", drinkCount)
	for i := 0; i < menuCount; i++ {
		if catalog[i].Category == "Drink" {
			fmt.Printf(" - %s: Rp%d\n", catalog[i].Name, catalog[i].Price)
		}
	}
	fmt.Printf("Food: %d\n", foodCount)
	for i := 0; i < menuCount; i++ {
		if catalog[i].Category == "Food" {
			fmt.Printf(" - %s: Rp%d\n", catalog[i].Name, catalog[i].Price)
		}
	}
	fmt.Printf("Total Orders: %d\n", totalOrders)
	fmt.Printf("Average Price: Rp%.0f\n", averagePrice)
}